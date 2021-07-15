// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ua

import (
	"reflect"
	"sync"

	"github.com/gopcua/opcua/errors"
)

// TypeRegistry provides a registry for Go types.
//
// Each type is registered with a unique identifier
// which cannot be changed for the lifetime of the component.
//
// Types can be registered multiple times under different
// identifiers.
//
// The implementation is safe for concurrent use.
type TypeRegistry struct {
	mu    sync.Mutex
	types map[string]reflect.Type
	ids   map[reflect.Type]string
}

// NewTypeRegistry returns a new type registry.
func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		types: make(map[string]reflect.Type),
		ids:   make(map[reflect.Type]string),
	}
}

// New returns a new instance of the type with the given id.
//
// If the id is not known the function returns nil.
//
// New panics if id is nil.
func (r *TypeRegistry) New(id *NodeID) interface{} {
	if id == nil {
		panic("opcua: missing id in call to TypeRegistry.New")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	typ, ok := r.types[id.String()]
	if !ok {
		return nil
	}
	return reflect.New(typ.Elem()).Interface()
}

// Lookup returns the id of the type of v or nil if
// the type is not registered.
//
// If the type was registered multiple times the first
// registered id for this type is returned.
func (r *TypeRegistry) Lookup(v interface{}) *NodeID {
	r.mu.Lock()
	defer r.mu.Unlock()
	if id, ok := r.ids[reflect.TypeOf(v)]; ok {
		return MustParseNodeID(id)
	}
	return nil
}

// Register adds a new type to the registry.
//
// If the id is already registered the function returns an error.
//
// Register panics if id is nil.
func (r *TypeRegistry) Register(id *NodeID, v interface{}) error {
	if id == nil {
		panic("opcua: missing id in call to TypeRegistry.Register")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	typ := reflect.TypeOf(v)
	ids := id.String()

	if r.types[ids] != nil {
		return errors.Errorf("%s is already registered", id)
	}
	r.types[ids] = typ

	if _, exists := r.ids[typ]; !exists {
		r.ids[typ] = ids
	}
	return nil
}
