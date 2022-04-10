package server

import (
	"braverats/brp"
	"log"
	"net"

	"github.com/google/uuid"
)

type client struct {
	id     uuid.UUID
	conn   net.Conn
	server *Server

	name       string
	lobby      *lobby
	lobbyOwner bool
	ready      bool
}

func newClient(conn net.Conn, server *Server) *client {
	return &client{
		id:     uuid.New(),
		conn:   conn,
		name:   "player",
		server: server,
	}
}

func (c *client) handleWriteErr(tag brp.TAG, err error) {
	if err != nil {
		log.Printf("client %s didn't receive an %s tag: %v", c.id, tag, err)
	}
}
