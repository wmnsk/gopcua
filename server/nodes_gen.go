// Generated code. DO NOT EDIT

// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package server

import 	"github.com/gopcua/opcua/ua"

type node struct {
	id   *ua.NodeID
	attr map[ua.AttributeID]*AttrValue

	superTypeID *ua.NodeID
}

func (n *node) ID() *ua.NodeID {
	return n.id
}

func (n *node) Attribute(id ua.AttributeID) (*AttrValue, error) {
	if n.attr == nil {
		return nil, ua.StatusBadAttributeIDInvalid
	}
	v := n.attr[id]
	if v == nil {
		return nil, ua.StatusBadAttributeIDInvalid
	}
	return v, nil
}

func PredefinedNodes() []Node{
	return []Node{
		&node{
			id: ua.NewNumericNodeID(0, 3062),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3063),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 1),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Boolean")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Boolean")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SByte")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SByte")),
			},
			superTypeID: ua.NewNumericNodeID(0, 27),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Byte")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Byte")),
			},
			superTypeID: ua.NewNumericNodeID(0, 28),
		},
		&node{
			id: ua.NewNumericNodeID(0, 4),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Int16")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Int16")),
			},
			superTypeID: ua.NewNumericNodeID(0, 27),
		},
		&node{
			id: ua.NewNumericNodeID(0, 5),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UInt16")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UInt16")),
			},
			superTypeID: ua.NewNumericNodeID(0, 28),
		},
		&node{
			id: ua.NewNumericNodeID(0, 6),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Int32")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Int32")),
			},
			superTypeID: ua.NewNumericNodeID(0, 27),
		},
		&node{
			id: ua.NewNumericNodeID(0, 7),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UInt32")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UInt32")),
			},
			superTypeID: ua.NewNumericNodeID(0, 28),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Int64")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Int64")),
			},
			superTypeID: ua.NewNumericNodeID(0, 27),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UInt64")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UInt64")),
			},
			superTypeID: ua.NewNumericNodeID(0, 28),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Float")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Float")),
			},
			superTypeID: ua.NewNumericNodeID(0, 26),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Double")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Double")),
			},
			superTypeID: ua.NewNumericNodeID(0, 26),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("String")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("String")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 13),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DateTime")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DateTime")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Guid")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Guid")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ByteString")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ByteString")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("XmlElement")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("XmlElement")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NodeId")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NodeId")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExpandedNodeId")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExpandedNodeId")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StatusCode")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StatusCode")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 20),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("QualifiedName")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("QualifiedName")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("LocalizedText")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("LocalizedText")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataValue")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataValue")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 25),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DiagnosticInfo")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DiagnosticInfo")),
			},
			superTypeID: ua.NewNumericNodeID(0, 24),
		},
		&node{
			id: ua.NewNumericNodeID(0, 50),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Decimal")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Decimal")),
			},
			superTypeID: ua.NewNumericNodeID(0, 26),
		},
		&node{
			id: ua.NewNumericNodeID(0, 35),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Organizes")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Organizes")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"OrganizedBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 36),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasEventSource")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasEventSource")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"EventSourceOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 37),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasModellingRule")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasModellingRule")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"ModellingRuleOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 38),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasEncoding")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasEncoding")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"EncodingOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 39),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasDescription")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasDescription")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"DescriptionOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 40),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasTypeDefinition")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasTypeDefinition")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"TypeDefinitionOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 41),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("GeneratesEvent")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("GeneratesEvent")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"GeneratedBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3065),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AlwaysGeneratesEvent")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AlwaysGeneratesEvent")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"AlwaysGeneratedBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 41),
		},
		&node{
			id: ua.NewNumericNodeID(0, 45),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasSubtype")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasSubtype")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"SubtypeOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 34),
		},
		&node{
			id: ua.NewNumericNodeID(0, 46),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasProperty")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasProperty")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"PropertyOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 44),
		},
		&node{
			id: ua.NewNumericNodeID(0, 47),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasComponent")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasComponent")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"ComponentOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 44),
		},
		&node{
			id: ua.NewNumericNodeID(0, 48),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasNotifier")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasNotifier")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"NotifierOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 36),
		},
		&node{
			id: ua.NewNumericNodeID(0, 49),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasOrderedComponent")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasOrderedComponent")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"OrderedComponentOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 51),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FromState")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FromState")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"ToTransition"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 52),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ToState")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ToState")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"FromTransition"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 53),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasCause")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasCause")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"MayBeCausedBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 54),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasEffect")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasEffect")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"MayBeEffectedBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 117),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasSubStateMachine")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasSubStateMachine")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"SubStateMachineOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 56),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasHistoricalConfiguration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasHistoricalConfiguration")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"HistoricalConfigurationOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 44),
		},
		&node{
			id: ua.NewNumericNodeID(0, 58),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BaseObjectType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BaseObjectType")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 61),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FolderType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FolderType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 63),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BaseDataVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BaseDataVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 62),
		},
		&node{
			id: ua.NewNumericNodeID(0, 68),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PropertyType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PropertyType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 62),
		},
		&node{
			id: ua.NewNumericNodeID(0, 69),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataTypeDescriptionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataTypeDescriptionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 72),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataTypeDictionaryType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataTypeDictionaryType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 75),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataTypeSystemType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataTypeSystemType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 76),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataTypeEncodingType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataTypeEncodingType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 120),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NamingRuleType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NamingRuleType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 77),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ModellingRuleType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ModellingRuleType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 78),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Mandatory")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Mandatory")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 80),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Optional")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Optional")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 83),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExposesItsArray")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExposesItsArray")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11508),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OptionalPlaceholder")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OptionalPlaceholder")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11510),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MandatoryPlaceholder")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MandatoryPlaceholder")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 84),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Root")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Root")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 85),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Objects")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Objects")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 86),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Types")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Types")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 87),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Views")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Views")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 88),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ObjectTypes")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ObjectTypes")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 89),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("VariableTypes")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("VariableTypes")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 90),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataTypes")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataTypes")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 91),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ReferenceTypes")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ReferenceTypes")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 92),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("XML Schema")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("XML Schema")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 93),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OPC Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OPC Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 129),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasArgumentDescription")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasArgumentDescription")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"ArgumentDescriptionOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 131),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasOptionalInputArgumentDescription")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasOptionalInputArgumentDescription")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"OptionalInputArgumentDescriptionOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 129),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15957),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("http://opcfoundation.org/UA/")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("http://opcfoundation.org/UA/")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3068),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NodeVersion")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NodeVersion")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12170),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ViewVersion")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ViewVersion")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3067),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Icon")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Icon")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3069),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("LocalTime")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("LocalTime")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3070),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AllowNulls")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AllowNulls")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11433),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ValueAsText")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ValueAsText")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11498),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MaxStringLength")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MaxStringLength")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15002),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MaxCharacters")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MaxCharacters")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12908),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MaxByteStringLength")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MaxByteStringLength")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11512),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MaxArrayLength")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MaxArrayLength")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11513),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EngineeringUnits")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EngineeringUnits")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11432),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EnumStrings")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EnumStrings")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3071),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EnumValues")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EnumValues")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12745),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OptionSetValues")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OptionSetValues")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3072),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("InputArguments")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("InputArguments")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 3073),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OutputArguments")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OutputArguments")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16306),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DefaultInputValues")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DefaultInputValues")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17605),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DefaultInstanceBrowseName")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DefaultInstanceBrowseName")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2000),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ImageBMP")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ImageBMP")),
			},
			superTypeID: ua.NewNumericNodeID(0, 30),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2001),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ImageGIF")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ImageGIF")),
			},
			superTypeID: ua.NewNumericNodeID(0, 30),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2002),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ImageJPG")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ImageJPG")),
			},
			superTypeID: ua.NewNumericNodeID(0, 30),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2003),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ImagePNG")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ImagePNG")),
			},
			superTypeID: ua.NewNumericNodeID(0, 30),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16307),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AudioDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AudioDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2004),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2013),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerCapabilitiesType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerCapabilitiesType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2020),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerDiagnosticsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerDiagnosticsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2026),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionsDiagnosticsSummaryType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionsDiagnosticsSummaryType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2029),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionDiagnosticsObjectType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionDiagnosticsObjectType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2033),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("VendorServerInfoType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("VendorServerInfoType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2034),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerRedundancyType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerRedundancyType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2036),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TransparentRedundancyType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TransparentRedundancyType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2034),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2039),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NonTransparentRedundancyType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NonTransparentRedundancyType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2034),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11945),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NonTransparentNetworkRedundancyType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NonTransparentNetworkRedundancyType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2039),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11564),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OperationLimitsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OperationLimitsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11575),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FileType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FileType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11595),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AddressSpaceFileType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AddressSpaceFileType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 11575),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11616),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NamespaceMetadataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NamespaceMetadataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11645),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NamespacesType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NamespacesType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2340),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AggregateFunctionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AggregateFunctionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2138),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerStatusType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerStatusType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3051),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BuildInfoType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BuildInfoType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2150),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerDiagnosticsSummaryType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerDiagnosticsSummaryType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2164),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SamplingIntervalDiagnosticsArrayType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SamplingIntervalDiagnosticsArrayType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2165),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SamplingIntervalDiagnosticsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SamplingIntervalDiagnosticsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2171),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SubscriptionDiagnosticsArrayType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SubscriptionDiagnosticsArrayType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2172),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SubscriptionDiagnosticsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SubscriptionDiagnosticsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2196),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionDiagnosticsArrayType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionDiagnosticsArrayType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2197),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionDiagnosticsVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionDiagnosticsVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2243),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionSecurityDiagnosticsArrayType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionSecurityDiagnosticsArrayType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2244),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionSecurityDiagnosticsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionSecurityDiagnosticsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11487),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OptionSetType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OptionSetType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16309),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SelectionListType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SelectionListType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17986),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AudioVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AudioVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3048),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EventTypes")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EventTypes")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2253),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Server")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Server")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11192),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HistoryServerCapabilities")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HistoryServerCapabilities")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11737),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BitFieldMaskDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BitFieldMaskDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 9),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14533),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("KeyValuePair")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("KeyValuePair")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15528),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EndpointType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EndpointType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2299),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StateMachineType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StateMachineType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2755),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StateVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StateVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2762),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TransitionVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TransitionVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2760),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FiniteStateVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FiniteStateVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2755),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2767),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FiniteTransitionVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FiniteTransitionVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2762),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2307),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2309),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("InitialStateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("InitialStateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2307),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2310),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TransitionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TransitionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15109),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ChoiceStateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ChoiceStateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2307),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15112),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasGuard")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasGuard")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"GuardOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15113),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("GuardVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("GuardVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15128),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExpressionGuardVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExpressionGuardVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15113),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15317),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ElseGuardVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ElseGuardVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15113),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17709),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RationalNumberType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RationalNumberType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17716),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("3DVectorType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("3DVectorType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17714),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18774),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("3DCartesianCoordinatesType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("3DCartesianCoordinatesType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18772),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18781),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("3DOrientationType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("3DOrientationType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18779),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18791),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("3DFrameType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("3DFrameType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18786),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18806),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RationalNumber")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RationalNumber")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18808),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("3DVector")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("3DVector")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18807),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18810),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("3DCartesianCoordinates")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("3DCartesianCoordinates")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18809),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18812),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("3DOrientation")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("3DOrientation")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18811),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18814),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("3DFrame")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("3DFrame")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18813),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11939),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OpenFileMode")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OpenFileMode")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 13353),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FileDirectoryType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FileDirectoryType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16314),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FileSystem")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FileSystem")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15744),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TemporaryFileTransferType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TemporaryFileTransferType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15803),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FileTransferStateMachineType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FileTransferStateMachineType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2771),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15607),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RoleSetType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RoleSetType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15620),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RoleType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RoleType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15632),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("IdentityCriteriaType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("IdentityCriteriaType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15634),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("IdentityMappingRuleType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("IdentityMappingRuleType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15644),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Anonymous")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Anonymous")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15656),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuthenticatedUser")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuthenticatedUser")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15668),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Observer")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Observer")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15680),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Operator")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Operator")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16036),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Engineer")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Engineer")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15692),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Supervisor")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Supervisor")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15716),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ConfigureAdmin")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ConfigureAdmin")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15704),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SecurityAdmin")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SecurityAdmin")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17591),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DictionaryFolderType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DictionaryFolderType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17594),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Dictionaries")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Dictionaries")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17597),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasDictionaryEntry")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasDictionaryEntry")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"DictionaryEntryOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17598),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("IrdiDictionaryEntryType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("IrdiDictionaryEntryType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17600),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UriDictionaryEntryType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UriDictionaryEntryType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17708),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("InterfaceTypes")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("InterfaceTypes")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17603),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasInterface")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasInterface")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"InterfaceOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17604),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasAddIn")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasAddIn")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"AddInOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23498),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("CurrencyUnitType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("CurrencyUnitType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 23501),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("CurrencyUnit")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("CurrencyUnit")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2365),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataItemType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataItemType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15318),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BaseAnalogType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BaseAnalogType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2365),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2368),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AnalogItemType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AnalogItemType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15318),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17497),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AnalogUnitType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AnalogUnitType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15318),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17570),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AnalogUnitRangeType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AnalogUnitRangeType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2368),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2373),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TwoStateDiscreteType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TwoStateDiscreteType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2372),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2376),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MultiStateDiscreteType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MultiStateDiscreteType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2372),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11238),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MultiStateValueDiscreteType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MultiStateValueDiscreteType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2372),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12029),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("YArrayItemType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("YArrayItemType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12038),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("XYArrayItemType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("XYArrayItemType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12047),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ImageItemType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ImageItemType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12057),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("CubeItemType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("CubeItemType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12068),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NDimensionArrayItemType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NDimensionArrayItemType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12021),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8995),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TwoStateVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TwoStateVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2755),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9002),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ConditionVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ConditionVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9004),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasTrueSubState")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasTrueSubState")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"IsTrueSubStateOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9005),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasFalseSubState")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasFalseSubState")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"IsFalseSubStateOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16361),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasAlarmSuppressionGroup")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasAlarmSuppressionGroup")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"IsAlarmSuppressionGroupOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16362),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AlarmGroupMember")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AlarmGroupMember")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"MemberOfAlarmGroup"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 35),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2830),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DialogConditionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DialogConditionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2782),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2881),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AcknowledgeableConditionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AcknowledgeableConditionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2782),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2915),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AlarmConditionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AlarmConditionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2881),
		},
		&node{
			id: ua.NewNumericNodeID(0, 16405),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AlarmGroupType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AlarmGroupType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2929),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ShelvedStateMachineType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ShelvedStateMachineType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2771),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2955),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("LimitAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("LimitAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2915),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9318),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExclusiveLimitStateMachineType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExclusiveLimitStateMachineType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2771),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9341),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExclusiveLimitAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExclusiveLimitAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2955),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9906),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NonExclusiveLimitAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NonExclusiveLimitAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2955),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10060),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NonExclusiveLevelAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NonExclusiveLevelAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 9906),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9482),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExclusiveLevelAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExclusiveLevelAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 9341),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10368),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NonExclusiveDeviationAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NonExclusiveDeviationAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 9906),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10214),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NonExclusiveRateOfChangeAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NonExclusiveRateOfChangeAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 9906),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9764),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExclusiveDeviationAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExclusiveDeviationAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 9341),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9623),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExclusiveRateOfChangeAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExclusiveRateOfChangeAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 9341),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10523),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DiscreteAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DiscreteAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2915),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10637),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OffNormalAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OffNormalAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 10523),
		},
		&node{
			id: ua.NewNumericNodeID(0, 10751),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TripAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TripAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 10637),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18347),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("InstrumentDiagnosticAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("InstrumentDiagnosticAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 10637),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18496),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SystemDiagnosticAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SystemDiagnosticAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 10637),
		},
		&node{
			id: ua.NewNumericNodeID(0, 13225),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("CertificateExpirationAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("CertificateExpirationAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 11753),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17080),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DiscrepancyAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DiscrepancyAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2915),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2790),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2127),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2803),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionEnableEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionEnableEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2829),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionCommentEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionCommentEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8927),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionRespondEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionRespondEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8944),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionAcknowledgeEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionAcknowledgeEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8961),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionConfirmEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionConfirmEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11093),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionShelvingEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionShelvingEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17225),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionSuppressionEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionSuppressionEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17242),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionSilenceEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionSilenceEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15013),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionResetEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionResetEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17259),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuditConditionOutOfServiceEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuditConditionOutOfServiceEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2790),
		},
		&node{
			id: ua.NewNumericNodeID(0, 9006),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasCondition")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasCondition")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"IsConditionOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 32),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17276),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasEffectDisable")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasEffectDisable")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"MayBeDisabledBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 54),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17983),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasEffectEnable")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasEffectEnable")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"MayBeEnabledBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 54),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17984),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasEffectSuppressed")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasEffectSuppressed")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"MayBeSuppressedBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 54),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17985),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasEffectUnsuppressed")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasEffectUnsuppressed")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"MayBeUnsuppressedBy"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 54),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17279),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AlarmMetricsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AlarmMetricsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17277),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AlarmRateVariableType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AlarmRateVariableType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2391),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ProgramStateMachineType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ProgramStateMachineType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2771),
		},
		&node{
			id: ua.NewNumericNodeID(0, 3806),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ProgramTransitionAuditEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ProgramTransitionAuditEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 2315),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2380),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ProgramDiagnosticType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ProgramDiagnosticType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15383),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ProgramDiagnostic2Type")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ProgramDiagnostic2Type")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11214),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Annotations")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Annotations")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2318),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HistoricalDataConfigurationType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HistoricalDataConfigurationType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11202),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HA Configuration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HA Configuration")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11215),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HistoricalEventFilter")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HistoricalEventFilter")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2330),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HistoryServerCapabilitiesType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HistoryServerCapabilitiesType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12522),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TrustListType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TrustListType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 11575),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12552),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TrustListMasks")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TrustListMasks")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12554),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TrustListDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TrustListDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19297),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TrustListOutOfDateAlarmType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TrustListOutOfDateAlarmType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 11753),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12555),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("CertificateGroupType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("CertificateGroupType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 13813),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("CertificateGroupFolderType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("CertificateGroupFolderType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12558),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HttpsCertificateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HttpsCertificateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12556),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15181),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UserCredentialCertificateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UserCredentialCertificateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12556),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12559),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RsaMinApplicationCertificateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RsaMinApplicationCertificateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12557),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12560),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RsaSha256ApplicationCertificateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RsaSha256ApplicationCertificateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12557),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12581),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerConfigurationType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerConfigurationType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12637),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerConfiguration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerConfiguration")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17496),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("KeyCredentialConfigurationFolderType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("KeyCredentialConfigurationFolderType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18155),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("KeyCredentialConfiguration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("KeyCredentialConfiguration")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18001),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("KeyCredentialConfigurationType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("KeyCredentialConfigurationType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18029),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("KeyCredentialUpdatedAuditEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("KeyCredentialUpdatedAuditEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18011),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18047),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("KeyCredentialDeletedAuditEventType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("KeyCredentialDeletedAuditEventType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 18011),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17732),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuthorizationServices")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuthorizationServices")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17852),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AuthorizationServiceConfigurationType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AuthorizationServiceConfigurationType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11187),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AggregateConfigurationType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AggregateConfigurationType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 2341),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Interpolative")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Interpolative")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2342),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Average")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Average")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2343),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TimeAverage")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TimeAverage")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11285),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TimeAverage2")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TimeAverage2")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2344),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Total")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Total")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11304),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Total2")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Total2")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2346),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Minimum")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Minimum")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2347),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Maximum")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Maximum")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2348),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MinimumActualTime")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MinimumActualTime")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2349),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MaximumActualTime")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MaximumActualTime")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2350),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Range")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Range")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11286),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Minimum2")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Minimum2")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11287),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Maximum2")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Maximum2")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11305),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MinimumActualTime2")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MinimumActualTime2")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11306),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MaximumActualTime2")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MaximumActualTime2")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11288),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Range2")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Range2")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2351),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AnnotationCount")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AnnotationCount")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2352),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Count")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Count")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11307),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DurationInStateZero")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DurationInStateZero")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11308),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DurationInStateNonZero")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DurationInStateNonZero")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2355),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NumberOfTransitions")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NumberOfTransitions")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2357),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Start")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Start")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2358),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("End")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("End")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2359),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Delta")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Delta")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11505),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StartBound")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StartBound")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11506),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EndBound")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EndBound")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11507),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DeltaBounds")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DeltaBounds")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2360),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DurationGood")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DurationGood")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2361),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DurationBad")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DurationBad")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2362),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PercentGood")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PercentGood")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2363),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PercentBad")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PercentBad")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 2364),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("WorstQuality")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("WorstQuality")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11292),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("WorstQuality2")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("WorstQuality2")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11426),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StandardDeviationSample")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StandardDeviationSample")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11427),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StandardDeviationPopulation")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StandardDeviationPopulation")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11428),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("VarianceSample")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("VarianceSample")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11429),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("VariancePopulation")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("VariancePopulation")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15487),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StructureDescription")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StructureDescription")),
			},
			superTypeID: ua.NewNumericNodeID(0, 14525),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15488),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EnumDescription")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EnumDescription")),
			},
			superTypeID: ua.NewNumericNodeID(0, 14525),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15005),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SimpleTypeDescription")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SimpleTypeDescription")),
			},
			superTypeID: ua.NewNumericNodeID(0, 14525),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15006),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UABinaryFileDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UABinaryFileDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15534),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14647),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubState")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubState")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14523),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetMetaDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetMetaDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15534),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14524),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FieldMetaData")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FieldMetaData")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15904),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetFieldFlags")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetFieldFlags")),
			},
			superTypeID: ua.NewNumericNodeID(0, 5),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14593),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ConfigurationVersionDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ConfigurationVersionDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15578),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishedDataSetDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishedDataSetDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14273),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishedVariableDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishedVariableDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15581),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishedDataItemsDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishedDataItemsDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15580),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15582),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishedEventsDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishedEventsDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15580),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15583),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetFieldContentMask")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetFieldContentMask")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15597),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetWriterDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetWriterDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15480),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("WriterGroupDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("WriterGroupDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15609),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15617),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubConnectionDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubConnectionDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15510),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NetworkAddressUrlDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NetworkAddressUrlDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15502),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15520),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ReaderGroupDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ReaderGroupDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15609),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15623),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetReaderDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetReaderDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15631),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TargetVariablesDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TargetVariablesDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15630),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14744),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FieldTargetDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FieldTargetDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15874),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OverrideValueHandling")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OverrideValueHandling")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15635),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SubscribedDataSetMirrorDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SubscribedDataSetMirrorDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15630),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15530),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubConfigurationDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubConfigurationDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 20408),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetOrderingType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetOrderingType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15642),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UadpNetworkMessageContentMask")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UadpNetworkMessageContentMask")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15645),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UadpWriterGroupMessageDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UadpWriterGroupMessageDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15616),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15646),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UadpDataSetMessageContentMask")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UadpDataSetMessageContentMask")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15652),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UadpDataSetWriterMessageDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UadpDataSetWriterMessageDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15605),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15653),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UadpDataSetReaderMessageDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UadpDataSetReaderMessageDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15629),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15654),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("JsonNetworkMessageContentMask")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("JsonNetworkMessageContentMask")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15657),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("JsonWriterGroupMessageDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("JsonWriterGroupMessageDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15616),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15658),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("JsonDataSetMessageContentMask")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("JsonDataSetMessageContentMask")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15664),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("JsonDataSetWriterMessageDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("JsonDataSetWriterMessageDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15605),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15665),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("JsonDataSetReaderMessageDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("JsonDataSetReaderMessageDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15629),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17467),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DatagramConnectionTransportDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DatagramConnectionTransportDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15618),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15532),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DatagramWriterGroupTransportDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DatagramWriterGroupTransportDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15611),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15007),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerConnectionTransportDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerConnectionTransportDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15618),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15008),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerTransportQualityOfService")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerTransportQualityOfService")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15667),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerWriterGroupTransportDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerWriterGroupTransportDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15611),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15669),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerDataSetWriterTransportDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerDataSetWriterTransportDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15598),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15670),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerDataSetReaderTransportDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerDataSetReaderTransportDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15628),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15906),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubKeyServiceType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubKeyServiceType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15452),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SecurityGroupFolderType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SecurityGroupFolderType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15471),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SecurityGroupType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SecurityGroupType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14416),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishSubscribeType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishSubscribeType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15906),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14443),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishSubscribe")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishSubscribe")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14476),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasPubSubConnection")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasPubSubConnection")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"PubSubConnectionOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14509),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishedDataSetType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishedDataSetType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15489),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExtensionFieldsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExtensionFieldsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14936),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetToWriter")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetToWriter")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"WriterToDataSet"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 33),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14534),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishedDataItemsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishedDataItemsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 14509),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14572),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PublishedEventsType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PublishedEventsType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 14509),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14477),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetFolderType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetFolderType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 61),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14209),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubConnectionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubConnectionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17725),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("WriterGroupType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("WriterGroupType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 14232),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15296),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasDataSetWriter")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasDataSetWriter")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"IsWriterInGroup"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18804),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasWriterGroup")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasWriterGroup")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"IsWriterGroupOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17999),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ReaderGroupType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ReaderGroupType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 14232),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15297),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasDataSetReader")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasDataSetReader")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"IsReaderInGroup"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 18805),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ReferenceType_32")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HasReaderGroup")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HasReaderGroup")),
				ua.AttributeIDInverseName: NewAttrValue(ua.MustVariant(&ua.LocalizedText{Locale:"", Text:"IsReaderGroupOf"})),
			},
			superTypeID: ua.NewNumericNodeID(0, 47),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15298),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetWriterType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetWriterType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15306),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DataSetReaderType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DataSetReaderType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15108),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SubscribedDataSetType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SubscribedDataSetType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15111),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TargetVariablesType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TargetVariablesType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15108),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15127),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SubscribedDataSetMirrorType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SubscribedDataSetMirrorType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15108),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14643),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubStatusType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubStatusType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 58),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19723),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DiagnosticsLevel")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DiagnosticsLevel")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19725),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsCounterType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsCounterType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 63),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19730),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsCounterClassification")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsCounterClassification")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19732),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsRootType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsRootType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19786),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsConnectionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsConnectionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19834),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsWriterGroupType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsWriterGroupType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19903),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsReaderGroupType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsReaderGroupType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19968),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsDataSetWriterType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsDataSetWriterType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 20027),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsDataSetReaderType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PubSubDiagnosticsDataSetReaderType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 19677),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21105),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UadpWriterGroupMessageType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UadpWriterGroupMessageType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17998),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21111),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UadpDataSetWriterMessageType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UadpDataSetWriterMessageType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 21096),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21116),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UadpDataSetReaderMessageType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UadpDataSetReaderMessageType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 21104),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21126),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("JsonWriterGroupMessageType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("JsonWriterGroupMessageType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17998),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21128),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("JsonDataSetWriterMessageType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("JsonDataSetWriterMessageType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 21096),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21130),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("JsonDataSetReaderMessageType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("JsonDataSetReaderMessageType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 21104),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15064),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DatagramConnectionTransportType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DatagramConnectionTransportType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17721),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21133),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DatagramWriterGroupTransportType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DatagramWriterGroupTransportType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17997),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15155),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerConnectionTransportType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerConnectionTransportType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17721),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21136),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerWriterGroupTransportType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerWriterGroupTransportType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17997),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21138),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerDataSetWriterTransportType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerDataSetWriterTransportType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15305),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21142),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BrokerDataSetReaderTransportType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BrokerDataSetReaderTransportType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15319),
		},
		&node{
			id: ua.NewNumericNodeID(0, 21147),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("ObjectType_8")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NetworkAddressUrlType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NetworkAddressUrlType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 21145),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19077),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MultiStateDictionaryEntryDiscreteBaseType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MultiStateDictionaryEntryDiscreteBaseType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 11238),
		},
		&node{
			id: ua.NewNumericNodeID(0, 19084),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("VariableType_16")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MultiStateDictionaryEntryDiscreteType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MultiStateDictionaryEntryDiscreteType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 19077),
		},
		&node{
			id: ua.NewNumericNodeID(0, 256),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("IdType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("IdType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 257),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NodeClass")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NodeClass")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 94),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PermissionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PermissionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15031),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AccessLevelType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AccessLevelType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 3),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15406),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AccessLevelExType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AccessLevelExType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15033),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EventNotifierType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EventNotifierType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 3),
		},
		&node{
			id: ua.NewNumericNodeID(0, 95),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AccessRestrictionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AccessRestrictionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 96),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RolePermissionType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RolePermissionType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 98),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StructureType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StructureType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 101),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StructureField")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StructureField")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 99),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StructureDefinition")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StructureDefinition")),
			},
			superTypeID: ua.NewNumericNodeID(0, 97),
		},
		&node{
			id: ua.NewNumericNodeID(0, 100),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EnumDefinition")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EnumDefinition")),
			},
			superTypeID: ua.NewNumericNodeID(0, 97),
		},
		&node{
			id: ua.NewNumericNodeID(0, 296),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Argument")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Argument")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 7594),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EnumValueType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EnumValueType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 102),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EnumField")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EnumField")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7594),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12755),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("OptionSet")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("OptionSet")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12877),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NormalizedString")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NormalizedString")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12878),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DecimalString")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DecimalString")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12879),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DurationString")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DurationString")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12880),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TimeString")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TimeString")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12881),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DateString")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DateString")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 290),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Duration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Duration")),
			},
			superTypeID: ua.NewNumericNodeID(0, 11),
		},
		&node{
			id: ua.NewNumericNodeID(0, 294),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UtcTime")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UtcTime")),
			},
			superTypeID: ua.NewNumericNodeID(0, 13),
		},
		&node{
			id: ua.NewNumericNodeID(0, 295),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("LocaleId")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("LocaleId")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 8912),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("TimeZoneDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("TimeZoneDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 17588),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Index")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Index")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 288),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("IntegerId")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("IntegerId")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 307),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ApplicationType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ApplicationType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 308),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ApplicationDescription")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ApplicationDescription")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 20998),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("VersionTime")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("VersionTime")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12189),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerOnNetwork")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerOnNetwork")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 311),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ApplicationInstanceCertificate")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ApplicationInstanceCertificate")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15),
		},
		&node{
			id: ua.NewNumericNodeID(0, 302),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MessageSecurityMode")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MessageSecurityMode")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 303),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UserTokenType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UserTokenType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 304),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UserTokenPolicy")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UserTokenPolicy")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 312),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EndpointDescription")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EndpointDescription")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 432),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RegisteredServer")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RegisteredServer")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12890),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DiscoveryConfiguration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DiscoveryConfiguration")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12891),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MdnsDiscoveryConfiguration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MdnsDiscoveryConfiguration")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12890),
		},
		&node{
			id: ua.NewNumericNodeID(0, 315),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SecurityTokenRequestType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SecurityTokenRequestType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 344),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SignedSoftwareCertificate")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SignedSoftwareCertificate")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 388),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionAuthenticationToken")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionAuthenticationToken")),
			},
			superTypeID: ua.NewNumericNodeID(0, 17),
		},
		&node{
			id: ua.NewNumericNodeID(0, 319),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AnonymousIdentityToken")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AnonymousIdentityToken")),
			},
			superTypeID: ua.NewNumericNodeID(0, 316),
		},
		&node{
			id: ua.NewNumericNodeID(0, 322),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("UserNameIdentityToken")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("UserNameIdentityToken")),
			},
			superTypeID: ua.NewNumericNodeID(0, 316),
		},
		&node{
			id: ua.NewNumericNodeID(0, 325),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("X509IdentityToken")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("X509IdentityToken")),
			},
			superTypeID: ua.NewNumericNodeID(0, 316),
		},
		&node{
			id: ua.NewNumericNodeID(0, 938),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("IssuedIdentityToken")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("IssuedIdentityToken")),
			},
			superTypeID: ua.NewNumericNodeID(0, 316),
		},
		&node{
			id: ua.NewNumericNodeID(0, 348),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NodeAttributesMask")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NodeAttributesMask")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 376),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AddNodesItem")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AddNodesItem")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 379),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AddReferencesItem")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AddReferencesItem")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 382),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DeleteNodesItem")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DeleteNodesItem")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 385),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DeleteReferencesItem")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DeleteReferencesItem")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 347),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AttributeWriteMask")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AttributeWriteMask")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 521),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ContinuationPoint")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ContinuationPoint")),
			},
			superTypeID: ua.NewNumericNodeID(0, 15),
		},
		&node{
			id: ua.NewNumericNodeID(0, 537),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RelativePathElement")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RelativePathElement")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 540),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RelativePath")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RelativePath")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 289),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Counter")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Counter")),
			},
			superTypeID: ua.NewNumericNodeID(0, 7),
		},
		&node{
			id: ua.NewNumericNodeID(0, 291),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NumericRange")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NumericRange")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 292),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Time")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Time")),
			},
			superTypeID: ua.NewNumericNodeID(0, 12),
		},
		&node{
			id: ua.NewNumericNodeID(0, 293),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Date")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Date")),
			},
			superTypeID: ua.NewNumericNodeID(0, 13),
		},
		&node{
			id: ua.NewNumericNodeID(0, 331),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EndpointConfiguration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EndpointConfiguration")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 576),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FilterOperator")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FilterOperator")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 583),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ContentFilterElement")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ContentFilterElement")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 586),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ContentFilter")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ContentFilter")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 589),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("FilterOperand")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("FilterOperand")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 592),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ElementOperand")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ElementOperand")),
			},
			superTypeID: ua.NewNumericNodeID(0, 589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 595),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("LiteralOperand")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("LiteralOperand")),
			},
			superTypeID: ua.NewNumericNodeID(0, 589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 598),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AttributeOperand")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AttributeOperand")),
			},
			superTypeID: ua.NewNumericNodeID(0, 589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 601),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SimpleAttributeOperand")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SimpleAttributeOperand")),
			},
			superTypeID: ua.NewNumericNodeID(0, 589),
		},
		&node{
			id: ua.NewNumericNodeID(0, 659),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HistoryEvent")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HistoryEvent")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11234),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HistoryUpdateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HistoryUpdateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11293),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("PerformUpdateType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("PerformUpdateType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 719),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("MonitoringFilter")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("MonitoringFilter")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 725),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EventFilter")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EventFilter")),
			},
			superTypeID: ua.NewNumericNodeID(0, 719),
		},
		&node{
			id: ua.NewNumericNodeID(0, 948),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AggregateConfiguration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AggregateConfiguration")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 920),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("HistoryEventFieldList")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("HistoryEventFieldList")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 338),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("BuildInfo")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("BuildInfo")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 851),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RedundancySupport")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RedundancySupport")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 852),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerState")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerState")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 853),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("RedundantServerDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("RedundantServerDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11943),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EndpointUrlListDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EndpointUrlListDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 11944),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("NetworkGroupDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("NetworkGroupDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 856),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SamplingIntervalDiagnosticsDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SamplingIntervalDiagnosticsDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 859),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerDiagnosticsSummaryDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerDiagnosticsSummaryDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 862),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServerStatusDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServerStatusDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 865),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionDiagnosticsDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionDiagnosticsDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 868),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SessionSecurityDiagnosticsDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SessionSecurityDiagnosticsDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 871),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ServiceCounterDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ServiceCounterDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 299),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("StatusResult")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("StatusResult")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 874),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SubscriptionDiagnosticsDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SubscriptionDiagnosticsDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 877),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ModelChangeStructureDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ModelChangeStructureDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 897),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("SemanticChangeStructureDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("SemanticChangeStructureDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 884),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Range")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Range")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 887),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("EUInformation")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("EUInformation")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12077),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AxisScaleEnumeration")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AxisScaleEnumeration")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12171),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ComplexNumberType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ComplexNumberType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12172),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("DoubleComplexNumberType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("DoubleComplexNumberType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12079),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("AxisInformation")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("AxisInformation")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 12080),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("XVType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("XVType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 894),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ProgramDiagnosticDataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ProgramDiagnosticDataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 15396),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ProgramDiagnostic2DataType")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ProgramDiagnostic2DataType")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 891),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Annotation")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Annotation")),
			},
			superTypeID: ua.NewNumericNodeID(0, 22),
		},
		&node{
			id: ua.NewNumericNodeID(0, 890),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("DataType_64")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("ExceptionDeviationFormat")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("ExceptionDeviationFormat")),
			},
			superTypeID: ua.NewNumericNodeID(0, 29),
		},
		&node{
			id: ua.NewNumericNodeID(0, 14846),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15671),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18815),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18816),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18817),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18818),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18819),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18820),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18821),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18822),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18823),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15736),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23507),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12680),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15676),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 125),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 126),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 127),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15421),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15422),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 124),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14839),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14847),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15677),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15678),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14323),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15679),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15681),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15682),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15683),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15688),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15689),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21150),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15691),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15693),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15694),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15695),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21151),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21152),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21153),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15701),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15702),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15703),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15705),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15706),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15707),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15712),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14848),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15713),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21154),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15715),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15717),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15718),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15719),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15724),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15725),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17468),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21155),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15479),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15727),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15729),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15733),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 128),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 121),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14844),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 122),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 123),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 298),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8251),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14845),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12765),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12766),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8917),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 310),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12207),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 306),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 314),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 434),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12900),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12901),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 346),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 318),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 321),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 324),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 327),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 940),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 378),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 381),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 384),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 387),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 539),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 542),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 333),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 585),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 588),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 591),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 594),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 597),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 600),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 603),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 661),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 721),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 727),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 950),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 922),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 340),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 855),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11957),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11958),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 858),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 861),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 864),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 867),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 870),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 873),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 301),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 876),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 879),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 899),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 886),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 889),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12181),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12182),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12089),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12090),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 896),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15397),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 893),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default Binary")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default Binary")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 7617),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Opc.Ua")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Opc.Ua")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14802),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15949),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18851),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18852),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18853),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18854),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18855),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18856),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18857),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18858),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 18859),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15728),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23520),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12676),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15950),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14796),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15589),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15590),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15529),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15531),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14794),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14795),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14803),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15951),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15952),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14319),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15953),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15954),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15955),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15956),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15987),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15988),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21174),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15990),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15991),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15992),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15993),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21175),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21176),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21177),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15995),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15996),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16007),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16008),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16009),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16010),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16011),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14804),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16012),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21178),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16014),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16015),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16016),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16017),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16018),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16019),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17472),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21179),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15579),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16021),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16022),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16023),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16126),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14797),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14800),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14798),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14799),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 297),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 7616),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 14801),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12757),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12758),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8913),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 309),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12195),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 305),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 313),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 433),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12892),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12893),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 345),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 317),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 320),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 323),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 326),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 939),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 377),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 380),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 383),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 386),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 538),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 541),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 332),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 584),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 587),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 590),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 593),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 596),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 599),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 602),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 660),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 720),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 726),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 949),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 921),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 339),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 854),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11949),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 11950),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 857),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 860),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 863),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 866),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 869),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 872),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 300),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 875),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 878),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 898),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 885),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 888),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12173),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12174),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12081),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 12082),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 895),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15401),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 892),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default XML")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default XML")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 8252),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Variable_2")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Opc.Ua")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Opc.Ua")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15041),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16150),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19064),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19065),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19066),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19067),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19068),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19069),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19070),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19071),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 19072),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15042),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 23528),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15044),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16151),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15057),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15058),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15059),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15700),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15714),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15050),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15051),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15049),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16152),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16153),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15060),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16154),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16155),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16156),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16157),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16158),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16159),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21198),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16161),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16280),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16281),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16282),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21199),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21200),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21201),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16284),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16285),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16286),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16287),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16288),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16308),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16310),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15061),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16311),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21202),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16323),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16391),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16392),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16393),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16394),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16404),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 17476),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 21203),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15726),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16524),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16525),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 16526),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15062),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15063),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15065),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15066),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15067),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15081),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15082),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15083),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15084),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15085),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15086),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15087),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15095),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15098),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15099),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15102),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15105),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15106),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15136),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15140),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15141),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15142),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15143),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15144),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15165),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15169),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15172),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15175),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15188),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15189),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15199),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15204),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15205),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15206),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15207),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15208),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15209),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15210),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15273),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15293),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15295),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15304),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15349),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15361),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15362),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15363),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15364),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15365),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15366),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15367),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15368),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15369),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15370),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15371),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15372),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15373),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15374),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15375),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15376),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15377),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15378),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15379),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15380),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15381),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15405),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
		&node{
			id: ua.NewNumericNodeID(0, 15382),
			attr: map[ua.AttributeID]*AttrValue{
				ua.AttributeIDNodeClass: NewAttrValue(ua.MustVariant("Object_1")),
				ua.AttributeIDBrowseName: NewAttrValue(ua.MustVariant("Default JSON")),
				ua.AttributeIDDisplayName: NewAttrValue(ua.MustVariant("Default JSON")),
			},
		},
	}
}
