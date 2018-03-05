package deck

import "testing"

func TestNew(t *testing.T) {
	cards := New()
	if got, expected := len(cards), len(suitsByRank)*len(facesByRank); got != expected {
		t.Errorf("Received the wrong number of cards. Expected %d but got %d",
			got, expected)
	}
	for i, suit := range suitsByRank {
		for j, face := range facesByRank {
			c := cards[len(facesByRank)*i+j]
			if c.Suit != suit || c.Face != face {
				t.Error("not sorted properly")
			}
		}
	}
}

func TestSort(t *testing.T) {
	cards := New()
	Sort(cards)
	for i, suit := range suitsByRank {
		for j, face := range facesByRank {
			c := cards[len(facesByRank)*i+j]
			if c.Suit != suit || c.Face != face {
				t.Error("not sorted properly")
			}
		}
	}
}

func TestShuffleThenSort(t *testing.T) {
	cards := New()
	Shuffle(cards)
	Sort(cards)
	for i, suit := range suitsByRank {
		for j, face := range facesByRank {
			c := cards[len(facesByRank)*i+j]
			if c.Suit != suit || c.Face != face {
				t.Error("not sorted properly")
			}
		}
	}
}

func TestJokers(t *testing.T) {
	jokers := Jokers(10)
	if numOfJokers := len(jokers); numOfJokers != 10 {
		t.Errorf("Expected %d jokers, but got %d", 10, numOfJokers)
	}
	for _, j := range jokers {
		if j.Suit != Joker {
			t.Error("This should be a joker")
		}
	}
}

func TestDecks(t *testing.T) {
	fiveDecks := Deck(5)
	for d := 0; d < 5; d++ {
		for i, suit := range suitsByRank {
			for j, face := range facesByRank {
				c := fiveDecks[len(facesByRank)*i+j]
				if c.Suit != suit || c.Face != face {
					t.Error("not sorted properly")
				}
			}
		}
	}
}

func TestCardString(t *testing.T) {
	joker := Card{Joker, Face(100)}
	if joker.String() != "Joker" {
		t.Error("Expected a Joker")
	}
	aceOfSpades := Card{Spade, Ace}
	if aceOfSpades.String() != "Ace of Spades" {
		t.Error("Expected a Ace of Spades")
	}
	excessive := Card{Suit(100), Face(100)}
	if got := excessive.String(); got != "Face(100) of Suit(100)s" {
		t.Errorf("Stringer didn't do what I expected. I got %s", got)
	}
}

func TestFilter(t *testing.T) {
	cards := New()
	noNumbersFilter := func(c Card) bool {
		for i := 1; i < 10; i++ {
			if c.Face == facesByRank[i] {
				return true
			}
		}
		return false
	}
	noNumbersCards := Filter(cards, noNumbersFilter)
	if num := len(noNumbersCards); num != 16 {
		t.Errorf("Expected 16 non-number cards, but got %d", num)
	}
}
