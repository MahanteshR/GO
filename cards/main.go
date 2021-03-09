package main

func main() {
	cards := newDeck()
	hands, remaining := deal(cards, 5)
	hands.print()
	remaining.print()
}
