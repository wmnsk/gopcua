// +build integration

package uatest

import (
	"context"
	"testing"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
	"github.com/pascaldekloe/goe/verify"
)

func TestCallMethod(t *testing.T) {
	tests := []struct {
		req *ua.CallMethodRequest
		out []*ua.Variant
	}{
		{
			req: &ua.CallMethodRequest{
				ObjectID: ua.NewStringNodeID(2, "main"),
				MethodID: ua.NewStringNodeID(2, "even"),
				InputArguments: []*ua.Variant{
					ua.MustVariant(int64(12)),
				},
			},
			out: []*ua.Variant{ua.MustVariant(true)},
		},
	}

	srv := NewPythonServer("method_server.py")
	defer srv.Close()

	c := opcua.NewClient(srv.Endpoint, srv.Opts...)
	if err := c.Connect(context.Background()); err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	for _, tt := range tests {
		t.Run(tt.req.ObjectID.String(), func(t *testing.T) {
			resp, err := c.Call(tt.req)
			if err != nil {
				t.Fatal(err)
			}
			if got, want := resp.StatusCode, ua.StatusOK; got != want {
				t.Fatalf("got status %v want %v", got, want)
			}
			if got, want := resp.OutputArguments, tt.out; !verify.Values(t, "", got, want) {
				t.Fail()
			}
		})
	}
}
