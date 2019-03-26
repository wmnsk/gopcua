// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package opcua

import (
	"context"

	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
	"github.com/gopcua/opcua/uacp"
	"github.com/gopcua/opcua/uasc"
)

var DefaultAck = &uacp.Acknowledge{
	ReceiveBufSize: 0xffff,
	SendBufSize:    0xffff,
	MaxMessageSize: 1 << 16,
	MaxChunkCount:  256,
}

// Server is a high-level OPC-UA Server
type Server struct {
	endpointURL string
	Handler     Handler
	Ack         *uacp.Acknowledge

	l *uacp.Listener

	// cfg is the configuration for the secure channel.
	cfg *uasc.Config

	// sessionCfg is the configuration for the session.
	sessionCfg *uasc.SessionConfig
}

func NewServer(endpoint string, opts ...Option) *Server {
	cfg, sessionCfg := ApplyConfig(DefaultServerConfig(), DefaultServerSessionConfig(), opts...)
	return &Server{
		endpointURL: endpoint,
		cfg:         cfg,
		sessionCfg:  sessionCfg,
	}
}

func (s *Server) ListenAndServe(ctx context.Context, h Handler) error {
	if h == nil {
		h = s.Handler
	}
	ack := s.Ack
	if ack == nil {
		ack = DefaultAck
	}
	l, err := uacp.Listen(s.endpointURL, ack)
	if err != nil {
		return err
	}
	s.l = l
	s.cfg = DefaultServerConfig()
	return s.serve(ctx, h)
}

func (s *Server) serve(ctx context.Context, h Handler) error {
	for {
		// establish uacp conn
		c, err := s.l.Accept(ctx)
		if err != nil {
			return err
		}

		// establish secure channel
		sechan, err := uasc.NewSecureChannel(s.endpointURL, c, s.cfg)
		if err != nil {
			_ = c.Close()
			return err
		}

		// establish session
		go s.handle(ctx, c.ID(), sechan)
	}
}

func (s *Server) handle(ctx context.Context, connID uint32, sechan *uasc.SecureChannel) {
	for {
		msg := sechan.Receive(ctx)
		if msg.Err != nil {
			debug.Printf("conn %d: recv %#v", connID, msg.Err)
			_ = sechan.Close()
			return
		}
		debug.Printf("conn %d: recv %#v", connID, msg)

		switch req := msg.V.(type) {
		case *ua.FindServersRequest:
			s.handleFindServers(ctx, connID, sechan, req)
		case *ua.GetEndpointsRequest:
			s.handleGetEndpoints(ctx, connID, sechan, req)
		default:
			debug.Printf("conn %d: cannot handle %T", connID, req)
		}
	}
}

func (s *Server) handleFindServers(ctx context.Context, connID uint32, sechan *uasc.SecureChannel, req *ua.FindServersRequest) {
	debug.Printf("conn %d: handle %T", connID, req)
	resp := &ua.FindServersResponse{
		Servers: []*ua.ApplicationDescription{
			&ua.ApplicationDescription{
				ApplicationURI:  s.sessionCfg.ClientDescription.ApplicationURI,
				ProductURI:      s.sessionCfg.ClientDescription.ProductURI,
				ApplicationName: s.sessionCfg.ClientDescription.ApplicationName,
				ApplicationType: s.sessionCfg.ClientDescription.ApplicationType,
				// GatewayServerURI    string
				// DiscoveryProfileURI string
				DiscoveryURLs: []string{s.endpointURL},
			},
		},
	}

	sechan.SendResponse(req, resp)
}

func (s *Server) handleGetEndpoints(ctx context.Context, connID uint32, sechan *uasc.SecureChannel, req *ua.GetEndpointsRequest) {
	debug.Printf("conn %d: handle %T", connID, req)
	resp := &ua.GetEndpointsResponse{
		Endpoints: []*ua.EndpointDescription{
			&ua.EndpointDescription{
				EndpointURL: s.endpointURL,
				Server: &ua.ApplicationDescription{
					ApplicationURI:  s.sessionCfg.ClientDescription.ApplicationURI,
					ProductURI:      s.sessionCfg.ClientDescription.ProductURI,
					ApplicationName: s.sessionCfg.ClientDescription.ApplicationName,
					ApplicationType: s.sessionCfg.ClientDescription.ApplicationType,
					// GatewayServerURI    string
					// DiscoveryProfileURI string
					DiscoveryURLs: []string{s.endpointURL},
				},
				ServerCertificate:  nil,
				SecurityMode:       s.cfg.SecurityMode,
				SecurityPolicyURI:  s.cfg.SecurityPolicyURI,
				UserIdentityTokens: []*ua.UserTokenPolicy{},
				// TransportProfileURI string
				// SecurityLevel: s.cfg.SecurityLevel,
			},
		},
	}

	sechan.SendResponse(req, resp)
}

func (s *Server) Close() error {
	return s.l.Close()
}

type Handler interface {
	ServeOPCUA(w ResponseWriter, r *Request)
}

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeOPCUA(w ResponseWriter, r *Request) {
	f(w, r)
}

type Request struct {
	Msg *uasc.Message
}

type ResponseWriter struct {
	Msg *uasc.Message
}

func (w *ResponseWriter) Send(m *uasc.Message) {
	w.Msg = m
}
