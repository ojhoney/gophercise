//go:generate stringer -type=Suit, Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

type Card struct {
	Suit
	Rank
}

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func NewDeck(opts ...Opts) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{suit, rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

type Opts func(cards []Card) []Card

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(0))
	ret := make([]Card, len(cards))
	for i, j := range r.Perm(len(cards)) {
		ret[i] = cards[j]
	}
	return ret
}

func AddJoker(num_of_joker int) Opts {
	return func(cards []Card) []Card {
		for i := 0; i < num_of_joker; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) Opts {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

func AddDeck(num_of_deck int) Opts {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < num_of_deck; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
