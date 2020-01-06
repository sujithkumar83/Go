package main

func main() {
	//var card string = "Ace of Spades"
	//card:="Ace of Spades"
	//var card := newCard()
	//cards := deck{"Ace of Spades", newCard()}
	//cards = append(cards, "Six of Spades")
	cards := newDeck()
	cards.print()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()

	//cards.savetoFile("myCards")
	//cardnew := newDeckFromFile("myCards")
	//cardnew.print()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
