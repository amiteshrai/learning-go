package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck' which is a slice of strings
type deck []string

// Create a new deck instance
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, suite+" of "+value)
		}
	}
	return cards
}

// Deal the cards and return the remaining cards along with the hands.
func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	updatedDeck := d[handSize:]

	return updatedDeck, hand
}

// Print deck cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Convert a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

// Write a deck to a file
func (d deck) saveToFile(filename string) error {
	deckStr := d.toString()
	deckBytes := []byte(deckStr)
	return ioutil.WriteFile(filename, deckBytes, 0666)
}

// Load a deck from a file
func loadDeckFromFile(filename string) deck {
	deckBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error loading the deck from file")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	deckStr := string(deckBytes)
	deckSlice := strings.Split(deckStr, "\n")

	return deck(deckSlice)
}

// Shuffle the deck
func (d deck) shuffle() {
	seed := time.Now().UnixNano() // To ensure that seed is always different from previous execution
	source := rand.NewSource(seed)
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
