package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: King, Suit: Spade})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// King of Spades
	// Joker
}

func TestNew(t *testing.T) {
	cards := NewDecks()

	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck")
	}

}
