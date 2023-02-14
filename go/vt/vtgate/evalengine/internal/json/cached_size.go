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
// Code generated by Sizegen. DO NOT EDIT.

package json

import hack "vitess.io/vitess/go/hack"

func (cached *Object) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(24)
	}
	// field kvs []vitess.io/vitess/go/vt/vtgate/evalengine/internal/json.kv
	{
		size += hack.RuntimeAllocSize(int64(cap(cached.kvs)) * int64(24))
		for _, elem := range cached.kvs {
			size += elem.CachedSize(false)
		}
	}
	return size
}
func (cached *Value) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(80)
	}
	// field o vitess.io/vitess/go/vt/vtgate/evalengine/internal/json.Object
	size += cached.o.CachedSize(false)
	// field a []*vitess.io/vitess/go/vt/vtgate/evalengine/internal/json.Value
	{
		size += hack.RuntimeAllocSize(int64(cap(cached.a)) * int64(8))
		for _, elem := range cached.a {
			size += elem.CachedSize(true)
		}
	}
	// field s string
	size += hack.RuntimeAllocSize(int64(len(cached.s)))
	return size
}
func (cached *kv) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(24)
	}
	// field k string
	size += hack.RuntimeAllocSize(int64(len(cached.k)))
	// field v *vitess.io/vitess/go/vt/vtgate/evalengine/internal/json.Value
	size += cached.v.CachedSize(true)
	return size
}
