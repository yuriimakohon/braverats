package server

import (
	"braverats/brp"
	"braverats/domain"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func (c *client) setName(args []byte) {
	if len(args) == 0 {
		c.respErr(errors.New("player`s name can not be empty"))
		return
	}

	c.name = string(args)
	c.respOk(fmt.Sprintf("name %q set", c.name))
}

func (c *client) createLobby(args []byte) {
	if c.lobby != nil {
		c.respErr(errors.New("you are already in a lobby"))
		return
	}

	if len(args) == 0 {
		c.respErr(errors.New("lobby name can not be empty"))
		return
	}

	if _, ok := c.server.lobbies[string(args)]; ok {
		c.respInfo(fmt.Sprintf("lobby with name %q already exists", string(args)))
		return
	}

	c.lobby = &lobby{
		name:        string(args),
		firstPlayer: c,
	}
	c.lobbyOwner = true
	c.server.lobbies[string(args)] = c.lobby

	c.respOk(fmt.Sprintf("lobby %q created", c.lobby.name))
}

func (c *client) joinLobby(args []byte) {
	if len(args) == 0 {
		c.respErr(errors.New("lobby name can not be empty"))
		return
	}

	if c.lobby != nil {
		c.respErr(errors.New("you are already in a lobby"))
		return
	}

	lobby, ok := c.server.lobbies[string(args)]
	if !ok {
		c.respInfo(fmt.Sprintf("lobby with name %q doesn't exists", string(args)))
		return
	}

	c.lobby = lobby
	c.lobby.secondPlayer = c
	log.Printf("client %s joined lobby %q\n", c.id, lobby.name)

	c.lobby.firstPlayer.joinedLobby(c.name)
	c.respLobby(c.lobby.firstPlayer.ready, c.lobby.firstPlayer.name)
}

func (c *client) leaveLobby() {
	if c.lobby == nil {
		c.respErr(errors.New("you are not in a lobby"))
		return
	}

	lobby := c.lobby

	if c.lobbyOwner {
		delete(c.server.lobbies, lobby.name)
		c.lobbyOwner = false

		if lobby.secondPlayer != nil {
			lobby.secondPlayer.lobbyClosed()
			lobby.removePlayer(lobby.secondPlayer.id)
		}
	} else {
		lobby.firstPlayer.leftLobby(c.name)
	}

	lobby.removePlayer(c.id)
	c.respOk(fmt.Sprintf("left %q lobby", lobby.name))
}

func (c *client) setReadiness(args []byte) {
	if c.lobby == nil {
		c.respErr(errors.New("you are not in a lobby"))
		return
	}

	ready, err := strconv.ParseBool(string(args))
	if err != nil {
		c.respErr(errors.New("invalid readiness value"))
		return
	}

	var currentPlayer = c
	var anotherPlayer *client

	if c.lobbyOwner {
		anotherPlayer = c.lobby.secondPlayer
	} else {
		anotherPlayer = c.lobby.firstPlayer
	}

	if currentPlayer.ready != ready {
		currentPlayer.ready = ready
		if anotherPlayer != nil {
			anotherPlayer.playerReadiness(ready)
		}
	}

	currentPlayer.respOk(fmt.Sprintf("readiness set to %t", ready))
}

func (c *client) startMatch() {
	if c.lobby == nil {
		c.respErr(errors.New("you are not in lobby"))
		return
	}

	if !c.lobbyOwner {
		c.respErr(errors.New("you are not lobby owner"))
		return
	}

	c.match = &match{
		m:            domain.NewStandardMatch(),
		firstPlayer:  c,
		secondPlayer: c.lobby.secondPlayer,
	}
	c.lobby.secondPlayer.match = c.match

	if !c.match.firstPlayer.ready && !c.match.secondPlayer.ready {
		c.respErr(errors.New("both players must be ready"))
		return
	}

	log.Printf("client %s started match in %s lobby\n", c.id, c.lobby.name)
	c.respOk("match started")
	c.matchStarted()
	c.lobby.secondPlayer.matchStarted()
}

func (c *client) putCard(args []byte) {
	if c.match == nil {
		c.respErr(errors.New("you are not in a match"))
		return
	}

	type player struct {
		c      *client
		faceUp bool
		card   **domain.Card
	}

	fp := player{
		c:    c.match.firstPlayer,
		card: &c.match.fpCard,
	}
	sp := player{
		c:    c.match.secondPlayer,
		card: &c.match.spCard,
	}
	var me, opponent *player
	if c.lobbyOwner {
		me = &fp
		opponent = &sp
	} else {
		me = &sp
		opponent = &fp
	}

	lastRound := c.match.m.GetLastRound()
	if lastRound != nil {
		fp.faceUp = lastRound.Effects.Has(domain.SPSpyPlayed)
		sp.faceUp = lastRound.Effects.Has(domain.FPSpyPlayed)

		if opponent.faceUp && *opponent.card == nil {
			c.respErr(errors.New("you can't put a card first after spy played"))
			return
		}
	}

	card, err := c.match.putCard(c.lobbyOwner, args)
	if err != nil {
		c.respErr(err)
		return
	}
	*me.card = card

	c.respOk("")
	opponent.c.cardPut(me.faceUp, (*me.card).ID)

	if *opponent.card != nil {
		round, err := c.match.m.PlayRound(**fp.card, **sp.card)
		if err != nil {
			c.respErr(err)
			return
		}
		switch round.Result {
		case domain.FPWR:
			fp.c.roundEnded(brp.WonRound, (*sp.card).ID)
			sp.c.roundEnded(brp.LoosedRound, (*fp.card).ID)
		case domain.SPWR:
			fp.c.roundEnded(brp.LoosedRound, (*sp.card).ID)
			sp.c.roundEnded(brp.WonRound, (*fp.card).ID)
		case domain.Hold:
			fp.c.roundEnded(brp.HeldRound, (*sp.card).ID)
			sp.c.roundEnded(brp.HeldRound, (*fp.card).ID)
		case domain.FPWG:
			fp.c.roundEnded(brp.WonGame, (*sp.card).ID)
			sp.c.roundEnded(brp.LoosedGame, (*fp.card).ID)
			fp.c.ready = false
			sp.c.ready = false
			c.match = nil
		case domain.SPWG:
			fp.c.roundEnded(brp.LoosedGame, (*sp.card).ID)
			sp.c.roundEnded(brp.WonGame, (*fp.card).ID)
			fp.c.ready = false
			sp.c.ready = false
		case domain.Draw:
			fp.c.roundEnded(brp.DrawGame, (*sp.card).ID)
			sp.c.roundEnded(brp.DrawGame, (*fp.card).ID)
		}
		*fp.card = nil
		*sp.card = nil
	}
}
