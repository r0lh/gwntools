package gwntools

import (
	"fmt"
	"io"
	"net"
	"os/exec"
)

// Remote connects to a remote TCP service
func Remote(addr string) (*Process, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to process: %v", err)
	}

	return NewProcess(conn, conn), nil
}

// Local connects to a local command
func Local(cmd *exec.Cmd) (*Process, error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("Error opening stdout from process: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, fmt.Errorf("Error opening stderr from process: %v", err)
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("Error opening stdin from process: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("Start cmd error: %v", err)
	}

	return NewProcess(io.MultiReader(stdout, stderr), stdin), nil
}
