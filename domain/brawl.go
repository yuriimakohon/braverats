package domain

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

var BrawlFunctions = map[[2]CardID]BrawlFunction{
	{Prince, Prince}:         anyAnyStronger,
	{Prince, General}:        princeGeneral,
	{Prince, Wizard}:         anyAnyStronger,
	{Prince, Ambassador}:     fpwr,
	{Prince, Assassin}:       fpwr,
	{Prince, Spy}:            princeSpy,
	{Prince, Princess}:       spwg,
	{Prince, Musician}:       hold,
	{General, General}:       generals,
	{General, Wizard}:        anyAnyStronger,
	{General, Ambassador}:    generalAmbassador,
	{General, Assassin}:      generalAssassin,
	{General, Spy}:           generalSpy,
	{General, Princess}:      generalPrincess,
	{General, Musician}:      generalMusician,
	{Wizard, Wizard}:         anyAnyStronger,
	{Wizard, Ambassador}:     wizardAmbassador,
	{Wizard, Assassin}:       fpwr,
	{Wizard, Spy}:            fpwr,
	{Wizard, Princess}:       fpwr,
	{Wizard, Musician}:       fpwr,
	{Ambassador, Ambassador}: ambassadors,
	{Ambassador, Assassin}:   ambassadorAssassin,
	{Ambassador, Spy}:        ambassadorSpy,
	{Ambassador, Princess}:   ambassadorPrincess,
	{Ambassador, Musician}:   ambassadorMusician,
	{Assassin, Assassin}:     anyAnyWeaker,
	{Assassin, Spy}:          assassinSpy,
	{Assassin, Princess}:     anyAnyWeaker,
	{Assassin, Musician}:     hold,
	{Spy, Spy}:               anyAnyStronger,
	{Spy, Princess}:          spyPrincess,
	{Spy, Musician}:          spyMusician,
	{Princess, Princess}:     anyAnyStronger,
	{Princess, Musician}:     hold,
	{Musician, Musician}:     hold,
}
