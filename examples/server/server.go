// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gopcua/opcua"
)

func main() {
	endpoint := flag.String("endpoint", "opc.tcp://localhost:4840", "OPC UA Endpoint URL")
	flag.Parse()

	ctx := context.Background()

	s := opcua.NewServer(*endpoint)
	h := opcua.HandlerFunc(func(w opcua.ResponseWriter, r *opcua.Request) {
		fmt.Println(r)
	})

	log.Println("Listening on ", *endpoint)
	err := s.ListenAndServe(ctx, h)

	log.Print(err)
}
