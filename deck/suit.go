//go:generate stringer -type=Suit

package deck

// Suit represents the suit family the card belongs to.
type Suit uint8

const (
	// Spade represents the spade suit.
	Spade Suit = iota
	// Diamond represents the diamond suit.
	Diamond
	// Club represents the club suit.
	Club
	// Heart represents the heart suit.
	Heart
	// Joker represents jokers.
	Joker
)

var suitsByRank = [...]Suit{Spade, Diamond, Club, Heart}
