package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"
	"net"
)

type Server struct {
	ln       net.Listener
	coffsets map[string]int
	buffer   []Message
}

func NewServer() *Server {
	return &Server{
		coffsets: make(map[string]int),
		buffer:   make([]Message, 0),
	}
}

func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", ":9091")
	if err != nil {
		return err
	}
	s.ln = ln
	for {
		conn, err := ln.Accept()
		if err != nil {
			if err == io.EOF {
				return err
			}
			slog.Error("server accept error", "err", err)
		}
		go s.HandleConn(conn)
	}
}

func (s *Server) HandleConn(conn net.Conn) {
	slog.Info("new connection coming", "the remote address", conn.RemoteAddr())

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			slog.Error("connection read error", "err", err)
			return
		}
		rawMsg := buf[:n]
		r := bytes.NewReader(rawMsg)
		var header Header
		binary.Read(r, binary.BigEndian, &header)
		fmt.Println(header)
	}
}
