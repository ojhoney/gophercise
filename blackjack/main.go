package main

import (
	"fmt"
	"gophercises/deck"
	"strings"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := 0; i < len(h); i++ {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func main() {
	cards := deck.NewDeck(deck.AddDeck(3), deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ {
		card, cards = cards[0], cards[1:]
		fmt.Println(card)
	}

	h := Hand(cards[0:3])
	fmt.Println(h)
}

func (h Hand) MinScore() int {
	score := 0
	for _, card := range h {
		score += min(10, int(card.Rank))
	}
	return score
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, card := range h {
		if card.Rank == deck.Ace {
			minScore += 10
		}
	}
	return minScore
}

func shouldHit(h Hand) {
	if h.Score() <= 16 || (h.Score() == 17 && h.MinScore() == 7) {
		return
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
