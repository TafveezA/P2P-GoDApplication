package main
import (
	"fmt"
	"github.com/TafveezA/P2P-GoDApplication/p2p"

)

func main(){
	cfg := p2p.ServerConfig{
		listenAddr: ":3000"
	}
	server := p2p.NewServer(cfg)
	server.start()
}