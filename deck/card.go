package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Card represents the combination of a suit and a face value.
type Card struct {
	Suit
	Face
}

// String returns the string representation of a card.
func (c Card) String() string {
	if c.Suit == Joker {
		return fmt.Sprintf("%s", c.Suit)
	}
	return fmt.Sprintf("%s of %ss", c.Face, c.Suit)
}

// New returns a slice of all 52 cards in sorted order.
func New() []Card {
	var result []Card
	for _, suit := range suitsByRank {
		for _, face := range facesByRank {
			result = append(result, Card{suit, face})
		}
	}
	return result
}

// Sort by default sorting behavior.
func Sort(cards []Card) {
	sort.Slice(cards, func(i, j int) bool {
		return sortValue(cards[i]) < sortValue(cards[j])
	})
}

func sortValue(c Card) int {
	return int(c.Suit)*len(facesByRank) + int(c.Face)
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Shuffle randomly rearranges the cards in a new slice.
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// Jokers return n number of joker cards
func Jokers(n int) []Card {
	var jokers []Card
	for i := 0; i < n; i++ {
		jokers = append(jokers, Card{Joker, Face(i)})
	}
	return jokers
}

// Deck returns n number of decks each with 52 cards.
// Each deck is sequential in the slice.
func Deck(n int) []Card {
	var decks []Card
	for i := 0; i < n; i++ {
		decks = append(decks, New()...)
	}
	return decks
}

// Filter returns a subset of cards that satisfy the predicate condition.
func Filter(cards []Card, f func(card Card) bool) []Card {
	var matchingCards []Card
	for _, c := range cards {
		if !f(c) {
			matchingCards = append(matchingCards, c)
		}
	}
	return matchingCards
}
