package gwntools

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Process is a local or a remote process to connect to
type Process struct {
	io.Reader
	io.Writer
	BufReader *bufio.Reader
}

// CreateProcess create a new process with an io.Reader and io.Writer
func NewProcess(r io.Reader, w io.Writer) *Process {
	return &Process{
		Reader:    r,
		Writer:    w,
		BufReader: bufio.NewReader(r),
	}
}

// Interactive interact with the process
func (p Process) Interactive() error {
	ch := make(chan error, 2)

	go func() {
		_, err := io.Copy(p.Writer, os.Stdin)
		ch <- err
	}()

	go func() {
		_, err := io.Copy(os.Stdout, p.BufReader)
		ch <- err
	}()

	for i := 0; i < 2; i++ {
		err := <-ch
		if err != nil {
			return fmt.Errorf("Error interacting with process: v%", err)
		}
	}

	return nil
}

// Write writes data to a process' stdin
func (p *Process) Write(data []byte) error {
	_, err := p.Writer.Write(data)
	if err != nil {
		return fmt.Errorf("Error writing data to process: %v", err)
	}

	return nil
}

// Readline read data from process until first newline "\n"
func (p *Process) ReadLine() ([]byte, error) {
	line, err := p.BufReader.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("Error reading line from process: %v", err)
	}

	return line, nil
}

// ReadByte reads and returns a single byte from process
func (p *Process) ReadByte() (byte, error) {
	b, err := p.BufReader.ReadByte()
	if err != nil {
		return 0, fmt.Errorf("Error reading byte from process: %v", err)
	}

	return b, nil
}
