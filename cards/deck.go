package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) display() {
	fmt.Println(strings.Join(d, ", "))
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toByteSlice() []byte {
	return []byte(strings.Join(d, ", "))
}

func toStringSlice(byteSlice []byte) deck {
	return strings.Split(string(byteSlice), ", ")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, d.toByteSlice(), 0666)
}

func readFromFile(filename string) deck {
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return toStringSlice(fileContents)
}

func (d deck) shuffleDeck() deck {

	for i := range d {
		newPosition := rand.IntN(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
