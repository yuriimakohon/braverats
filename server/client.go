package server

import (
	"braverats/brp"
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
)

type client struct {
	id     uuid.UUID
	conn   net.Conn
	server *Server
	name   string

	lobby      *lobby
	lobbyOwner bool
	ready      bool

	match *match
}

func newClient(conn net.Conn, server *Server) *client {
	return &client{
		id:     uuid.New(),
		conn:   conn,
		name:   "player",
		server: server,
	}
}

func (c *client) handleRequest(packet brp.Packet) error {
	if packet.Type != brp.TypeReq {
		return errors.New("client sent a non request packet")
	}

	switch packet.Tag {
	case brp.ReqSetName:
		c.handleSetName(packet.Payload)
	case brp.ReqCreateLobby:
		c.handleCreateLobby(packet.Payload)
	case brp.ReqJoinLobby:
		c.handleJoinLobby(packet.Payload)
	case brp.ReqLeaveLobby:
		c.handleLeaveLobby()
	case brp.ReqSetReadiness:
		c.handleSetReadiness(packet.Payload)
	case brp.ReqStartMatch:
		c.handleStartMatch()
	case brp.ReqPutCard:
		c.handlePutCard(packet.Payload)
	}

	return nil
}

func (c *client) handleWriteErr(tag brp.TAG, err error) {
	if err != nil {
		log.Printf("client %s didn't receive an %s tag: %v", c.id, tag, err)
	}
}
