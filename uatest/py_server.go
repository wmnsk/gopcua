// +build integration

package uatest

import (
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/errors"
	"github.com/gopcua/opcua/ua"
)

type Server interface {
	URL() string
	Close() error
}

// PythonServer runs a python test server.
type PythonServer struct {
	// Path is the path to the Python server.
	Path string

	// Endpoint is the endpoint address which will be set
	// after the server has started.
	Endpoint string

	// Opts contains the client options required to connect to the server.
	// They are valid after the server has been started.
	Opts []opcua.Option

	cmd *exec.Cmd
}

// NewPythonServer creates a test server and starts it. The function
// panics if the server cannot be started.
func NewPythonServer(path string) *PythonServer {
	s := &PythonServer{Path: path}
	if err := s.Run(); err != nil {
		panic(err)
	}
	return s
}

func (s *PythonServer) URL() string {
	return s.Endpoint
}

func (s *PythonServer) Run() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(wd, s.Path)

	py, err := exec.LookPath("python3")
	if err != nil {
		// fallback to python and hope it still points to a python3 version.
		// the Windows python3 installer doesn't seem to create a `python3.exe`
		py, err = exec.LookPath("python")
		if err != nil {
			return errors.Errorf("unable to find Python executable")
		}
	}

	s.cmd = exec.Command(py, path)
	s.Endpoint = "opc.tcp://127.0.0.1:4840"
	s.Opts = []opcua.Option{opcua.SecurityMode(ua.MessageSecurityModeNone)}
	if err := s.cmd.Start(); err != nil {
		return err
	}

	// wait until endpoint is available
	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		c, err := net.Dial("tcp", "127.0.0.1:4840")
		if err != nil {
			time.Sleep(100 * time.Millisecond)
			continue
		}
		c.Close()
		return nil
	}
	return errors.Errorf("timeout")
}

func (s *PythonServer) Close() error {
	if s.cmd == nil {
		return errors.Errorf("not running")
	}
	go func() { s.cmd.Process.Kill() }()
	return s.cmd.Wait()
}
