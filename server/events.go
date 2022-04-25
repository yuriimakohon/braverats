package server

import (
	"braverats/brp"
	"strconv"
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

func (c *client) cardPut(faceUp bool, card brp.CardID) {
	c.logResponse(brp.EventCardPut, strconv.Itoa(card.Int()))
	_, err := c.conn.Write(brp.NewEventCardPut(faceUp, card))
	c.handleWriteErr(brp.EventCardPut, err)
}

func (c *client) roundEnded(result brp.RoundResult, card brp.CardID) {
	c.logResponse(brp.EventRoundEnded, strconv.Itoa(result.Int()))
	_, err := c.conn.Write(brp.NewEventRoundEnded(result, card))
	c.handleWriteErr(brp.EventRoundEnded, err)
}
