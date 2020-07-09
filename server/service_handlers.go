// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package server

import (
	"crypto/rand"
	"log"
	"time"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uasc"
)

func (s *Server) initHandlers() {
	s.serviceHandlers = make(map[uint16]handlerFunc)
	// s.registerHandler(id.ServiceFault_Encoding_DefaultBinary, handleServiceFault)
	s.registerHandler(id.FindServersRequest_Encoding_DefaultBinary, handleFindServersRequest)
	s.registerHandler(id.FindServersOnNetworkRequest_Encoding_DefaultBinary, handleFindServersOnNetworkRequest)
	s.registerHandler(id.GetEndpointsRequest_Encoding_DefaultBinary, handleGetEndpointsRequest)
	s.registerHandler(id.RegisterServerRequest_Encoding_DefaultBinary, handleRegisterServerRequest)
	s.registerHandler(id.RegisterServer2Request_Encoding_DefaultBinary, handleRegisterServer2Request)
	//s.registerHandler(id.OpenSecureChannelRequest_Encoding_DefaultBinary, handleOpenSecureChannelRequest)
	//s.registerHandler(id.CloseSecureChannelRequest_Encoding_DefaultBinary, handleCloseSecureChannelRequest)
	s.registerHandler(id.CreateSessionRequest_Encoding_DefaultBinary, handleCreateSessionRequest)
	s.registerHandler(id.ActivateSessionRequest_Encoding_DefaultBinary, handleActivateSessionRequest)
	s.registerHandler(id.CloseSessionRequest_Encoding_DefaultBinary, handleCloseSessionRequest)
	s.registerHandler(id.CancelRequest_Encoding_DefaultBinary, handleCancelRequest)
	s.registerHandler(id.AddNodesRequest_Encoding_DefaultBinary, handleAddNodesRequest)
	s.registerHandler(id.AddReferencesRequest_Encoding_DefaultBinary, handleAddReferencesRequest)
	s.registerHandler(id.DeleteNodesRequest_Encoding_DefaultBinary, handleDeleteNodesRequest)
	s.registerHandler(id.DeleteReferencesRequest_Encoding_DefaultBinary, handleDeleteReferencesRequest)
	s.registerHandler(id.BrowseRequest_Encoding_DefaultBinary, handleBrowseRequest)
	s.registerHandler(id.BrowseNextRequest_Encoding_DefaultBinary, handleBrowseNextRequest)
	s.registerHandler(id.TranslateBrowsePathsToNodeIDsRequest_Encoding_DefaultBinary, handleTranslateBrowsePathsToNodeIDsRequest)
	s.registerHandler(id.RegisterNodesRequest_Encoding_DefaultBinary, handleRegisterNodesRequest)
	s.registerHandler(id.UnregisterNodesRequest_Encoding_DefaultBinary, handleUnregisterNodesRequest)
	s.registerHandler(id.QueryFirstRequest_Encoding_DefaultBinary, handleQueryFirstRequest)
	s.registerHandler(id.QueryNextRequest_Encoding_DefaultBinary, handleQueryNextRequest)
	s.registerHandler(id.ReadRequest_Encoding_DefaultBinary, handleReadRequest)
	s.registerHandler(id.HistoryReadRequest_Encoding_DefaultBinary, handleHistoryReadRequest)
	s.registerHandler(id.WriteRequest_Encoding_DefaultBinary, handleWriteRequest)
	s.registerHandler(id.HistoryUpdateRequest_Encoding_DefaultBinary, handleHistoryUpdateRequest)
	s.registerHandler(id.CallMethodRequest_Encoding_DefaultBinary, handleCallMethodRequest)
	s.registerHandler(id.CallRequest_Encoding_DefaultBinary, handleCallRequest)
	s.registerHandler(id.MonitoredItemCreateRequest_Encoding_DefaultBinary, handleMonitoredItemCreateRequest)
	s.registerHandler(id.CreateMonitoredItemsRequest_Encoding_DefaultBinary, handleCreateMonitoredItemsRequest)
	s.registerHandler(id.MonitoredItemModifyRequest_Encoding_DefaultBinary, handleMonitoredItemModifyRequest)
	s.registerHandler(id.ModifyMonitoredItemsRequest_Encoding_DefaultBinary, handleModifyMonitoredItemsRequest)
	s.registerHandler(id.SetMonitoringModeRequest_Encoding_DefaultBinary, handleSetMonitoringModeRequest)
	s.registerHandler(id.SetTriggeringRequest_Encoding_DefaultBinary, handleSetTriggeringRequest)
	s.registerHandler(id.DeleteMonitoredItemsRequest_Encoding_DefaultBinary, handleDeleteMonitoredItemsRequest)
	s.registerHandler(id.CreateSubscriptionRequest_Encoding_DefaultBinary, handleCreateSubscriptionRequest)
	s.registerHandler(id.ModifySubscriptionRequest_Encoding_DefaultBinary, handleModifySubscriptionRequest)
	s.registerHandler(id.SetPublishingModeRequest_Encoding_DefaultBinary, handleSetPublishingModeRequest)
	s.registerHandler(id.PublishRequest_Encoding_DefaultBinary, handlePublishRequest)
	s.registerHandler(id.RepublishRequest_Encoding_DefaultBinary, handleRepublishRequest)
	s.registerHandler(id.TransferSubscriptionsRequest_Encoding_DefaultBinary, handleTransferSubscriptionsRequest)
	s.registerHandler(id.DeleteSubscriptionsRequest_Encoding_DefaultBinary, handleDeleteSubscriptionsRequest)
}

type response struct {
	hdr *ua.ResponseHeader
}

func (r *response) Header() *ua.ResponseHeader {
	return r.hdr
}

func (r *response) SetHeader(hdr *ua.ResponseHeader) {
	r.hdr = hdr
}

type handlerFunc func(*Server, *uasc.SecureChannel, ua.Request) (ua.Response, error)

func (s *Server) registerHandler(typeID uint16, f handlerFunc) {
	s.serviceHandlers[typeID] = f
}

func (s *Server) handleService(sc *uasc.SecureChannel, req ua.Request) {
	debug.Printf("handleService: Got: %T\n", req)

	var resp ua.Response
	var err error

	typeID := ua.ServiceTypeID(req)
	h, ok := s.serviceHandlers[typeID]
	if ok {
		resp, err = h(s, sc, req)
	} else {
		if typeID == 0 {
			debug.Printf("unknown service %T. Did you call register?", req)
		}
		err = ua.StatusBadServiceUnsupported
	}

	// todo(dh): Find the actual request ID for the response headers
	if err != nil {
		if statusCode, ok := err.(ua.StatusCode); ok {
			resp = &ua.ServiceFault{ResponseHeader: responseHeader(0, statusCode)}
		} else {
			resp = &ua.ServiceFault{ResponseHeader: responseHeader(0, ua.StatusBadUnexpectedError)}
		}
	}

	if resp == nil {
		return
	}

	err = sc.SendResponse(resp)
	if err != nil {
		debug.Printf("Error sending response: %s\n", err)
	}
}

func responseHeader(reqID uint32, statusCode ua.StatusCode) *ua.ResponseHeader {
	return &ua.ResponseHeader{
		Timestamp:          time.Now(),
		RequestHandle:      reqID,
		ServiceResult:      statusCode,
		ServiceDiagnostics: &ua.DiagnosticInfo{},
		StringTable:        []string{},
		AdditionalHeader:   ua.NewExtensionObject(nil),
	}
}

// func handleServiceFault(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
// 	debug.Printf("Handling %T\n", r)

// 	req, ok := r.(*ua.ServiceFault)
// 	if !ok {
// 		debug.Printf("handleServiceFault: Expected *ua.ServiceFault, got %T", r)
// 		return nil, ua.StatusBadRequestTypeInvalid
// 	}
// 	debug.Printf("Got ServiceFault: %s", req.ResponseHeader.ServiceResult)

// 	// No response required
// 	return nil, nil
// }

func handleFindServersRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.FindServersRequest)
	if !ok {
		debug.Printf("handleFindServersRequest: Expected *ua.FindServersRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	response := &ua.FindServersResponse{
		ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
		Servers: []*ua.ApplicationDescription{
			s.Endpoints[0].Server,
		},
	}

	return response, nil
}

func handleFindServersOnNetworkRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.FindServersOnNetworkRequest)
	if !ok {
		debug.Printf("handleFindServersOnNetworkRequest: Expected *ua.FindServersOnNetworkRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.FindServersOnNetworkResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleGetEndpointsRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.GetEndpointsRequest)
	if !ok {
		debug.Printf("handleGetEndpointsRequest: Expected *ua.GetEndpointsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	response := &ua.GetEndpointsResponse{
		ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
		Endpoints:      s.Endpoints,
	}

	return response, nil
}

func handleRegisterServerRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.RegisterServerRequest)
	if !ok {
		debug.Printf("handleRegisterServerRequest: Expected *ua.RegisterServerRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.RegisterServerResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleRegisterServer2Request(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.RegisterServer2Request)
	if !ok {
		debug.Printf("handleRegisterServer2Request: Expected *ua.RegisterServer2Request, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.RegisterServer2Response{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

// Handled in the UASC stack
//func handleOpenSecureChannelRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
//	debug.Printf("Handling %T\n", r)
//
//	req, ok := r.(*ua.OpenSecureChannelRequest)
//	if !ok {
//		debug.Printf("handleOpenSecureChannelRequest: Expected *ua.OpenSecureChannelRequest, got %T", r)
//		return nil, ua.StatusBadRequestTypeInvalid
//	}
//
//	return response, nil
//}

// Handled in the UASC stack
//func handleCloseSecureChannelRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
//	debug.Printf("Handling %T\n", r)
//
//	_, ok := r.(*ua.CloseSecureChannelRequest)
//	if !ok {
//		debug.Printf("handleCloseSecureChannelRequest: Expected *ua.CloseSecureChannelRequest, got %T", r)
//		return nil, ua.StatusBadRequestTypeInvalid
//	}
//
//	sc.Close()
//
//	// No response required
//	return nil, nil
//}

const (
	sessionTimeoutMin     = 100            // 100ms
	sessionTimeoutMax     = 30 * 60 * 1000 // 30 minutes
	sessionTimeoutDefault = 60 * 1000      // 60s

	sessionNonceLength = 32
)

func handleCreateSessionRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CreateSessionRequest)
	if !ok {
		debug.Printf("handleCreateSessionRequest: Expected *ua.CreateSessionRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	// New session
	sess := s.sb.NewSession()

	// Ensure session timeout is reasonable
	sess.cfg.sessionTimeout = time.Duration(req.RequestedSessionTimeout) * time.Millisecond
	if sess.cfg.sessionTimeout > sessionTimeoutMax || sess.cfg.sessionTimeout < sessionTimeoutMin {
		sess.cfg.sessionTimeout = sessionTimeoutDefault
	}

	nonce := make([]byte, sessionNonceLength)
	if _, err := rand.Read(nonce); err != nil {
		log.Printf("error creating session nonce")
		return nil, ua.StatusBadInternalError
	}
	sess.serverNonce = nonce
	sess.remoteCertificate = req.ClientCertificate

	sig, alg, err := sc.NewSessionSignature(req.ClientCertificate, req.ClientNonce)
	if err != nil {
		log.Printf("error creating session signature")
		return nil, ua.StatusBadInternalError
	}

	response := &ua.CreateSessionResponse{
		ResponseHeader:        responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
		SessionID:             sess.ID,
		AuthenticationToken:   sess.AuthTokenID,
		RevisedSessionTimeout: float64(sess.cfg.sessionTimeout / time.Millisecond),
		MaxRequestMessageSize: 0, // Not used
		ServerSignature: &ua.SignatureData{
			Signature: sig,
			Algorithm: alg,
		},
		ServerCertificate: s.cfg.certificate,
		ServerNonce:       nonce,
		ServerEndpoints:   s.Endpoints,
	}

	return response, nil
}

func handleActivateSessionRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.ActivateSessionRequest)
	if !ok {
		debug.Printf("handleActivateSessionRequest: Expected *ua.ActivateSessionRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	sess := s.sb.Session(req.RequestHeader.AuthenticationToken)
	if sess == nil {
		return nil, ua.StatusBadSessionIDInvalid
	}

	err := sc.VerifySessionSignature(sess.remoteCertificate, sess.serverNonce, req.ClientSignature.Signature)
	if err != nil {
		debug.Printf("error verifying session signature with nonce: %s", err)
		return nil, ua.StatusBadSecurityChecksFailed
	}

	nonce := make([]byte, sessionNonceLength)
	if _, err := rand.Read(nonce); err != nil {
		log.Printf("error creating session nonce")
		return nil, ua.StatusBadInternalError
	}
	sess.serverNonce = nonce

	response := &ua.ActivateSessionResponse{
		ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
		ServerNonce:    nonce,
		// Results:         []ua.StatusCode{},
		// DiagnosticInfos: []*ua.DiagnosticInfo{},
	}

	return response, nil
}

func handleCloseSessionRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CloseSessionRequest)
	if !ok {
		debug.Printf("handleCloseSessionRequest: Expected *ua.CloseSessionRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	err := s.sb.Close(req.RequestHeader.AuthenticationToken)
	if err != nil {
		return nil, ua.StatusBadSessionIDInvalid
	}

	//TODO: deal with 'delete subscriptions' field in request
	response := &ua.CloseSessionResponse{
		ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	}

	return response, nil
}

func handleCancelRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CancelRequest)
	if !ok {
		debug.Printf("handleCancelRequest: Expected *ua.CancelRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.CancelResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleAddNodesRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.AddNodesRequest)
	if !ok {
		debug.Printf("handleAddNodesRequest: Expected *ua.AddNodesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.AddNodesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleAddReferencesRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.AddReferencesRequest)
	if !ok {
		debug.Printf("handleAddReferencesRequest: Expected *ua.AddReferencesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.AddReferencesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleDeleteNodesRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.DeleteNodesRequest)
	if !ok {
		debug.Printf("handleDeleteNodesRequest: Expected *ua.DeleteNodesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.DeleteNodesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleDeleteReferencesRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.DeleteReferencesRequest)
	if !ok {
		debug.Printf("handleDeleteReferencesRequest: Expected *ua.DeleteReferencesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.DeleteReferencesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleBrowseRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.BrowseRequest)
	if !ok {
		debug.Printf("handleBrowseRequest: Expected *ua.BrowseRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.BrowseResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleBrowseNextRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.BrowseNextRequest)
	if !ok {
		debug.Printf("handleBrowseNextRequest: Expected *ua.BrowseNextRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.BrowseNextResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleTranslateBrowsePathsToNodeIDsRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.TranslateBrowsePathsToNodeIDsRequest)
	if !ok {
		debug.Printf("handleTranslateBrowsePathsToNodeIDsRequest: Expected *ua.TranslateBrowsePathsToNodeIDsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.TranslateBrowsePathsToNodeIDsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleRegisterNodesRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.RegisterNodesRequest)
	if !ok {
		debug.Printf("handleRegisterNodesRequest: Expected *ua.RegisterNodesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.RegisterNodesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleUnregisterNodesRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.UnregisterNodesRequest)
	if !ok {
		debug.Printf("handleUnregisterNodesRequest: Expected *ua.UnregisterNodesRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.UnregisterNodesResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleQueryFirstRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.QueryFirstRequest)
	if !ok {
		debug.Printf("handleQueryFirstRequest: Expected *ua.QueryFirstRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.QueryFirstResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleQueryNextRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.QueryNextRequest)
	if !ok {
		debug.Printf("handleQueryNextRequest: Expected *ua.QueryNextRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.QueryNextResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleReadRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.ReadRequest)
	if !ok {
		debug.Printf("handleReadRequest: Expected *ua.ReadRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	results := make([]*ua.DataValue, len(req.NodesToRead))
	for i, n := range req.NodesToRead {
		debug.Printf("read: node=%s attr=%s", n.NodeID, n.AttributeID)

		dv := &ua.DataValue{
			EncodingMask:    ua.DataValueServerTimestamp,
			ServerTimestamp: time.Now(),
		}

		v, err := s.as.Attribute(n.NodeID, n.AttributeID)
		switch x := err.(type) {
		case nil:
			dv.EncodingMask |= ua.DataValueStatusCode | ua.DataValueValue
			dv.Status = ua.StatusOK
			dv.Value = v.Value

		case ua.StatusCode:
			dv.EncodingMask |= ua.DataValueStatusCode
			dv.Status = x

		default:
			debug.Printf("read: node=%s attr=%s err=%s", n.NodeID, n.AttributeID, err)
			dv.EncodingMask |= ua.DataValueStatusCode
			dv.Status = ua.StatusBadInternalError
		}

		if v != nil && !v.SourceTimestamp.IsZero() {
			dv.EncodingMask |= ua.DataValueSourceTimestamp
			dv.SourceTimestamp = v.SourceTimestamp
		}

		results[i] = dv
	}

	response := &ua.ReadResponse{
		ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
		Results:        results,
	}

	return response, nil
}

func handleHistoryReadRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.HistoryReadRequest)
	if !ok {
		debug.Printf("handleHistoryReadRequest: Expected *ua.HistoryReadRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.HistoryReadResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleWriteRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.WriteRequest)
	if !ok {
		debug.Printf("handleWriteRequest: Expected *ua.WriteRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.WriteResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleHistoryUpdateRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.HistoryUpdateRequest)
	if !ok {
		debug.Printf("handleHistoryUpdateRequest: Expected *ua.HistoryUpdateRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.HistoryUpdateResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleCallMethodRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	//req, ok := r.(*ua.CallMethodRequest)
	//if !ok {
	//	debug.Printf("handleCallMethodRequest: Expected *ua.CallMethodRequest, got %T", r)
	//	return nil, ua.StatusBadRequestTypeInvalid
	//}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(0, ua.StatusBadServiceUnsupported)}
	// response := &ua.CallMethodResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleCallRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CallRequest)
	if !ok {
		debug.Printf("handleCallRequest: Expected *ua.CallRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.CallResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleMonitoredItemCreateRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	//req, ok := r.(*ua.MonitoredItemCreateRequest)
	//if !ok {
	//	debug.Printf("handleMonitoredItemCreateRequest: Expected *ua.MonitoredItemCreateRequest, got %T", r)
	//	return nil, ua.StatusBadRequestTypeInvalid
	//}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(0, ua.StatusBadServiceUnsupported)}
	// response := &ua.MonitoredItemCreateResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleCreateMonitoredItemsRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CreateMonitoredItemsRequest)
	if !ok {
		debug.Printf("handleCreateMonitoredItemsRequest: Expected *ua.CreateMonitoredItemsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.CreateMonitoredItemsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleMonitoredItemModifyRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	//req, ok := r.(*ua.MonitoredItemModifyRequest)
	//if !ok {
	//	debug.Printf("handleMonitoredItemModifyRequest: Expected *ua.MonitoredItemModifyRequest, got %T", r)
	//	return nil, ua.StatusBadRequestTypeInvalid
	//}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(0, ua.StatusBadServiceUnsupported)}
	// response := &ua.MonitoredItemModifyResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleModifyMonitoredItemsRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.ModifyMonitoredItemsRequest)
	if !ok {
		debug.Printf("handleModifyMonitoredItemsRequest: Expected *ua.ModifyMonitoredItemsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.ModifyMonitoredItemsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleSetMonitoringModeRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.SetMonitoringModeRequest)
	if !ok {
		debug.Printf("handleSetMonitoringModeRequest: Expected *ua.SetMonitoringModeRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.SetMonitoringModeResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleSetTriggeringRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.SetTriggeringRequest)
	if !ok {
		debug.Printf("handleSetTriggeringRequest: Expected *ua.SetTriggeringRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.SetTriggeringResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleDeleteMonitoredItemsRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.DeleteMonitoredItemsRequest)
	if !ok {
		debug.Printf("handleDeleteMonitoredItemsRequest: Expected *ua.DeleteMonitoredItemsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.DeleteMonitoredItemsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleCreateSubscriptionRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.CreateSubscriptionRequest)
	if !ok {
		debug.Printf("handleCreateSubscriptionRequest: Expected *ua.CreateSubscriptionRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.CreateSubscriptionResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleModifySubscriptionRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.ModifySubscriptionRequest)
	if !ok {
		debug.Printf("handleModifySubscriptionRequest: Expected *ua.ModifySubscriptionRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.ModifySubscriptionResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleSetPublishingModeRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.SetPublishingModeRequest)
	if !ok {
		debug.Printf("handleSetPublishingModeRequest: Expected *ua.SetPublishingModeRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.SetPublishingModeResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handlePublishRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.PublishRequest)
	if !ok {
		debug.Printf("handlePublishRequest: Expected *ua.PublishRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.PublishResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleRepublishRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.RepublishRequest)
	if !ok {
		debug.Printf("handleRepublishRequest: Expected *ua.RepublishRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.RepublishResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleTransferSubscriptionsRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.TransferSubscriptionsRequest)
	if !ok {
		debug.Printf("handleTransferSubscriptionsRequest: Expected *ua.TransferSubscriptionsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.TransferSubscriptionsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}

func handleDeleteSubscriptionsRequest(s *Server, sc *uasc.SecureChannel, r ua.Request) (ua.Response, error) {
	debug.Printf("Handling %T\n", r)

	req, ok := r.(*ua.DeleteSubscriptionsRequest)
	if !ok {
		debug.Printf("handleDeleteSubscriptionsRequest: Expected *ua.DeleteSubscriptionsRequest, got %T", r)
		return nil, ua.StatusBadRequestTypeInvalid
	}

	//TODO: Replace with proper response once implemented
	response := &ua.ServiceFault{ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusBadServiceUnsupported)}
	// response := &ua.DeleteSubscriptionsResponse{
	//	ResponseHeader: responseHeader(req.RequestHeader.RequestHandle, ua.StatusOK),
	//  ... remaining fields
	//}

	return response, nil
}
