package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards, leftCards := deal(cards, 6)
	cards.show()
	leftCards.saveToFile("myCards.txt")
}
