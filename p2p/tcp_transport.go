package p2p

import (
	"bytes"
	"net"

	"github.com/sirupsen/logrus"
)

type Peer struct {
	conn     net.Conn
	outbound bool
}

func (p *Peer) Send(b []byte) error {
	_, err := p.conn.Write(b)
	return err
}
func (p *Peer) ReadLoop(msgch chan *Message) {

	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {

			break
		}
		msgch <- &Message{
			From:    p.conn.RemoteAddr(),
			Payload: bytes.NewReader(buf[:n]),
		}

	}
	p.conn.Close()

}

type TCPTransport struct {
	ListenAddr string
	listener   net.Listener
	addPeer    chan *Peer
	delPeer    chan *Peer
}

func NewTCPTransport(addr string) *TCPTransport {
	return &TCPTransport{
		ListenAddr: addr,
	}
}
func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	t.listener = ln

	for {
		conn, err := ln.Accept()
		if err != nil {
			logrus.Error(err)
			continue
		}
		peer := &Peer{
			conn:     conn,
			outbound: false,
		}
		t.addPeer <- peer
		//peer.ReadLoop()
	}

}
