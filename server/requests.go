package server

import (
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

	err := c.lobby.startMatch()
	if err != nil {
		c.respErr(err)
	}

	log.Printf("client %s started match in %s lobby\n", c.id, c.lobby.name)
	c.matchStarted()
	c.lobby.secondPlayer.matchStarted()
}
