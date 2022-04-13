package server

import (
	"errors"

	"github.com/google/uuid"
)

type lobby struct {
	name         string  // Unique lobby name
	firstPlayer  *client // First player in the lobby is owner
	secondPlayer *client // Second player in the lobby
}

func (l *lobby) removePlayer(id uuid.UUID) {
	if l.firstPlayer.id == id {
		l.firstPlayer.lobby = nil
		l.firstPlayer.ready = false
		l.firstPlayer = nil
	} else if l.secondPlayer.id == id {
		l.secondPlayer.lobby = nil
		l.secondPlayer.ready = false
		l.secondPlayer = nil
	}
}

func (l *lobby) startMatch() error {
	if l.firstPlayer == nil || l.secondPlayer == nil {
		return errors.New("need two players to start match")
	}

	if !l.firstPlayer.ready || !l.secondPlayer.ready {
		return errors.New("all players must be ready")
	}

	// TODO: create and start match
	return nil
}
