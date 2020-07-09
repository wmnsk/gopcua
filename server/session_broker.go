package server

import (
	mrand "math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gopcua/opcua/debug"
	"github.com/gopcua/opcua/ua"
)

type session struct {
	cfg sessionConfig

	ID                *ua.NodeID
	AuthTokenID       *ua.NodeID
	serverNonce       []byte
	remoteCertificate []byte
}

type sessionConfig struct {
	sessionTimeout time.Duration
}

type sessionBroker struct {
	wg sync.WaitGroup

	// mu protects concurrent modification of s
	mu sync.Mutex
	// s is a slice of all sessions watched by the session broker
	s map[string]*session
}

func newSessionBroker() *sessionBroker {
	return &sessionBroker{
		s: make(map[string]*session),
	}
}

func (sb *sessionBroker) NewSession() *session {
	s := &session{
		ID:          ua.NewGUIDNodeID(1, uuid.New().String()),
		AuthTokenID: ua.NewNumericNodeID(0, uint32(mrand.Int31())),
	}

	sb.mu.Lock()
	sb.s[s.AuthTokenID.String()] = s
	sb.mu.Unlock()

	return s
}

func (sb *sessionBroker) Close(authToken *ua.NodeID) error {
	s, ok := sb.s[authToken.String()]
	if !ok {
		debug.Printf("sessionBroker.Close: error looking up session %v", authToken)
	}

	sb.mu.Lock()
	delete(sb.s, s.AuthTokenID.String())
	sb.mu.Unlock()

	return nil
}

func (sb *sessionBroker) Session(authToken *ua.NodeID) *session {
	s, ok := sb.s[authToken.String()]
	if !ok {
		debug.Printf("sessionBroker.Session: error looking up session %v", authToken)
	}

	return s
}
