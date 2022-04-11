package server

import "braverats/brp"

func (c *client) joinedLobby(name string) {
	_, err := c.conn.Write(brp.NewEventJoinedLobby(name))
	c.handleWriteErr(brp.EventJoinedLobby, err)
}

func (c *client) leftLobby(name string) {
	_, err := c.conn.Write(brp.NewEventLeftLobby(name))
	c.handleWriteErr(brp.EventLeftLobby, err)
}

func (c *client) playerReadiness(ready bool) {
	_, err := c.conn.Write(brp.NewEventPlayerReadiness(ready))
	c.handleWriteErr(brp.EventPlayerReadiness, err)
}

func (c *client) matchStarted() {
	_, err := c.conn.Write(brp.NewEventMatchStarted())
	c.handleWriteErr(brp.EventMatchStarted, err)
}
