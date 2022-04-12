package server

import (
	"errors"
	"log"
	"strconv"
)

func (c *client) setName(args []byte) {
	if len(args) == 0 {
		c.err(errors.New("player`s name can not be empty"))
		return
	}

	c.name = string(args)
	log.Printf("client %s set name to %s\n", c.id, c.name)
	c.ok()
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
		c.err(errors.New("lobby with such name already exists"))
		return
	}

	c.lobby = &lobby{
		name:        string(args),
		firstPlayer: c,
	}
	c.lobbyOwner = true
	c.server.lobbies[string(args)] = c.lobby

	log.Printf("client %s created lobby %s\n", c.id, c.lobby.name)
	c.ok()
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
		c.err(errors.New("lobby with such name doesn't exists"))
		return
	}

	c.lobby = lobby
	c.lobby.secondPlayer = c
	log.Printf("client %s joined lobby %s\n", c.id, lobby.name)

	c.lobby.firstPlayer.joinedLobby(c.name)
	c.ok()
}

func (c *client) leaveLobby() {
	if c.lobby == nil {
		c.err(errors.New("you are not in a lobby"))
		return
	}

	if c.lobbyOwner {
		c.lobbyOwner = false
		delete(c.server.lobbies, c.lobby.name)

		if c.lobby.secondPlayer != nil {
			c.lobby.secondPlayer.lobby = nil
			c.lobby.secondPlayer.leftLobby(c.name)
		}
	} else {
		c.lobby.firstPlayer.leftLobby(c.name)
	}

	c.lobby = nil
	log.Printf("client %s left lobby %s\n", c.id, c.lobby.name)
	c.ok()
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
		log.Printf("client %s set readiness to %t\n", c.id, ready)
	}

	currentPlayer.ok()
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
