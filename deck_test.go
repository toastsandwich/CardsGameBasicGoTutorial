package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // t -> test handler
	d := newDeck()
	size := 52
	fCard := "A Hearts"
	lCard := "K Clubs"
	if len(d) != size {
		t.Errorf("Expected deck length of %v, but got %v", size, len(d))
	}
	if d[0] != fCard {
		t.Errorf("Expected first card as %v but got : %v", fCard, d[0])
	}
	if d[size-1] != lCard {
		t.Errorf("Expeceted last card as %v but got : %v", lCard, d[size-1])
	}
}

func TestSaveToFileAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadDeck := newDeck().readFile("_decktesting")
	if len(loadDeck) != 52 {
		t.Errorf("Expected deck length of %v, but got %v", 52, len(d))
	}
}
