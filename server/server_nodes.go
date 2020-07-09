package server

import (
	"time"

	"github.com/gopcua/opcua/ua"
)

type currentTime struct{}

func (n *currentTime) ID() *ua.NodeID {
	return ua.NewNumericNodeID(0, 2258)
}

func (n *currentTime) Attribute(attr ua.AttributeID) (*AttrValue, error) {
	switch attr {
	case ua.AttributeIDBrowseName:
		return NewAttrValue(ua.MustVariant("CurrentTime")), nil
	case ua.AttributeIDValue:
		return NewAttrValue(ua.MustVariant(time.Now())), nil
	default:
		return nil, ua.StatusBadAttributeIDInvalid
	}
}

type namespaces struct {
	s *Server
}

func (n *namespaces) ID() *ua.NodeID {
	return ua.NewNumericNodeID(0, 2255)
}

func (n *namespaces) Attribute(attr ua.AttributeID) (*AttrValue, error) {
	switch attr {
	case ua.AttributeIDBrowseName:
		return NewAttrValue(ua.MustVariant("Namespaces")), nil
	case ua.AttributeIDValue:
		return NewAttrValue(ua.MustVariant(n.s.Namespaces())), nil
	default:
		return nil, ua.StatusBadAttributeIDInvalid
	}

}

type serverStatus struct {
	s *Server
}

func (n *serverStatus) ID() *ua.NodeID {
	return ua.NewNumericNodeID(0, 2256)
}

func (n *serverStatus) Attribute(attr ua.AttributeID) (*AttrValue, error) {
	switch attr {
	case ua.AttributeIDBrowseName:
		return NewAttrValue(ua.MustVariant("ServerStatus")), nil
	case ua.AttributeIDValue:
		return NewAttrValue(ua.MustVariant(ua.NewExtensionObject(n.s.Status()))), nil
	default:
		return nil, ua.StatusBadAttributeIDInvalid
	}
}
