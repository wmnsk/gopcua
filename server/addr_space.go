// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

//go:generate go run ../cmd/predefined-nodes/main.go

package server

import (
	"sync"
	"time"

	"github.com/gopcua/opcua/ua"
)

type Node interface {
	ID() *ua.NodeID
	Attribute(ua.AttributeID) (*AttrValue, error)
}

type AttrValue struct {
	Value           *ua.Variant
	SourceTimestamp time.Time
}

func NewAttrValue(v *ua.Variant) *AttrValue {
	return &AttrValue{Value: v}
}

type AddressSpace struct {
	mu    sync.RWMutex
	nodes map[string]Node
}

func newAddressSpace(nodes ...Node) *AddressSpace {
	return &AddressSpace{nodes: make(map[string]Node)}
}

func (a *AddressSpace) AddNodes(nodes ...Node) error {
	for _, n := range nodes {
		if err := a.AddNode(n); err != nil {
			return err
		}
	}
	return nil
}

func (a *AddressSpace) AddNode(n Node) error {
	a.mu.Lock()
	a.nodes[n.ID().String()] = n
	a.mu.Unlock()
	return nil
}

func (a *AddressSpace) Attribute(id *ua.NodeID, attr ua.AttributeID) (*AttrValue, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	n := a.nodes[id.String()]
	if n == nil {
		return nil, ua.StatusBadNodeIDUnknown
	}
	return n.Attribute(attr)
}
