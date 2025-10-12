package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("First card isn't Ace of Spades, its %s", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Last card isn't Four of Clubs, its %s", d[0])
	}
}
