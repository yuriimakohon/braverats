package server

import (
	"braverats/protocol"
	"log"
)

func (c *client) err(err error) {
	log.Printf("client %s caused error: %v\n", c.id, err)
	_, err = c.conn.Write(protocol.RespErr(err))
	c.handleWriteErr(protocol.Err, err)
}

func (c *client) ok() {
	_, err := c.conn.Write(protocol.RespOk())
	c.handleWriteErr(protocol.Ok, err)
}

func (c *client) joinedLobby(name string) {
	_, err := c.conn.Write(protocol.RespJoinedLobby(name))
	c.handleWriteErr(protocol.JoinedLobby, err)
}

func (c *client) leftLobby(name string) {
	_, err := c.conn.Write(protocol.RespLeftLobby(name))
	c.handleWriteErr(protocol.LeftLobby, err)
}

func (c *client) playerReadiness(ready bool) {
	_, err := c.conn.Write(protocol.RespPlayerReadiness(ready))
	c.handleWriteErr(protocol.PlayerReadiness, err)
}

func (c *client) matchStarted() {
	_, err := c.conn.Write(protocol.RespMatchStarted())
	c.handleWriteErr(protocol.MatchStarted, err)
}
