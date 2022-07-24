package main

import (
	"amiteshrai/learning-go/src/deck"
	"fmt"
)

func main() {
	cards := deck.NewDeck()
	fmt.Println("---------- Printing All Cards...")
	cards.Print()

	fmt.Println("---------- Printing Remaining Cards...")
	hand, remainingCards := deck.Deal(cards, 5)
	remainingCards.Print()

	fmt.Println("---------- Printing Cards In Hand...")
	hand.Print()

	fmt.Println("---------- Converting deck to string...")
	cardsString := cards.ToString()
	fmt.Println(cardsString)

	fmt.Println("---------- Writing deck to the file system...")
	filename := "output/deck.txt"
	cards.SaveToFile(filename)

	fmt.Println("---------- Loading deck from the file system...")
	newCards := deck.LoadDeckFromFile(filename)
	newCards.Print()

	// newCards2 := loadDeckFromFile("cards.txt")
	// fmt.Println(newCards2)
	fmt.Println("---------- Shuffling cards...")
	cards.Shuffle()
	cards.Print()
}
