package brp

type CardID uint8

const (
	CardUnknown CardID = iota
	CardMusician
	CardPrincess
	CardSpy
	CardAssassin
	CardAmbassador
	CardWizard
	CardGeneral
	CardPrince
)

func (c *CardID) Int() int {
	return int(*c)
}

func IsCardID(id CardID) bool {
	_, ok := cardIDs[id]
	return ok
}

// cardIDs is a set of all CardIDs.
var cardIDs = map[CardID]struct{}{
	CardUnknown:    {},
	CardMusician:   {},
	CardPrincess:   {},
	CardSpy:        {},
	CardAssassin:   {},
	CardAmbassador: {},
	CardWizard:     {},
	CardGeneral:    {},
	CardPrince:     {},
}
