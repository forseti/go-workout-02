package main

func main() {
	// cards.print()

	// hand, remainingCards := deal(cards, 4)

	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards := newDeckFromFile("my_123")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
