package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// Very Custom Logic for generating English greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// Very Custom Logic for generating Spanish greeting
	return "Hola!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
