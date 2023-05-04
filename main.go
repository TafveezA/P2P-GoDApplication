package main

import (
	"log"
	"time"

	"github.com/TafveezA/P2P-GoDApplication/p2p"
)

func main() {
	cfg := p2p.ServerConfig{
		Version:     "P2PGame v0.1-alpha",
		ListenAddr:  ":3000",
		GameVariant: p2p.TexasHoldem,
	}
	server := p2p.NewServer(cfg)
	go server.Start()
	time.Sleep(2 * time.Second)

	remotecfg := p2p.ServerConfig{
		Version:     "P2PGame v0.1-alpha",
		ListenAddr:  ":4000",
		GameVariant: p2p.TexasHoldem,
	}

	remoteServer := p2p.NewServer(remotecfg)
	go remoteServer.Start()
	if err := remoteServer.Connect(":3000"); err != nil {
		log.Fatal(err)
	}

	select {}

}
