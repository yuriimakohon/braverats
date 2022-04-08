package main

import (
	"braverats/domain"
	"fmt"
	"log"
)

func main() {
	match := domain.NewStandardMatch()

	var round domain.Round
	for round.Result == domain.FPWR || round.Result == domain.SPWR || round.Result == domain.Hold {
		printRoundStart(*match)
		var fpChoice, spChoice uint8
		_, err := fmt.Scanf("%d %d", &fpChoice, &spChoice)
		round, err = match.PlayRound(fpChoice, spChoice)
		if err != nil {
			log.Fatal(err)
		}
		printResult(round.Result)
		fmt.Printf("SCORE %d : %d\n\n", match.FPScore, match.SPScore)
	}
}

func printRoundStart(match domain.Match) {
	round := match.GetLastRound()
	if round != nil {
		if round.Effects.Has(domain.SPSpyPlayed) {
			fmt.Println("First Player shows card -_-")
		}
		if round.Effects.Has(domain.FPGeneralPlayed) {
			fmt.Println("First Player +2 Power")
		}
		if round.Effects.Has(domain.FPSpyPlayed) {
			fmt.Println("Second Player shows card -_-")
		}
		if round.Effects.Has(domain.SPGeneralPlayed) {
			fmt.Println("Second Player +2 Power")
		}
	}
	printHand(match.FPHand)
	printHand(match.SPHand)
}

func printHand(hand domain.CardHand) {
	for i, card := range hand {
		fmt.Printf("|(%d) %s %d|  ", i, card.Name, card.Power)
	}
	fmt.Println()
}

func printResult(result domain.RoundResult) {
	switch result {
	case domain.FPWR:
		fmt.Println("First player won round")
	case domain.SPWR:
		fmt.Println("Second player won round")
	case domain.Hold:
		fmt.Println("Round is held")
	case domain.FPWG:
		fmt.Println("First player won game !!!!!!!")
	case domain.SPWG:
		fmt.Println("Second player won game !!!!!!!")
	case domain.Draw:
		fmt.Println("Draw, congratulations for both players !!!!!!")
	}
}
