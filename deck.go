package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string
//printing the deck
func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}
//creating the new deck
func newDeck() deck {
	cards := deck{}
	cardsuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardvalues := []string{"Ace", "Two", "Three", "Four"}
	for _, suits := range cardsuits {
		for _, values := range cardvalues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}
//handsize meanse to divide the slice
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]

}
// join the slice of strings in the file
func (d deck) toString() string {
	a := []string(d)
	result := strings.Join(a, ",")
	return result

}
//create a text file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
//creating a new deck from file
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	a := string(bs)
	s := strings.Split(a, ",")
	return deck(s)
}
//it is imcomplete
func (d deck) shuffle() {
	for index := range d {
		newposition := rand.Intn(len(d) - 1)
		d[index], d[newposition] = d[newposition], d[index]

	}

}
