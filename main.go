package main

func main() {
	// slice := deck{"0 Ace of cards", "1 Ace of cards", "2 Ace of cards"}
	cards := newDeck()
	/*cards1, cards2 := deal(cards, 5)
	fmt.Println("CARDS1-----------CARDS1")

	cards1.print()
	fmt.Println("CARDS2-----------CARDS2")
	cards2.print()
	// slice:=[]int{12,3}
	// fmt.Println(slice)
	// greeting := "Hi There!"
	// fmt.Println([]byte(greeting))*/
	//fmt.Println(cards.toString())
	// fmt.Println(cards.saveToFile("my_cards"))
	// fmt.Println(newDeckFromFile("my_cards"))
	// slice := newDeckFromFile("my_cards")
	// slice.print()
	cards.shuffle()
	cards.print()

}
