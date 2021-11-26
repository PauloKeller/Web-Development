package main

func main() {
	cards := newDeckFromFile("cards")
	cards.shuffle()
	cards.print()
}
