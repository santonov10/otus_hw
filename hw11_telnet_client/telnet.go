package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &tcpTelnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

type tcpTelnetClient struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func (t *tcpTelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	t.conn = conn
	return nil
}

func (t *tcpTelnetClient) Close() error {
	if t.conn != nil {
		err := t.conn.Close()
		if err != nil {
			return fmt.Errorf("close: %w", err)
		}
	}
	return nil
}

func (t *tcpTelnetClient) Send() error {
	_, err := io.Copy(t.conn, os.Stdin)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

func (t *tcpTelnetClient) Receive() error {
	_, err := io.Copy(os.Stdout, t.conn)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}
