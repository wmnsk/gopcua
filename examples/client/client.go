// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

/*
Command client provides a connection establishment of OPC UA Secure Conversation.

XXX - Currently this command just initiates the connection(UACP) to the specified endpoint.
*/
package main

import (
	"encoding/hex"
	"flag"
	"log"

	"github.com/wmnsk/gopcua/uacp"
)

func main() {
	var (
		endpoint   = flag.String("endpoint", "opc.tcp://example.com/foo/bar", "OPC UA Endpoint URL")
		payloadHex = flag.String("payload", "deadbeef", "Payload to send in hex stream format")
	)
	flag.Parse()

	conn, err := uacp.Dial(*endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("Successfully established connection with %v", conn.RemoteEndpoint())

	payload, err := hex.DecodeString(*payloadHex)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Write(payload); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully sent message: %x", payload)
}
