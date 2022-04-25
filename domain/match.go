package domain

import (
	"braverats/brp"
	"errors"
)

type Match struct {
	Rounds   []*Round
	FPHand   CardHand
	SPHand   CardHand
	FPScore  uint8
	SPScore  uint8
	WinScore uint8
}

func NewStandardMatch() *Match {
	match := &Match{
		Rounds:   make([]*Round, 0, 8),
		FPHand:   StandardHand(),
		SPHand:   StandardHand(),
		FPScore:  0,
		SPScore:  0,
		WinScore: 4,
	}
	return match
}

func (m *Match) PlayRound(fpCard, spCard Card) (Round, error) {
	if len(m.Rounds) != 0 {
		lastRound := m.GetLastRound()
		if lastRound.Effects.Has(FPGeneralPlayed) {
			fpCard.Power += 2
		}
		if lastRound.Effects.Has(SPGeneralPlayed) {
			spCard.Power += 2
		}
	}

	round, err := m.brawlCards(fpCard, spCard)
	if err != nil {
		return round, err
	}

	return m.finishRound(round), nil
}

func (m *Match) GetLastRound() *Round {
	if m.Rounds == nil || len(m.Rounds) == 0 {
		return nil
	}
	return m.Rounds[len(m.Rounds)-1]
}

func (m *Match) brawlCards(fpCard, spCard Card) (Round, error) {
	var round Round

	brawl, ok := BrawlFunctions[[2]brp.CardID{fpCard.ID, spCard.ID}]
	if ok {
		round = brawl(fpCard, spCard)
	} else if brawl, ok = BrawlFunctions[[2]brp.CardID{spCard.ID, fpCard.ID}]; ok {
		round = brawl(spCard, fpCard)
		inverseRound(&round)
	} else {
		return round, errors.New("card pair for brawl doesn't exists")
	}

	return round, nil
}

// finishRound counts all scores and review result of brawl
func (m *Match) finishRound(round Round) Round {
	switch round.Result {
	case FPWR:
		for i := len(m.Rounds) - 1; i >= 0 && m.Rounds[i].Result == Hold; i-- {
			m.Rounds[i].Result = FPWR
			if m.Rounds[i].Effects.Has(FPAmbassadorPlayed) {
				m.FPScore++
			}
			m.FPScore++
		}
		if round.Effects.Has(FPAmbassadorPlayed) {
			m.FPScore++
		}
		m.FPScore++
		if m.FPScore >= m.WinScore {
			round.Result = FPWG
		}
	case SPWR:
		for i := len(m.Rounds) - 1; i >= 0 && m.Rounds[i].Result == Hold; i-- {
			m.Rounds[i].Result = SPWR
			if m.Rounds[i].Effects.Has(SPAmbassadorPlayed) {
				m.SPScore++
			}
			m.SPScore++
		}
		if round.Effects.Has(SPAmbassadorPlayed) {
			m.SPScore++
		}
		m.SPScore++
		if m.SPScore >= m.WinScore {
			round.Result = SPWG
		}
	case Hold:
		if len(m.FPHand) == 0 && len(m.SPHand) == 0 {
			round.Result = Draw
		}
	}
	m.Rounds = append(m.Rounds, &round)
	return round
}

func inverseRound(r *Round) {
	switch r.Result {
	case FPWR:
		r.Result = SPWR
	case SPWR:
		r.Result = FPWR
	case FPWG:
		r.Result = SPWG
	case SPWG:
		r.Result = FPWG
	}

	r.Effects.Swap(FPSpyPlayed, SPSpyPlayed)
	r.Effects.Swap(FPGeneralPlayed, SPGeneralPlayed)
	r.Effects.Swap(FPAmbassadorPlayed, SPAmbassadorPlayed)
}
