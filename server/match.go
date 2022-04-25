package server

import (
	"braverats/brp"
	"braverats/domain"
	"errors"
	"strconv"
)

type match struct {
	m            *domain.Match
	firstPlayer  *client // First player in the lobby is owner
	secondPlayer *client // Second player in the lobby
	fpCard       *domain.Card
	spCard       *domain.Card
}

func (m *match) putCard(firstPlayer bool, payload []byte) (*domain.Card, error) {
	var card *domain.Card

	if firstPlayer {
		card = m.fpCard
	} else {
		card = m.spCard
	}

	if card != nil {
		return nil, errors.New("you already played a card")
	}

	idInt, err := strconv.Atoi(string(payload))
	if err != nil {
		return nil, err
	}

	if firstPlayer {
		card, err = m.m.FPHand.ExtractCard(brp.CardID(idInt))
	} else {
		card, err = m.m.SPHand.ExtractCard(brp.CardID(idInt))
	}
	if err != nil {
		return nil, err
	}

	return card, nil
}
