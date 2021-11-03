package gwntools

import (
	"io"
	"log"
	"net"
	"os/exec"
)

// Remote connects to a remote TCP service
func Remote(addr string) *Process {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("Error connecting to process: %v", err)
	}

	return NewProcess(conn, conn)
}

// Local connects to a local command
func Local(cmd *exec.Cmd) *Process {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error opening stdout from process: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("Error opening stderr from process: %v", err)
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal("Error opening stdin from process: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal("Start cmd error: %v", err)
	}

	return NewProcess(io.MultiReader(stdout, stderr), stdin)
}
