package netsrv

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

type HandlerFunc func(r *bufio.Reader, w io.ReadWriter) error

type Server struct {
	h       HandlerFunc
	timeout time.Duration
}

func New(timeout time.Duration) *Server {
	return &Server{
		timeout: timeout,
	}
}

func (s *Server) Listen(address string) error {
	listener, err := net.Listen("tcp4", address)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		fmt.Println("Connection established")
		conn.SetDeadline(time.Now().Add(s.timeout * time.Second))
		go s.handler(conn)
	}
}

// Добавить функцию обработчик
func (s *Server) RegisterHandler(f HandlerFunc) {
	s.h = f
}

// обработчик подключения
func (s *Server) handler(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("Connection closed")

	r := bufio.NewReader(conn)
	for {
		err := s.h(r, conn)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		conn.SetDeadline(time.Now().Add(s.timeout * time.Second))
	}
}
