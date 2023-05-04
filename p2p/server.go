package p2p

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"net"

	"github.com/sirupsen/logrus"
)

type GameVariant uint8

func (gv GameVariant) String() string {
	switch gv {
	case TexasHoldem:
		return "TEXAS HOLDEM"
	case Other:
		return "Other"
	default:
		return "unknown"

	}
}

const (
	TexasHoldem GameVariant = iota
	Other
)

type ServerConfig struct {
	Version     string
	ListenAddr  string
	GameVariant GameVariant
}

type Server struct {
	ServerConfig
	handler   Handler
	transport *TCPTransport
	listener  net.Listener
	peers     map[net.Addr]*Peer
	addPeer   chan *Peer
	delPeer   chan *Peer
	msgCh     chan *Message
	GameState *GameState
}

func NewServer(cfg ServerConfig) *Server {
	s := &Server{

		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
		delPeer:      make(chan *Peer),
		msgCh:        make(chan *Message),
		GameState:    NewGameState(),
	}
	tr := NewTCPTransport(s.ListenAddr)
	s.transport = tr
	tr.addPeer = s.addPeer
	tr.delPeer = s.delPeer
	return s
}

func (s *Server) Start() {
	go s.loop()

	fmt.Printf("Game Server is running on port %s\n ", s.ListenAddr)
	logrus.WithFields(logrus.Fields{
		"port":    s.ListenAddr,
		"variant": s.GameVariant,
	}).Info("Started New game Server")
	s.transport.ListenAndAccept()

}
func (s *Server) sendHandshake(p *Peer) error {
	hs := &Handshake{
		GameVariant: s.GameVariant,
		Version:     s.Version,
	}
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(hs); err != nil {
		return err
	}
	// if err := hs.Encode(buf); err != nil {
	// 	return err
	// }
	return p.Send(buf.Bytes())
}

// improvement is needed to the game network, handshake protocol
func (s *Server) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	peer := &Peer{
		conn: conn,
	}
	s.addPeer <- peer
	return peer.Send([]byte(s.Version))
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}
		peer := &Peer{
			conn: conn,
		}
		s.addPeer <- peer
		peer.Send([]byte(s.Version))

	}
}

func (s *Server) loop() {

	for {
		select {
		case peer := <-s.delPeer:
			logrus.WithFields(logrus.Fields{
				"addr": peer.conn.RemoteAddr(),
			}).Info("Player disconnected")
			delete(s.peers, peer.conn.RemoteAddr())
			// if new person connects  to the server
			fmt.Printf("player disconnected %s\n", peer.conn.RemoteAddr())
		case peer := <-s.addPeer:
			go s.sendHandshake(peer)
			if err := <-s.addPeer; err != nil {
				logrus.Info("handshake successful:handshake with incoming player failed")
				continue
			}
			// to check max player logic
			go peer.ReadLoop(s.msgCh)
			logrus.WithFields(logrus.Fields{
				"addr": peer.conn.RemoteAddr(),
			}).Info("New Player connected")
			s.peers[peer.conn.RemoteAddr()] = peer
			fmt.Printf("new player connected %s\n", peer.conn.RemoteAddr())
		case msg := <-s.msgCh:
			if err := s.handleMessage(msg); err != nil {
				panic(err)
			}
		}
	}
}
func (hs *Handshake) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, []byte(hs.Version)); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, &hs.GameVariant)
}
func (hs *Handshake) Decode(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, []byte(hs.Version)); err != nil {
		return err
	}

	return binary.Read(r, binary.LittleEndian, &hs.GameVariant)
}

type Handshake struct {
	Version     string
	GameVariant GameVariant
}

func (s *Server) handshake(p *Peer) error {
	hs := &Handshake{}
	if err := gob.NewDecoder(p.conn).Decode(hs); err != nil {

		return err
	}
	// if err := hs.Decode(p.conn); err !=nil{
	// 	return err
	// }
	if s.GameVariant != hs.GameVariant {
		return fmt.Errorf("Invalid Game Variant %s", hs.GameVariant)
	}
	if s.GameVariant != hs.GameVariant {
		return fmt.Errorf("Invalid version %s", hs.Version)
	}

	logrus.WithFields(logrus.Fields{
		"peer":    p.conn.RemoteAddr(),
		"version": hs.Version,
		"variant": hs.GameVariant}).Info("Recieved Handshake")
	return nil

}

func (s *Server) handleMessage(msg *Message) error {
	fmt.Printf("%v\n", msg)
	return nil
}
