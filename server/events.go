package server

import (
	"braverats/brp"
)

func (c *client) joinedLobby(name string) {
	c.logResponse(brp.EventJoinedLobby, name)
	_, err := c.conn.Write(brp.NewEventJoinedLobby(name))
	c.handleWriteErr(brp.EventJoinedLobby, err)
}

func (c *client) leftLobby(name string) {
	c.logResponse(brp.EventLeftLobby, name)
	_, err := c.conn.Write(brp.NewEventLeftLobby(name))
	c.handleWriteErr(brp.EventLeftLobby, err)
}

func (c *client) lobbyClosed() {
	c.logResponse(brp.EventLobbyClosed, "")
	_, err := c.conn.Write(brp.NewEventLobbyClosed())
	c.handleWriteErr(brp.EventLobbyClosed, err)
}

func (c *client) playerReadiness(ready bool) {
	c.logResponse(brp.EventPlayerReadiness, "")
	_, err := c.conn.Write(brp.NewEventPlayerReadiness(ready))
	c.handleWriteErr(brp.EventPlayerReadiness, err)
}

func (c *client) matchStarted() {
	c.logResponse(brp.EventMatchStarted, "")
	_, err := c.conn.Write(brp.NewEventMatchStarted())
	c.handleWriteErr(brp.EventMatchStarted, err)
}
