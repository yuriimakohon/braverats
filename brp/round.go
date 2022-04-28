package brp

type RoundResult uint8

const (
	WonRound RoundResult = iota
	LoosedRound
	HeldRound
	WonGame
	LoosedGame
	DrawGame
)

func (r *RoundResult) Int() int {
	return int(*r)
}
