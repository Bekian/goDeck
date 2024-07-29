package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

// idiomatic deck struct
// has methods for dealing, shuffling, and printing the deck
type Deck struct {
	cards     []int
	cardCount int
}

//Think: is there an instance where a deck should be created that has card values and is not length 52?

// create a zeroed deck of size amt
func initZeroedDeck(amt int) Deck {
	// create a zeroed slice who is of length `amt` but cannot exceed the size of 52
	deck := make([]int, amt, 52)
	return Deck{cards: deck, cardCount: amt}
}

// create a full deck
func initDeck() Deck {
	deck := make([]int, 52, 52)
	for i, _ := range deck {
		deck[i] = i + 1
	}
	return Deck{cards: deck, cardCount: 52}
}

func (d *Deck) print() {
	fmt.Println(d.cards)
}

// removes an amt items from the top of the deck into a new deck
func (d *Deck) deal(amt int) Deck {
	// make a new deck of size amt
	newDeck := make([]int, amt, 52)
	for i, _ := range newDeck {
		// assign the last items from the original deck into the new deck
		newDeck[i] = d.cards[d.cardCount-i-1]
	}
	// remove the items from the original deck
	d.cards = d.cards[:d.cardCount-amt]
	// return the new deck
	return Deck{cards: newDeck, cardCount: amt}
}

func (d *Deck) compareDeck(other *Deck) bool {
	itemMap := make(map[int]bool)

	for _, card := range d.cards {
		itemMap[card] = true
	}

	for _, card := range other.cards {
		if itemMap[card] {
			return true
		}
	}

	return false
}

// merges 2 decks into one
func (d *Deck) combine(otherDeck *Deck) (err error) {
	// ensure total length of deck is not greater than 52
	if d.cardCount+otherDeck.cardCount > 52 {
		return errors.New("Total card count too high")
	}
	// compare the decks to ensure they dont have one or more of the same cards in both decks
	if d.compareDeck(otherDeck) {
		return errors.New("Decks have more than one of the same card")
	}
	// move items to d
	d.cards = append(d.cards, otherDeck.cards...)
	d.cardCount += otherDeck.cardCount
	// remove items from otherDeck
	otherDeck.cards = []int{}
	otherDeck.cardCount = 0
	// return no error
	return nil
}

// shuffles a deck in place
func (d *Deck) shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

func main() {
	deck := initDeck()
	deck.print()

	deck.shuffle()
	deck.print()

	hand := deck.deal(5)
	hand.print()
	deck.print()
	err := deck.combine(&hand)
	if err != nil {
		fmt.Println(err)
	}
	deck.print()
	hand.print()
}
