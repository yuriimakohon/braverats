package domain

import "braverats/brp"

type BrawlFunction func(firstCard, secondCard Card) Round

func anyAnyStronger(c1, c2 Card) Round {
	round := Round{}

	if c1.Power < c2.Power {
		round.Result = SPWR
	} else if c1.Power > c2.Power {
		round.Result = FPWR
	} else {
		round.Result = Hold
	}
	return round
}

func anyAnyWeaker(c1, c2 Card) Round {
	round := Round{}

	if c1.Power < c2.Power {
		round.Result = FPWR
	} else if c1.Power > c2.Power {
		round.Result = SPWR
	} else {
		round.Result = Hold
	}
	return round
}

func fpwr(_, _ Card) Round {
	return Round{Result: FPWR}
}

func spwg(_, _ Card) Round {
	return Round{Result: SPWG}
}

// hold holds the round
func hold(_, _ Card) Round {
	return Round{Result: Hold}
}

// Prince
func princeGeneral(p, g Card) Round {
	round := fpwr(p, g)
	round.Effects |= SPGeneralPlayed
	return round
}

func princeSpy(c1, c2 Card) Round {
	round := anyAnyStronger(c1, c2)
	round.Effects |= SPSpyPlayed
	return round
}

// General
func generals(g1, g2 Card) Round {
	round := anyAnyStronger(g1, g2)
	round.Effects |= FPGeneralPlayed | SPGeneralPlayed
	return round
}

func generalAmbassador(g, a Card) Round {
	round := anyAnyStronger(g, a)
	round.Effects |= FPGeneralPlayed | SPAmbassadorPlayed
	return round
}

func generalAssassin(g, a Card) Round {
	round := anyAnyWeaker(g, a)
	round.Effects |= FPGeneralPlayed
	return round
}

func generalSpy(g, s Card) Round {
	round := anyAnyStronger(g, s)
	round.Effects |= FPGeneralPlayed | SPSpyPlayed
	return round
}

func generalPrincess(g, p Card) Round {
	round := anyAnyStronger(g, p)
	round.Effects |= FPGeneralPlayed
	return round
}

// generalMusician holds round and assigns cards effects
func generalMusician(g, m Card) Round {
	return Round{
		Result:  Hold,
		Effects: FPGeneralPlayed,
	}
}

// Wizard
func wizardAmbassador(w, a Card) Round {
	round := anyAnyStronger(w, a)
	round.Effects |= SPAmbassadorPlayed
	return round
}

// Ambassador
func ambassadors(a1, a2 Card) Round {
	round := anyAnyStronger(a1, a2)
	round.Effects |= FPAmbassadorPlayed | SPAmbassadorPlayed
	return round
}

func ambassadorAssassin(a, a2 Card) Round {
	round := anyAnyWeaker(a, a2)
	round.Effects |= FPAmbassadorPlayed
	return round
}

// ambassadorSpy determines stronger card and assigns theirs effects
func ambassadorSpy(a, s Card) Round {
	round := anyAnyStronger(a, s)
	round.Effects |= FPAmbassadorPlayed | SPSpyPlayed
	return round
}

// ambassadorPrincess determines stronger card and assigns theirs effects
func ambassadorPrincess(a, p Card) Round {
	round := anyAnyStronger(a, p)
	round.Effects |= FPAmbassadorPlayed
	return round
}

// ambassadorMusician holds round and assigns cards effects
func ambassadorMusician(a, m Card) Round {
	return Round{
		Result:  Hold,
		Effects: FPAmbassadorPlayed,
	}
}

// Assassin
// assassinSpy determines weaker card and assigns theirs effects
func assassinSpy(a, s Card) Round {
	round := anyAnyWeaker(a, s)
	round.Effects |= SPSpyPlayed
	return round
}

// Spy
// spyPrincess determines stronger card and assigns theirs effects
func spyPrincess(s, p Card) Round {
	round := anyAnyStronger(s, p)
	round.Effects |= FPSpyPlayed
	return round
}

// spyMusician holds round and assigns cards effects
func spyMusician(s, m Card) Round {
	return Round{
		Result:  Hold,
		Effects: FPSpyPlayed,
	}
}

var BrawlFunctions = map[[2]brp.CardID]BrawlFunction{
	{brp.CardPrince, brp.CardPrince}:         anyAnyStronger,
	{brp.CardPrince, brp.CardGeneral}:        princeGeneral,
	{brp.CardPrince, brp.CardWizard}:         anyAnyStronger,
	{brp.CardPrince, brp.CardAmbassador}:     fpwr,
	{brp.CardPrince, brp.CardAssassin}:       fpwr,
	{brp.CardPrince, brp.CardSpy}:            princeSpy,
	{brp.CardPrince, brp.CardPrincess}:       spwg,
	{brp.CardPrince, brp.CardMusician}:       hold,
	{brp.CardGeneral, brp.CardGeneral}:       generals,
	{brp.CardGeneral, brp.CardWizard}:        anyAnyStronger,
	{brp.CardGeneral, brp.CardAmbassador}:    generalAmbassador,
	{brp.CardGeneral, brp.CardAssassin}:      generalAssassin,
	{brp.CardGeneral, brp.CardSpy}:           generalSpy,
	{brp.CardGeneral, brp.CardPrincess}:      generalPrincess,
	{brp.CardGeneral, brp.CardMusician}:      generalMusician,
	{brp.CardWizard, brp.CardWizard}:         anyAnyStronger,
	{brp.CardWizard, brp.CardAmbassador}:     wizardAmbassador,
	{brp.CardWizard, brp.CardAssassin}:       fpwr,
	{brp.CardWizard, brp.CardSpy}:            fpwr,
	{brp.CardWizard, brp.CardPrincess}:       fpwr,
	{brp.CardWizard, brp.CardMusician}:       fpwr,
	{brp.CardAmbassador, brp.CardAmbassador}: ambassadors,
	{brp.CardAmbassador, brp.CardAssassin}:   ambassadorAssassin,
	{brp.CardAmbassador, brp.CardSpy}:        ambassadorSpy,
	{brp.CardAmbassador, brp.CardPrincess}:   ambassadorPrincess,
	{brp.CardAmbassador, brp.CardMusician}:   ambassadorMusician,
	{brp.CardAssassin, brp.CardAssassin}:     anyAnyWeaker,
	{brp.CardAssassin, brp.CardSpy}:          assassinSpy,
	{brp.CardAssassin, brp.CardPrincess}:     anyAnyWeaker,
	{brp.CardAssassin, brp.CardMusician}:     hold,
	{brp.CardSpy, brp.CardSpy}:               anyAnyStronger,
	{brp.CardSpy, brp.CardPrincess}:          spyPrincess,
	{brp.CardSpy, brp.CardMusician}:          spyMusician,
	{brp.CardPrincess, brp.CardPrincess}:     anyAnyStronger,
	{brp.CardPrincess, brp.CardMusician}:     hold,
	{brp.CardMusician, brp.CardMusician}:     hold,
}
