package main

import (
	"math/rand"
	"time"
)

func generateDeck() []*card {
	cards := make([]*card, 0, suites*cardTypes)

	for i := 1; i <= suites; i++ {
		for j := 1; j <= cardTypes; j++ {
			cards = append(cards, &card{
				suite:  i,
				number: j,
			})
		}
	}

	return shuffleCards(cards)
}

func shuffleCards(cards []*card) []*card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}

func splitDeck(cards []*card) ([]*card, []*card) {
	deck1 := make([]*card, 0, suites*cardTypes)
	deck2 := make([]*card, 0, suites*cardTypes)

	// Deal out half deck to each player
	for i, card := range cards {
		if i%2 == 0 {
			deck1 = append(deck1, card)
		} else {
			deck2 = append(deck2, card)
		}
	}

	return deck1, deck2
}

// Split one card off the end of the deck.
// Returns the popped card and the deck.
func popCard(cards []*card) (choosenCard *card, deck []*card) {
	choosenCard, deck = cards[len(cards)-1], cards[:len(cards)-1]
	return choosenCard, deck
}

// Split three cards off the end of the deck.
// Returns the popped cards and the deck.
func popWarCards(cards []*card) (downCards []*card, deck []*card, enoughCards bool) {
	if len(cards) < 3 {
		return cards, nil, false
	}
	downCards, deck = cards[len(cards)-3:], cards[:len(cards)-3]
	return downCards, deck, true
}

// Prepends the given cards to the deck and returns the deck.
func prependCards(cards []*card, newCards ...*card) (newDeck []*card) {
	if len(newCards) > 1 {
		newCards = shuffleCards(newCards)
	}

	newDeck = append(newCards, cards...)
	return newDeck
}
