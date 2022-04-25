package domain

import (
	"braverats/brp"
	"errors"
	"sort"
)

type Card struct {
	ID    brp.CardID
	Power uint8
}

type CardHand []*Card

func (hand *CardHand) ExtractCard(id brp.CardID) (*Card, error) {
	for i, card := range *hand {
		if card.ID == id {
			*hand = append((*hand)[:i], (*hand)[i+1:]...)
			return card, nil
		}
	}
	return nil, errors.New("card not found")
}

func (hand *CardHand) Sort() {
	sort.Slice(*hand, func(i, j int) bool {
		return (*hand)[i].Power < (*hand)[j].Power
	})
}

var standardDeck = map[brp.CardID]Card{
	brp.CardPrince: {
		ID:    brp.CardPrince,
		Power: 7,
	},
	brp.CardGeneral: {
		ID:    brp.CardGeneral,
		Power: 6,
	},
	brp.CardWizard: {
		ID:    brp.CardWizard,
		Power: 5,
	},
	brp.CardAmbassador: {
		ID:    brp.CardAmbassador,
		Power: 4,
	},
	brp.CardAssassin: {
		ID:    brp.CardAssassin,
		Power: 3,
	},
	brp.CardSpy: {
		ID:    brp.CardSpy,
		Power: 2,
	},
	brp.CardPrincess: {
		ID:    brp.CardPrincess,
		Power: 1,
	},
	brp.CardMusician: {
		ID:    brp.CardMusician,
		Power: 0,
	},
}

func StandardHand() CardHand {
	var hand CardHand

	for _, card := range standardDeck {
		hand = append(hand, &Card{
			ID:    card.ID,
			Power: card.Power,
		})
	}

	hand.Sort()
	return hand
}
