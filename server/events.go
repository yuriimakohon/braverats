package server

import (
	"braverats/brp"
	"strconv"
)

func (c *client) emitEvent(tag brp.TAG, event []byte, msg string) {
	c.logOutgoingTag(tag, msg)
	_, err := c.conn.Write(event)
	c.handleWriteErr(tag, err)
}

func (c *client) eventJoinedLobby(name string) {
	c.emitEvent(brp.EventJoinedLobby, brp.NewEventJoinedLobby(name), name)
}

func (c *client) eventLeftLobby(name string) {
	c.emitEvent(brp.EventLeftLobby, brp.NewEventLeftLobby(name), name)
}

func (c *client) eventLobbyClosed() {
	c.emitEvent(brp.EventLobbyClosed, brp.NewEventLobbyClosed(), "")
}

func (c *client) eventPlayerReadiness(ready bool) {
	c.emitEvent(brp.EventPlayerReadiness, brp.NewEventPlayerReadiness(ready), strconv.FormatBool(ready))
}

func (c *client) eventMatchStarted() {
	c.emitEvent(brp.EventMatchStarted, brp.NewEventMatchStarted(), "")
}

func (c *client) eventCardPut(faceUp bool, card brp.CardID) {
	c.emitEvent(brp.EventCardPut, brp.NewEventCardPut(faceUp, card), strconv.Itoa(card.Int()))
}

func (c *client) eventRoundEnded(result brp.RoundResult, card brp.CardID) {
	c.emitEvent(brp.EventRoundEnded, brp.NewEventRoundEnded(result, card), strconv.Itoa(result.Int()))
}
