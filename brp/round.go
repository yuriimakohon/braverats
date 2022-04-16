package brp

type RoundResult uint8

const (
	WinRound RoundResult = iota
	LoseRound
	HoldRound
	WinGame
	LoseGame
	DrawGame
)

func (r *RoundResult) Int() int {
	return int(*r)
}
