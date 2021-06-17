/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package planbuilder

import (
	"sort"
	"strings"

	vtrpcpb "vitess.io/vitess/go/vt/proto/vtrpc"
	"vitess.io/vitess/go/vt/sqlparser"
	"vitess.io/vitess/go/vt/vtgate/engine"
	"vitess.io/vitess/go/vt/vtgate/semantics"
	"vitess.io/vitess/go/vt/vtgate/vindexes"

	"vitess.io/vitess/go/vt/vterrors"
)

func transformToLogicalPlan(tree joinTree, semTable *semantics.SemTable) (logicalPlan, error) {
	switch n := tree.(type) {
	case *routePlan:
		return transformRoutePlan(n)

	case *joinPlan:
		return transformJoinPlan(n, semTable)
	}

	return nil, vterrors.Errorf(vtrpcpb.Code_INTERNAL, "[BUG] unknown type encountered: %T", tree)
}

func transformJoinPlan(n *joinPlan, semTable *semantics.SemTable) (logicalPlan, error) {
	lhs, err := transformToLogicalPlan(n.lhs, semTable)
	if err != nil {
		return nil, err
	}
	rhs, err := transformToLogicalPlan(n.rhs, semTable)
	if err != nil {
		return nil, err
	}
	return &joinGen4{
		Left:  lhs,
		Right: rhs,
		Cols:  n.columns,
		Vars:  n.vars,
	}, nil
}

func transformRoutePlan(n *routePlan) (*route, error) {
	var tablesForSelect sqlparser.TableExprs
	tableNameMap := map[string]interface{}{}

	sort.Sort(n._tables)
	for _, t := range n._tables {
		tablesForSelect = append(tablesForSelect, relToTableExpr(t))
		for _, name := range t.tableNames() {
			tableNameMap[name] = nil
		}
	}

	for _, predicate := range n.vindexPredicates {
		switch predicate := predicate.(type) {
		case *sqlparser.ComparisonExpr:
			if predicate.Operator == sqlparser.InOp {
				predicate.Right = sqlparser.ListArg(engine.ListVarName)
			}
		}
	}

	for _, leftJoin := range n.leftJoins {
		var lft sqlparser.TableExpr
		if len(tablesForSelect) == 1 {
			lft = tablesForSelect[0]
		} else {
			lft = &sqlparser.ParenTableExpr{Exprs: tablesForSelect}
		}

		rightExpr := relToTableExpr(leftJoin.tbl)
		joinExpr := &sqlparser.JoinTableExpr{
			Join: sqlparser.LeftJoinType,
			Condition: sqlparser.JoinCondition{
				On: leftJoin.pred,
			},
			RightExpr: rightExpr,
			LeftExpr:  lft,
		}
		tablesForSelect = sqlparser.TableExprs{joinExpr}
		// todo: add table
		// tableNameMap[sqlparser.String()] = nil
	}

	predicates := n.Predicates()
	var where *sqlparser.Where
	if predicates != nil {
		where = &sqlparser.Where{Expr: predicates, Type: sqlparser.WhereClause}
	}

	var singleColumn vindexes.SingleColumn
	if n.vindex != nil {
		singleColumn = n.vindex.(vindexes.SingleColumn)
	}

	var expressions sqlparser.SelectExprs
	for _, col := range n.columns {
		expressions = append(expressions, &sqlparser.AliasedExpr{Expr: col})
	}

	var tableNames []string
	for name := range tableNameMap {
		tableNames = append(tableNames, name)
	}
	sort.Strings(tableNames)

	return &route{
		eroute: &engine.Route{
			Opcode:    n.routeOpCode,
			TableName: strings.Join(tableNames, ", "),
			Keyspace:  n.keyspace,
			Vindex:    singleColumn,
			Values:    n.vindexValues,
		},
		Select: &sqlparser.Select{
			SelectExprs: expressions,
			From:        tablesForSelect,
			Where:       where,
		},
		tables: n.solved,
	}, nil
}

func relToTableExpr(t relation) sqlparser.TableExpr {
	switch t := t.(type) {
	case *routeTable:
		return &sqlparser.AliasedTableExpr{
			Expr: sqlparser.TableName{
				Name: t.vtable.Name,
			},
			Partitions: nil,
			As:         t.qtable.Alias.As,
			Hints:      nil,
		}
	case parenTables:
		tables := sqlparser.TableExprs{}
		for _, t := range t {
			tableExpr := relToTableExpr(t)
			tables = append(tables, tableExpr)
		}
		return &sqlparser.ParenTableExpr{Exprs: tables}
	case *leJoin:
		return &sqlparser.JoinTableExpr{
			LeftExpr:  relToTableExpr(t.a),
			Join:      sqlparser.NormalJoinType,
			RightExpr: relToTableExpr(t.b),
			Condition: sqlparser.JoinCondition{
				On: t.pred,
			},
		}
	default:
		panic("should never happen")
	}
}
