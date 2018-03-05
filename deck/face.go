//go:generate stringer -type=Face

package deck

// Face represents the face value of the card.
type Face uint8

const (
	_ Face = iota
	// Ace represents an ace.
	Ace
	// Two represents a two.
	Two
	// Three represents a three.
	Three
	// Four represents a four.
	Four
	// Five represents a five.
	Five
	// Six represents a six.
	Six
	// Seven represents a seven.
	Seven
	// Eight represents an eight.
	Eight
	// Nine represents a nine.
	Nine
	// Ten represents a ten.
	Ten
	// Jack represents a jack.
	Jack
	// Queen represents a queen.
	Queen
	// King represents a king.
	King
)
