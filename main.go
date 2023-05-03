package main

import (
	"fmt"
	"time"

	"github.com/TafveezA/P2P-GoDApplication/deck"
	"github.com/TafveezA/P2P-GoDApplication/p2p"
)

func main() {
	cfg := p2p.ServerConfig{
		Version:    "P2PGame v0.1-alpha",
		ListenAddr: ":3000",
	}
	server := p2p.NewServer(cfg)
	go server.Start()
	time.Sleep(1 * time.Second)

	remotecfg := p2p.ServerConfig{
		Version:    "P2PGame v0.1-alpha",
		ListenAddr: ":4000",
	}

	remoteServer := p2p.NewServer(remotecfg)
	go remoteServer.Start()
	if err := remoteServer.Connect(":3000"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(deck.New())

}
