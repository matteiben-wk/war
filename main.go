package main

import (
	"fmt"
	"strconv"
)

const (
	suites    = 4
	cardTypes = 13
)

type card struct {
	suite  int
	number int
}

func main() {
	cards := generateDeck()
	deck1, deck2 := splitDeck(cards)

	playWar(deck1, deck2)
}

func playWar(deck1, deck2 []*card) {
	var rounds int
	var pot []*card

	// Keep playing until either deck is empty.
	for len(deck1) != 0 && len(deck2) != 0 {
		var card1, card2 *card
		card1, deck1 = popCard(deck1)
		card2, deck2 = popCard(deck2)
		var player1Wins, player2Wins bool

		if card1.number > card2.number {
			pot = append(pot, card1, card2)
			player1Wins = true
		} else if card1.number < card2.number {
			pot = append(pot, card1, card2)
			player2Wins = true
		} else {
			// WAR

			var enoughCards1, enoughCards2 bool
			var downCards1, downCards2 []*card

			downCards1, deck1, enoughCards1 = popWarCards(deck1)
			downCards2, deck2, enoughCards2 = popWarCards(deck2)

			if !enoughCards1 {
				fmt.Println(">>> WAR")

				// Add all remaining cards in play to deck2 to trigger win condition
				deck2 = append(deck2, card1, card2)
				deck2 = append(deck2, downCards1...)
				deck2 = append(deck2, downCards2...)
				deck2 = append(deck2, pot...)

				rounds++
				continue
			}
			if !enoughCards2 {
				fmt.Println(">>> WAR")

				// Add all remaining cards in play to deck1 to trigger win condition
				deck1 = append(deck1, card1, card2)
				deck1 = append(deck1, downCards1...)
				deck1 = append(deck1, downCards2...)
				deck1 = append(deck1, pot...)

				rounds++
				continue
			}

			// Put all cards (2 played, 6 down) into the pot, then draw again
			pot = append(pot, card1, card2)
			pot = append(pot, downCards1...)
			pot = append(pot, downCards2...)

			rounds++
			continue
		}

		if player1Wins {
			deck1 = prependCards(deck1, pot...)
		} else if player2Wins {
			deck2 = prependCards(deck2, pot...)
		}

		pot = nil

		rounds++

		// fmt.Printf(">>> Total Length: %d - Deck1: %d, Deck2: %d\n", len(deck1)+len(deck2), len(deck1), len(deck2))

		if rounds == 1000000 {
			panic(`Something went horrifically awry ... you've played over ` + strconv.Itoa(rounds) + ` rounds and lost all sanity`)
		}
	}

	// Cleanup before final card tally

	if len(deck1) == 0 {
		fmt.Println(">>> Player 2 wins!!! - Player 1 out of cards.")
	}
	if len(deck2) == 0 {
		fmt.Println(">>> Player 1 wins!!! - Player 2 out of cards.")
	}

	fmt.Printf(">>> Total Rounds: %d\n", rounds)
	// fmt.Printf(">>> Total Length: %d - Deck1: %d, Deck2: %d\n", len(deck1)+len(deck2), len(deck1), len(deck2))
}
