package server

import (
	"net"
	"sync"
)

type Peer struct {
	conn net.com
}
type ServerConfig struct {
	listenAddr string
}
type Server struct {
	ServerConfig
	mu      sync.RWMutex
	peers   map[net.Addr]*Peer
	addpeer chan *Peer
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
	}
}

func (s *Server) Start() {

}
func (s *Server) loop() {

	for {
		select {}
	}
}
