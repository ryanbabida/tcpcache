package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
)

type store[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(key K, val V)
}

type Server[K comparable, V any] struct {
	cfg       Config
	container store[K, V]
	logger    *slog.Logger
}

func NewServer[K comparable, V any](c store[K, V], cfg Config) *Server[K, V] {
	return &Server[K, V]{
		container: c,
		cfg:       cfg,
		logger:    slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

func (s *Server[K, V]) Run() {
	listener, err := net.Listen("tcp", ":"+*s.cfg.Port)
	if err != nil {
		s.logger.Error("unable to start tcp server")
		os.Exit(1)
	}

	s.logger.Info("TCP server listening on port :" + *s.cfg.Port)

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Error("failed to accept request", "error", err)
		}

		go s.HandleRequest(conn)
	}
}

func (s *Server[K, V]) HandleRequest(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadBytes('\n')
		if err != nil {
			s.logger.Error("unable to read message")
			conn.Write([]byte("unable to read message\n"))
			continue
		}

		var r request[K, V]
		err = json.Unmarshal(message, &r)
		if err != nil {
			s.logger.Error("unable to parse request")
			conn.Write([]byte("unable to parse request\n"))
			continue
		}

		err = r.isValid()
		if err != nil {
			s.logger.Error("invalid request", "error", err)
			conn.Write([]byte("invalid request\n"))
			continue
		}

		err = s.ExecuteAction(&r, conn)
		if err != nil {
			s.logger.Error("failed to execute action", "error", err)
			conn.Write([]byte("failed to execute action\n"))
			continue
		}

		conn.Write([]byte("Message received.\n"))
	}
}

func (s *Server[K, V]) ExecuteAction(r *request[K, V], conn net.Conn) error {
	switch r.Action {
	case Get:
		v, ok := s.container.Get(*r.Key)
		if ok {
			s := fmt.Sprintf("hit: %v\n", v)
			log.Println(s)
			conn.Write([]byte(s))
		} else {
			log.Printf("miss\n")
			conn.Write([]byte("miss\n"))
		}
	case Set:
		s.container.Set(*r.Key, *r.Value)
	default:
		return fmt.Errorf("unable to execute action")
	}

	return nil
}
