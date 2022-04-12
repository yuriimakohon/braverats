package server

import (
	"braverats/brp"
	"io"
	"log"
	"net"

	"github.com/google/uuid"
)

type Server struct {
	ln      net.Listener
	clients map[uuid.UUID]*client
	lobbies map[string]*lobby
}

// NewServer creates a new server
func NewServer() *Server {
	return &Server{
		clients: make(map[uuid.UUID]*client, 0),
		lobbies: make(map[string]*lobby, 0),
	}
}

// Start the server and listen for connections
func (s *Server) Start(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	s.ln = ln

	log.Println("Server started on port " + port)

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) addClient(conn net.Conn) *client {
	c := newClient(conn, s)
	s.clients[c.id] = c
	log.Println("new client added with id ", c.id.String())
	return c
}

func (s *Server) removeClient(id uuid.UUID) {
	c, ok := s.clients[id]
	if !ok {
		log.Printf("can't remove client %s doesn't exists", id.String())
		return
	}

	lobby := c.lobby

	if lobby != nil {
		if c.lobbyOwner {
			delete(c.server.lobbies, lobby.name)

			if lobby.secondPlayer != nil {
				lobby.secondPlayer.lobbyClosed()
				lobby.removePlayer(lobby.secondPlayer.id)
			}
		} else {
			lobby.firstPlayer.leftLobby(c.name)
		}

		lobby.removePlayer(c.id)
	}

	c.conn.Close()
	delete(s.clients, id)
	log.Println("client ", c.conn.RemoteAddr().String(), " ", c.id.String(), " disconnected")
}

func (s *Server) handleConnection(conn net.Conn) {
	client := s.addClient(conn)

	for {
		packet, err := brp.ReadPacket(client.conn)
		if err == io.EOF {
			s.removeClient(client.id)
			return
		}
		if err != nil {
			log.Printf("Error read packet from %s: %s\n", conn.RemoteAddr().String(), err)
			client.err(err)
			continue
		}
		s.handleReq(packet, client)
	}
}

func (s *Server) handleReq(packet brp.Packet, c *client) {
	switch packet.Tag {
	case brp.ReqSetName:
		c.setName(packet.Payload)
	case brp.ReqCreateLobby:
		c.createLobby(packet.Payload)
	case brp.ReqJoinLobby:
		c.joinLobby(packet.Payload)
	case brp.ReqLeaveLobby:
		c.leaveLobby()
	case brp.ReqSetReadiness:
		c.setReadiness(packet.Payload)
	case brp.ReqStartMatch:
		c.startMatch()
	}
}
