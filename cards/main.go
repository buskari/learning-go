package main

func main() {
	cards := newDeckFromFile("my_first_deck")
	// cards.saveToFile("my_first_deck")
	cards.shuffle()
	cards.print()
}
