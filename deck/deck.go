package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

// card struct makes up a deck
type Card struct {
	rank 	int 	// 1-13 card's value within the suit
	suit 	string 	// hearts, clubs, diamonds, spades
	id 		int 	// 1-52 position in an unshuffled deck
	flipped bool 	// whether or not the specific card is flipped
}

// idiomatic deck struct
// has methods for dealing, shuffling, and printing the deck
type Deck struct {
	cards     []Card
	cardCount int
}

// flips a single card
func (c *Card) flipCard() {
	c.flipped = !c.flipped
}
// flips a whole deck
func (d *Deck) flipDeck() {
	for _, card := range d.cards {
		card.flipCard()
	}
}

//Think: is there an instance where a deck should be created that has card values and is not length 52?

// TODO: this function needs to be refactored
// create a zeroed deck of size amt 
func initZeroedDeck(amt int) Deck {
	// create a zeroed slice who is of length `amt` but cannot exceed the size of 52
	deck := make([]int, amt, 52)
	return Deck{cards: deck, cardCount: amt}
}

// TODO: this function needs to be refactored
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

// TODO: this function needs to be refactored
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
	d.cardCount -= amt
	// return the new deck
	return Deck{cards: newDeck, cardCount: amt}
}

// TODO: this function needs to be refactored
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

// TODO: this function needs to be refactored
// merges 2 decks into one
func (d *Deck) combine(otherDeck *Deck) (err error) {
	// ensure total length of deck is not greater than 52
	if d.cardCount+otherDeck.cardCount > 52 {
		new_amt := d.cardCount+otherDeck.cardCount
		errorString := fmt.Sprintf("Total card count too high, amount: %d", new_amt)
		return errors.New(errorString)
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
	// init deck
	deck := initDeck()
	deck.print()
	// shuffle deck
	deck.shuffle()
	deck.print()
	// deal cards from the deck to a hand
	hand := deck.deal(5)
	hand.print()
	deck.print()
	// combine the hand back into the main deck
	err := deck.combine(&hand)
	if err != nil {
		fmt.Println(err)
	}
	deck.print()
	hand.print()
}
