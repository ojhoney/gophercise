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
	cards := NewDeck()

	if len(cards) != 52 {
		t.Error("Wrong number of cards in a new deck")
	}
}

func TestJoker(t *testing.T) {
	got := 0
	want := 4
	cards := NewDeck(AddJoker(want))
	for _, card := range cards {
		if card.Suit == Joker {
			got++
		}
	}
	if got != 4 {
		t.Errorf("Expected %v Jokers, got %v", want, got)
	}
}

func TestFilter(t *testing.T) {
	f := func(card Card) bool {
		return card.Suit == Spade
	}
	cards := NewDeck(Filter(f))
	for _, card := range cards {
		if card.Suit == Spade {
			t.Errorf("Expected no Spades, got Spades")
			return
		}
	}
}

func TestDeck(t *testing.T) {
	cards := NewDeck(AddDeck(3))
	got := len(cards)
	want := 13 * 4 * 3
	if got != want {
		t.Errorf("Expected %d cards, got %d cards", want, got)
	}
}
