package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Suits = [...]string{
	"Hearts", "Spades", "Clubs", "Diamonds",
}

var Ranks = [...]string{
	"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King",
}

type Card struct {
	Suit string
	Rank string
}

func (c Card) String() string {
	return fmt.Sprintf("%5s of %s", c.Rank, c.Suit)
}

type Deck []Card

func NewDeck() Deck {
	d := make([]Card, 0)

	for s := 0; s < len(Suits); s++ {
		for r := 0; r < len(Ranks); r++ {
			d = append(d, Card{Suit: Suits[s], Rank: Ranks[r]})
		}
	}

	return d
}

func (d Deck) Shuffle() Deck {
	size := len(d)

	var s Deck = make([]Card, size)
	rand.Seed(time.Now().Unix())
	m := make(map[int]bool)

	for i := 0; i < size; i++ {
		var index int
		ok := true

		for ok {
			index = rand.Intn(size)
			_, ok = m[index]
		}

		m[index] = true
		s[index] = d[i]
	}

	return s
}

func (d Deck) Show() {
	for i := 0; i < len(d); i++ {
		fmt.Printf("%02d: %s\n", i+1, d[i].String())
	}
}

func main() {
	deck := NewDeck()
	deck.Show()

	fmt.Printf("\nShuffling deck...\n\n")

	deck = deck.Shuffle()
	deck.Show()
}
