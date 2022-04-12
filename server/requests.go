package server

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func (c *client) setName(args []byte) {
	if len(args) == 0 {
		c.err(errors.New("player`s name can not be empty"))
		return
	}

	c.name = string(args)
	c.ok(fmt.Sprintf("name %q set", c.name))
}

func (c *client) createLobby(args []byte) {
	if c.lobby != nil {
		c.err(errors.New("you are already in a lobby"))
		return
	}

	if len(args) == 0 {
		c.err(errors.New("lobby name can not be empty"))
		return
	}

	if _, ok := c.server.lobbies[string(args)]; ok {
		c.info(fmt.Sprintf("lobby with name %q already exists", string(args)))
		return
	}

	c.lobby = &lobby{
		name:        string(args),
		firstPlayer: c,
	}
	c.lobbyOwner = true
	c.server.lobbies[string(args)] = c.lobby

	c.ok(fmt.Sprintf("lobby %q created", c.lobby.name))
}

func (c *client) joinLobby(args []byte) {
	if len(args) == 0 {
		c.err(errors.New("lobby name can not be empty"))
		return
	}

	if c.lobby != nil {
		c.err(errors.New("you are already in a lobby"))
		return
	}

	lobby, ok := c.server.lobbies[string(args)]
	if !ok {
		c.info(fmt.Sprintf("lobby with name %q doesn't exists", string(args)))
		return
	}

	c.lobby = lobby
	c.lobby.secondPlayer = c
	log.Printf("client %s joined lobby %q\n", c.id, lobby.name)

	c.lobby.firstPlayer.joinedLobby(c.name)
	c.ok(fmt.Sprintf("joined to %q lobby", lobby.name))
}

func (c *client) leaveLobby() {
	if c.lobby == nil {
		c.err(errors.New("you are not in a lobby"))
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
	c.ok(fmt.Sprintf("left %q lobby", lobby.name))
}

func (c *client) setReadiness(args []byte) {
	if c.lobby == nil {
		c.err(errors.New("you are not in a lobby"))
		return
	}

	ready, err := strconv.ParseBool(string(args))
	if err != nil {
		c.err(errors.New("invalid readiness value"))
		return
	}

	var currentPlayer, anotherPlayer *client

	if c.lobbyOwner {
		currentPlayer = c
		anotherPlayer = c.lobby.secondPlayer
	} else {
		currentPlayer = c.lobby.secondPlayer
		anotherPlayer = c
	}

	if currentPlayer.ready != ready {
		currentPlayer.ready = ready
		if anotherPlayer != nil {
			anotherPlayer.playerReadiness(ready)
		}
	}

	currentPlayer.ok(fmt.Sprintf("readiness set to %t", ready))
}

func (c *client) startMatch() {
	if c.lobby == nil {
		c.err(errors.New("you are not in lobby"))
		return
	}

	if !c.lobbyOwner {
		c.err(errors.New("you are not lobby owner"))
		return
	}

	err := c.lobby.startMatch()
	if err != nil {
		c.err(err)
	}

	log.Printf("client %s started match in %s lobby\n", c.id, c.lobby.name)
	c.matchStarted()
	c.lobby.secondPlayer.matchStarted()
}
