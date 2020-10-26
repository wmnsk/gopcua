// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package uacp

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/pascaldekloe/goe/verify"
)

func TestConn(t *testing.T) {
	t.Run("server exists ", func(t *testing.T) {
		ep := "opc.tcp://127.0.0.1:4840/foo/bar"
		ln, err := Listen(context.Background(), ep, nil)
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		done := make(chan struct{})
		acceptErr := make(chan error, 1)
		go func() {
			c, err := ln.Accept(ctx)
			if err != nil {
				acceptErr <- err
				return
			}
			defer c.Close()
			close(done)
		}()

		if _, err = Dial(ctx, ep); err != nil {
			t.Error(err)
		}

		select {
		case <-done:
		case err := <-acceptErr:
			t.Fatalf("accept fail: %v", err)
		case <-time.After(time.Second):
			t.Fatal("timed out")
		}
	})

	t.Run("Address resolves, but does not implement a opcua-server", func(t *testing.T) {
		ep := "opc.tcp://example.com:56789"

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, err := Dial(ctx, ep)
		if !err.(*net.OpError).Timeout() {
			t.Error(err)
		}
	})
}

func TestClientWrite(t *testing.T) {
	ep := "opc.tcp://127.0.0.1:4840/foo/bar"
	ln, err := Listen(context.Background(), ep, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srvConn *Conn
	done := make(chan int)
	acceptErr := make(chan error, 1)
	go func() {
		defer ln.Close()
		var err error
		srvConn, err = ln.Accept(ctx)
		if err != nil {
			acceptErr <- err
			return
		}
		done <- 0
	}()

	cliConn, err := Dial(ctx, ep)
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case _, ok := <-done:
			if !ok {
				t.Fatal("failed to setup secure channel")
			}
			goto NEXT
		case err := <-acceptErr:
			t.Fatalf("accept fail: %v", err)
		case <-time.After(time.Second):
			t.Fatal("timed out")
		}
	}

NEXT:
	msg := &Message{Data: []byte{0xde, 0xad, 0xbe, 0xef}}
	if err := cliConn.Send("MSGF", msg); err != nil {
		t.Fatal(err)
	}

	got, err := srvConn.Receive()
	if err != nil {
		t.Fatal(err)
	}
	got = got[hdrlen:]

	want, err := msg.Encode()
	if err != nil {
		t.Fatal(err)
	}
	verify.Values(t, "", got, want)
}

func TestServerWrite(t *testing.T) {
	ep := "opc.tcp://127.0.0.1:4840/foo/bar"
	ln, err := Listen(context.Background(), ep, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srvConn *Conn
	done := make(chan int)
	acceptErr := make(chan error, 1)
	go func() {
		defer ln.Close()
		var err error
		srvConn, err = ln.Accept(ctx)
		if err != nil {
			acceptErr <- err
			return
		}
		done <- 0
	}()

	cliConn, err := Dial(ctx, ep)
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case _, ok := <-done:
			if !ok {
				t.Fatal("failed to setup secure channel")
			}
			goto NEXT
		case err := <-acceptErr:
			t.Fatalf("accept fail: %v", err)
		case <-time.After(time.Second):
			t.Fatal("timed out")
		}
	}

NEXT:
	want := []byte{0xde, 0xad, 0xbe, 0xef}
	if _, err := srvConn.Write(want); err != nil {
		t.Fatal(err)
	}

	got := make([]byte, cliConn.ReceiveBufSize())
	n, err := cliConn.Read(got)
	if err != nil {
		t.Fatal(err)
	}
	got = got[:n]
	verify.Values(t, "", got, want)
}
