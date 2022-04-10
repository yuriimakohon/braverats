package server

import (
	"braverats/brp"
	"log"
)

func (c *client) err(err error) {
	log.Printf("client %s caused error: %v\n", c.id, err)
	_, err = c.conn.Write(brp.RespErr(err))
	c.handleWriteErr(brp.TagErr, err)
}

func (c *client) ok() {
	_, err := c.conn.Write(brp.RespOk())
	c.handleWriteErr(brp.TagOk, err)
}

func (c *client) joinedLobby(name string) {
	_, err := c.conn.Write(brp.RespJoinedLobby(name))
	c.handleWriteErr(brp.TagJoinedLobby, err)
}

func (c *client) leftLobby(name string) {
	_, err := c.conn.Write(brp.RespLeftLobby(name))
	c.handleWriteErr(brp.LeftLobby, err)
}

func (c *client) playerReadiness(ready bool) {
	_, err := c.conn.Write(brp.RespPlayerReadiness(ready))
	c.handleWriteErr(brp.TagPlayerReadiness, err)
}

func (c *client) matchStarted() {
	_, err := c.conn.Write(brp.RespMatchStarted())
	c.handleWriteErr(brp.TagMatchStarted, err)
}
