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

	isFirstPlayer := c.id == c.match.firstPlayer.id

	card, err := c.match.putCard(isFirstPlayer, args)
	if err != nil {
		c.respErr(err)
		return
	}

	var opponentCard, playerCard **domain.Card
	var opponent *client
	if isFirstPlayer {
		playerCard = &c.match.fpCard
		opponentCard = &c.match.spCard
		opponent = c.match.secondPlayer
	} else {
		playerCard = &c.match.spCard
		opponentCard = &c.match.fpCard
		opponent = c.match.firstPlayer
	}

	var faceUp, opponentFaceUp bool

	lastRound := c.match.m.GetLastRound()
	if lastRound != nil {
		faceUp = isFirstPlayer && lastRound.Effects.Has(domain.SPSpyPlayed) || !isFirstPlayer && lastRound.Effects.Has(domain.FPSpyPlayed)
		opponentFaceUp = isFirstPlayer && lastRound.Effects.Has(domain.FPSpyPlayed) || !isFirstPlayer && lastRound.Effects.Has(domain.SPSpyPlayed)

		if opponentFaceUp && opponentCard == nil {
			c.respErr(errors.New("you can't put a card first after spy played"))
			return
		}
	}

	*playerCard = card

	c.respOk("")
	opponent.cardPut(faceUp, (*playerCard).ID)

	if *opponentCard != nil {
		round, err := c.match.m.PlayRound(*card, **opponentCard)
		if err != nil {
			c.respErr(err)
			return
		}
		switch round.Result {
		case domain.FPWR:
			c.roundEnded(brp.WinRound, (*opponentCard).ID)
			opponent.roundEnded(brp.LoseRound, card.ID)
		case domain.SPWR:
			c.roundEnded(brp.LoseRound, (*opponentCard).ID)
			opponent.roundEnded(brp.WinRound, card.ID)
		case domain.Hold:
			c.roundEnded(brp.HoldRound, (*opponentCard).ID)
			opponent.roundEnded(brp.HoldRound, card.ID)
		case domain.FPWG:
			c.roundEnded(brp.WinGame, (*opponentCard).ID)
			opponent.roundEnded(brp.LoseGame, card.ID)
		case domain.SPWG:
			c.roundEnded(brp.LoseGame, (*opponentCard).ID)
			opponent.roundEnded(brp.WinGame, card.ID)
		case domain.Draw:
			c.roundEnded(brp.DrawGame, (*opponentCard).ID)
			opponent.roundEnded(brp.DrawGame, card.ID)
		}
		*playerCard = nil
		*opponentCard = nil
	}
}
