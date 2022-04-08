package domain

type Round struct {
	Effects RoundEffect
	Result  RoundResult
}

type RoundResult uint8

const (
	FPWR RoundResult = iota // First player wins round
	SPWR                    // Second player wins round
	FPWG                    // First player one wins game
	SPWG                    // Second player wins game
	Hold                    // Hold this round
	Draw                    // the game ended in a Draw
)

type RoundEffect uint8

const (
	FPSpyPlayed RoundEffect = 1 << iota
	SPSpyPlayed
	FPGeneralPlayed
	SPGeneralPlayed
	FPAmbassadorPlayed
	SPAmbassadorPlayed
)

// Swap first and second RoundEffect bit
func (re *RoundEffect) Swap(a, b RoundEffect) {
	mask := a | b
	if !(*re&mask == mask || *re&mask^mask == mask) {
		*re ^= mask
	}
}

func (re RoundEffect) Has(e RoundEffect) bool {
	return re&e != 0
}
