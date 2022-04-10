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
	c := s.clients[id]
	if c.lobby != nil && c.lobbyOwner {
		delete(s.lobbies, c.lobby.name)
	}
	c.conn.Close()
	delete(s.clients, id)
	log.Println("client " + c.conn.RemoteAddr().String() + " " + c.id.String() + "  disconnected")
}

func (s *Server) handleConnection(conn net.Conn) {
	client := s.addClient(conn)

	for {
		tag, args, err := brp.ReadPacket(client.conn)
		if err == io.EOF {
			s.removeClient(client.id)
			return
		}
		if err != nil {
			log.Printf("Error read packet from %s: %s\n", conn.RemoteAddr().String(), err)
			client.err(err)
			continue
		}
		s.handleReq(tag, args, client)
	}
}

func (s *Server) handleReq(tag brp.TAG, args []byte, c *client) {
	switch tag {
	case brp.TagSetName:
		c.setName(args)
	case brp.TagCreateLobby:
		c.createLobby(args)
	case brp.TagJoinLobby:
		c.joinLobby(args)
	case brp.TagLeaveLobby:
		c.leaveLobby()
	case brp.TagSetReadiness:
		c.setReadiness(args)
	case brp.TagStartMatch:
		c.startMatch()
	}
}
