package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("---------- Printing All Cards...")
	cards.print()

	fmt.Println("---------- Printing Remaining Cards...")
	hand, remainingCards := deal(cards, 5)
	remainingCards.print()

	fmt.Println("---------- Printing Cards In Hand...")
	hand.print()

	fmt.Println("---------- Converting deck to string...")
	cardsString := cards.toString()
	fmt.Println(cardsString)

	fmt.Println("---------- Writing deck to the file system...")
	filename := "deck.txt"
	cards.saveToFile(filename)

	fmt.Println("---------- Loading deck from the file system...")
	newCards := loadDeckFromFile(filename)
	newCards.print()

	// newCards2 := loadDeckFromFile("cards.txt")
	// fmt.Println(newCards2)
	fmt.Println("---------- Shuffling cards...")
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five Of Diamonds"
}
