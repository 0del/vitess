// Code generated by localvtctldclient-generator. DO NOT EDIT.

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

package localvtctldclient

import (
	"context"

	"google.golang.org/grpc"

	vtctldatapb "vitess.io/vitess/go/vt/proto/vtctldata"
)

// AddCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) AddCellInfo(ctx context.Context, in *vtctldatapb.AddCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.AddCellInfoResponse, error) {
	return client.s.AddCellInfo(ctx, in)
}

// AddCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) AddCellsAlias(ctx context.Context, in *vtctldatapb.AddCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.AddCellsAliasResponse, error) {
	return client.s.AddCellsAlias(ctx, in)
}

// ApplyRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ApplyRoutingRules(ctx context.Context, in *vtctldatapb.ApplyRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplyRoutingRulesResponse, error) {
	return client.s.ApplyRoutingRules(ctx, in)
}

// ApplySchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ApplySchema(ctx context.Context, in *vtctldatapb.ApplySchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplySchemaResponse, error) {
	return client.s.ApplySchema(ctx, in)
}

// ApplyVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ApplyVSchema(ctx context.Context, in *vtctldatapb.ApplyVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ApplyVSchemaResponse, error) {
	return client.s.ApplyVSchema(ctx, in)
}

// ChangeTabletType is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ChangeTabletType(ctx context.Context, in *vtctldatapb.ChangeTabletTypeRequest, opts ...grpc.CallOption) (*vtctldatapb.ChangeTabletTypeResponse, error) {
	return client.s.ChangeTabletType(ctx, in)
}

// CreateKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) CreateKeyspace(ctx context.Context, in *vtctldatapb.CreateKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.CreateKeyspaceResponse, error) {
	return client.s.CreateKeyspace(ctx, in)
}

// CreateShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) CreateShard(ctx context.Context, in *vtctldatapb.CreateShardRequest, opts ...grpc.CallOption) (*vtctldatapb.CreateShardResponse, error) {
	return client.s.CreateShard(ctx, in)
}

// DeleteCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteCellInfo(ctx context.Context, in *vtctldatapb.DeleteCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteCellInfoResponse, error) {
	return client.s.DeleteCellInfo(ctx, in)
}

// DeleteCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteCellsAlias(ctx context.Context, in *vtctldatapb.DeleteCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteCellsAliasResponse, error) {
	return client.s.DeleteCellsAlias(ctx, in)
}

// DeleteKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteKeyspace(ctx context.Context, in *vtctldatapb.DeleteKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteKeyspaceResponse, error) {
	return client.s.DeleteKeyspace(ctx, in)
}

// DeleteShards is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteShards(ctx context.Context, in *vtctldatapb.DeleteShardsRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteShardsResponse, error) {
	return client.s.DeleteShards(ctx, in)
}

// DeleteSrvVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteSrvVSchema(ctx context.Context, in *vtctldatapb.DeleteSrvVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteSrvVSchemaResponse, error) {
	return client.s.DeleteSrvVSchema(ctx, in)
}

// DeleteTablets is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) DeleteTablets(ctx context.Context, in *vtctldatapb.DeleteTabletsRequest, opts ...grpc.CallOption) (*vtctldatapb.DeleteTabletsResponse, error) {
	return client.s.DeleteTablets(ctx, in)
}

// EmergencyReparentShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) EmergencyReparentShard(ctx context.Context, in *vtctldatapb.EmergencyReparentShardRequest, opts ...grpc.CallOption) (*vtctldatapb.EmergencyReparentShardResponse, error) {
	return client.s.EmergencyReparentShard(ctx, in)
}

// ExecuteHook is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ExecuteHook(ctx context.Context, in *vtctldatapb.ExecuteHookRequest, opts ...grpc.CallOption) (*vtctldatapb.ExecuteHookResponse, error) {
	return client.s.ExecuteHook(ctx, in)
}

// FindAllShardsInKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) FindAllShardsInKeyspace(ctx context.Context, in *vtctldatapb.FindAllShardsInKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.FindAllShardsInKeyspaceResponse, error) {
	return client.s.FindAllShardsInKeyspace(ctx, in)
}

// GetBackups is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetBackups(ctx context.Context, in *vtctldatapb.GetBackupsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetBackupsResponse, error) {
	return client.s.GetBackups(ctx, in)
}

// GetCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetCellInfo(ctx context.Context, in *vtctldatapb.GetCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellInfoResponse, error) {
	return client.s.GetCellInfo(ctx, in)
}

// GetCellInfoNames is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetCellInfoNames(ctx context.Context, in *vtctldatapb.GetCellInfoNamesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellInfoNamesResponse, error) {
	return client.s.GetCellInfoNames(ctx, in)
}

// GetCellsAliases is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetCellsAliases(ctx context.Context, in *vtctldatapb.GetCellsAliasesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetCellsAliasesResponse, error) {
	return client.s.GetCellsAliases(ctx, in)
}

// GetKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetKeyspace(ctx context.Context, in *vtctldatapb.GetKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.GetKeyspaceResponse, error) {
	return client.s.GetKeyspace(ctx, in)
}

// GetKeyspaces is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetKeyspaces(ctx context.Context, in *vtctldatapb.GetKeyspacesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetKeyspacesResponse, error) {
	return client.s.GetKeyspaces(ctx, in)
}

// GetRoutingRules is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetRoutingRules(ctx context.Context, in *vtctldatapb.GetRoutingRulesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetRoutingRulesResponse, error) {
	return client.s.GetRoutingRules(ctx, in)
}

// GetSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSchema(ctx context.Context, in *vtctldatapb.GetSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSchemaResponse, error) {
	return client.s.GetSchema(ctx, in)
}

// GetShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetShard(ctx context.Context, in *vtctldatapb.GetShardRequest, opts ...grpc.CallOption) (*vtctldatapb.GetShardResponse, error) {
	return client.s.GetShard(ctx, in)
}

// GetSrvKeyspaceNames is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSrvKeyspaceNames(ctx context.Context, in *vtctldatapb.GetSrvKeyspaceNamesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvKeyspaceNamesResponse, error) {
	return client.s.GetSrvKeyspaceNames(ctx, in)
}

// GetSrvKeyspaces is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSrvKeyspaces(ctx context.Context, in *vtctldatapb.GetSrvKeyspacesRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvKeyspacesResponse, error) {
	return client.s.GetSrvKeyspaces(ctx, in)
}

// GetSrvVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSrvVSchema(ctx context.Context, in *vtctldatapb.GetSrvVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvVSchemaResponse, error) {
	return client.s.GetSrvVSchema(ctx, in)
}

// GetSrvVSchemas is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetSrvVSchemas(ctx context.Context, in *vtctldatapb.GetSrvVSchemasRequest, opts ...grpc.CallOption) (*vtctldatapb.GetSrvVSchemasResponse, error) {
	return client.s.GetSrvVSchemas(ctx, in)
}

// GetTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetTablet(ctx context.Context, in *vtctldatapb.GetTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTabletResponse, error) {
	return client.s.GetTablet(ctx, in)
}

// GetTablets is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetTablets(ctx context.Context, in *vtctldatapb.GetTabletsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetTabletsResponse, error) {
	return client.s.GetTablets(ctx, in)
}

// GetVSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetVSchema(ctx context.Context, in *vtctldatapb.GetVSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.GetVSchemaResponse, error) {
	return client.s.GetVSchema(ctx, in)
}

// GetWorkflows is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) GetWorkflows(ctx context.Context, in *vtctldatapb.GetWorkflowsRequest, opts ...grpc.CallOption) (*vtctldatapb.GetWorkflowsResponse, error) {
	return client.s.GetWorkflows(ctx, in)
}

// InitShardPrimary is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) InitShardPrimary(ctx context.Context, in *vtctldatapb.InitShardPrimaryRequest, opts ...grpc.CallOption) (*vtctldatapb.InitShardPrimaryResponse, error) {
	return client.s.InitShardPrimary(ctx, in)
}

// PingTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) PingTablet(ctx context.Context, in *vtctldatapb.PingTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.PingTabletResponse, error) {
	return client.s.PingTablet(ctx, in)
}

// PlannedReparentShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) PlannedReparentShard(ctx context.Context, in *vtctldatapb.PlannedReparentShardRequest, opts ...grpc.CallOption) (*vtctldatapb.PlannedReparentShardResponse, error) {
	return client.s.PlannedReparentShard(ctx, in)
}

// RebuildKeyspaceGraph is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RebuildKeyspaceGraph(ctx context.Context, in *vtctldatapb.RebuildKeyspaceGraphRequest, opts ...grpc.CallOption) (*vtctldatapb.RebuildKeyspaceGraphResponse, error) {
	return client.s.RebuildKeyspaceGraph(ctx, in)
}

// RebuildVSchemaGraph is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RebuildVSchemaGraph(ctx context.Context, in *vtctldatapb.RebuildVSchemaGraphRequest, opts ...grpc.CallOption) (*vtctldatapb.RebuildVSchemaGraphResponse, error) {
	return client.s.RebuildVSchemaGraph(ctx, in)
}

// RefreshState is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RefreshState(ctx context.Context, in *vtctldatapb.RefreshStateRequest, opts ...grpc.CallOption) (*vtctldatapb.RefreshStateResponse, error) {
	return client.s.RefreshState(ctx, in)
}

// RefreshStateByShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RefreshStateByShard(ctx context.Context, in *vtctldatapb.RefreshStateByShardRequest, opts ...grpc.CallOption) (*vtctldatapb.RefreshStateByShardResponse, error) {
	return client.s.RefreshStateByShard(ctx, in)
}

// ReloadSchema is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ReloadSchema(ctx context.Context, in *vtctldatapb.ReloadSchemaRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaResponse, error) {
	return client.s.ReloadSchema(ctx, in)
}

// ReloadSchemaKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ReloadSchemaKeyspace(ctx context.Context, in *vtctldatapb.ReloadSchemaKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaKeyspaceResponse, error) {
	return client.s.ReloadSchemaKeyspace(ctx, in)
}

// ReloadSchemaShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ReloadSchemaShard(ctx context.Context, in *vtctldatapb.ReloadSchemaShardRequest, opts ...grpc.CallOption) (*vtctldatapb.ReloadSchemaShardResponse, error) {
	return client.s.ReloadSchemaShard(ctx, in)
}

// RemoveKeyspaceCell is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RemoveKeyspaceCell(ctx context.Context, in *vtctldatapb.RemoveKeyspaceCellRequest, opts ...grpc.CallOption) (*vtctldatapb.RemoveKeyspaceCellResponse, error) {
	return client.s.RemoveKeyspaceCell(ctx, in)
}

// RemoveShardCell is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RemoveShardCell(ctx context.Context, in *vtctldatapb.RemoveShardCellRequest, opts ...grpc.CallOption) (*vtctldatapb.RemoveShardCellResponse, error) {
	return client.s.RemoveShardCell(ctx, in)
}

// ReparentTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ReparentTablet(ctx context.Context, in *vtctldatapb.ReparentTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.ReparentTabletResponse, error) {
	return client.s.ReparentTablet(ctx, in)
}

// RunHealthCheck is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) RunHealthCheck(ctx context.Context, in *vtctldatapb.RunHealthCheckRequest, opts ...grpc.CallOption) (*vtctldatapb.RunHealthCheckResponse, error) {
	return client.s.RunHealthCheck(ctx, in)
}

// SetKeyspaceServedFrom is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetKeyspaceServedFrom(ctx context.Context, in *vtctldatapb.SetKeyspaceServedFromRequest, opts ...grpc.CallOption) (*vtctldatapb.SetKeyspaceServedFromResponse, error) {
	return client.s.SetKeyspaceServedFrom(ctx, in)
}

// SetKeyspaceShardingInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetKeyspaceShardingInfo(ctx context.Context, in *vtctldatapb.SetKeyspaceShardingInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.SetKeyspaceShardingInfoResponse, error) {
	return client.s.SetKeyspaceShardingInfo(ctx, in)
}

// SetShardIsPrimaryServing is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetShardIsPrimaryServing(ctx context.Context, in *vtctldatapb.SetShardIsPrimaryServingRequest, opts ...grpc.CallOption) (*vtctldatapb.SetShardIsPrimaryServingResponse, error) {
	return client.s.SetShardIsPrimaryServing(ctx, in)
}

// SetShardTabletControl is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetShardTabletControl(ctx context.Context, in *vtctldatapb.SetShardTabletControlRequest, opts ...grpc.CallOption) (*vtctldatapb.SetShardTabletControlResponse, error) {
	return client.s.SetShardTabletControl(ctx, in)
}

// SetWritable is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SetWritable(ctx context.Context, in *vtctldatapb.SetWritableRequest, opts ...grpc.CallOption) (*vtctldatapb.SetWritableResponse, error) {
	return client.s.SetWritable(ctx, in)
}

// ShardReplicationPositions is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ShardReplicationPositions(ctx context.Context, in *vtctldatapb.ShardReplicationPositionsRequest, opts ...grpc.CallOption) (*vtctldatapb.ShardReplicationPositionsResponse, error) {
	return client.s.ShardReplicationPositions(ctx, in)
}

// SleepTablet is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) SleepTablet(ctx context.Context, in *vtctldatapb.SleepTabletRequest, opts ...grpc.CallOption) (*vtctldatapb.SleepTabletResponse, error) {
	return client.s.SleepTablet(ctx, in)
}

// StartReplication is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) StartReplication(ctx context.Context, in *vtctldatapb.StartReplicationRequest, opts ...grpc.CallOption) (*vtctldatapb.StartReplicationResponse, error) {
	return client.s.StartReplication(ctx, in)
}

// StopReplication is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) StopReplication(ctx context.Context, in *vtctldatapb.StopReplicationRequest, opts ...grpc.CallOption) (*vtctldatapb.StopReplicationResponse, error) {
	return client.s.StopReplication(ctx, in)
}

// TabletExternallyReparented is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) TabletExternallyReparented(ctx context.Context, in *vtctldatapb.TabletExternallyReparentedRequest, opts ...grpc.CallOption) (*vtctldatapb.TabletExternallyReparentedResponse, error) {
	return client.s.TabletExternallyReparented(ctx, in)
}

// UpdateCellInfo is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) UpdateCellInfo(ctx context.Context, in *vtctldatapb.UpdateCellInfoRequest, opts ...grpc.CallOption) (*vtctldatapb.UpdateCellInfoResponse, error) {
	return client.s.UpdateCellInfo(ctx, in)
}

// UpdateCellsAlias is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) UpdateCellsAlias(ctx context.Context, in *vtctldatapb.UpdateCellsAliasRequest, opts ...grpc.CallOption) (*vtctldatapb.UpdateCellsAliasResponse, error) {
	return client.s.UpdateCellsAlias(ctx, in)
}

// Validate is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) Validate(ctx context.Context, in *vtctldatapb.ValidateRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateResponse, error) {
	return client.s.Validate(ctx, in)
}

// ValidateKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateKeyspace(ctx context.Context, in *vtctldatapb.ValidateKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateKeyspaceResponse, error) {
	return client.s.ValidateKeyspace(ctx, in)
}

// ValidateSchemaKeyspace is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateSchemaKeyspace(ctx context.Context, in *vtctldatapb.ValidateSchemaKeyspaceRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateSchemaKeyspaceResponse, error) {
	return client.s.ValidateSchemaKeyspace(ctx, in)
}

// ValidateShard is part of the vtctlservicepb.VtctldClient interface.
func (client *localVtctldClient) ValidateShard(ctx context.Context, in *vtctldatapb.ValidateShardRequest, opts ...grpc.CallOption) (*vtctldatapb.ValidateShardResponse, error) {
	return client.s.ValidateShard(ctx, in)
}
