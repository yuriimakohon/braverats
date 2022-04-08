package domain

import (
	"errors"
	"sort"
)

type CardID uint8

const (
	Prince CardID = iota
	General
	Wizard
	Ambassador
	Assassin
	Spy
	Princess
	Musician
)

type Card struct {
	CardID
	Name        string
	Description string
	Power       uint8
}

type CardHand []Card

func (hand *CardHand) ExtractCard(idx uint8) (Card, error) {
	if idx < 0 || idx >= uint8(len(*hand)) {
		return Card{}, errors.New("index of card is out of hand range")
	}

	card := (*hand)[idx]
	*hand = append((*hand)[:idx], (*hand)[idx+1:]...)

	return card, nil
}

func (hand *CardHand) Sort() {
	sort.Slice(*hand, func(i, j int) bool {
		return (*hand)[i].Power < (*hand)[j].Power
	})
}

var standardDeck = map[CardID]Card{
	Prince: {
		CardID: Prince,
		Name:   "Prince",
		Power:  7,
	},
	General: {
		CardID: General,
		Name:   "General",
		Power:  6,
	},
	Wizard: {
		CardID: Wizard,
		Name:   "Wizard",
		Power:  5,
	},
	Ambassador: {
		CardID: Ambassador,
		Name:   "Ambassador",
		Power:  4,
	},
	Assassin: {
		CardID: Assassin,
		Name:   "Assassin",
		Power:  3,
	},
	Spy: {
		CardID: Spy,
		Name:   "Spy",
		Power:  2,
	},
	Princess: {
		CardID: Princess,
		Name:   "Princess",
		Power:  1,
	},
	Musician: {
		CardID: Musician,
		Name:   "Musician",
		Power:  0,
	},
}

func StandardHand() CardHand {
	var hand CardHand

	for _, card := range standardDeck {
		hand = append(hand, card)
	}

	hand.Sort()
	return hand
}
