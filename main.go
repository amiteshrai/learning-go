package main

import "fmt"

type greetingBot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func printGreeting(gb greetingBot) {
	fmt.Println(gb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

func main() {
	fmt.Println("Starting Application ...")

	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
