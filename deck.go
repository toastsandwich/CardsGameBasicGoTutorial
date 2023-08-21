package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creating a type which is slice of string
type deck []string

// new deck
func newDeck() deck {
	cards := deck{}
	suites := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	for _, suite := range suites {
		for _, value := range values {
			var card string = value + " " + suite
			cards = append(cards, card)
		}
	}
	return cards
}

// this is a reciever function
func (d deck) show() {
	fmt.Print("Your hand :\n")
	for _, card := range d {
		fmt.Printf("%s ", card)
	}
}

// shuffle the deck
func (d deck) shuffle() {
	//creating a randomizer so that the Intn gets a diffrent int 64 every time.
	//The time package is used so that to get a different Int64 number using UnixNano.
	source := rand.NewSource(time.Now().UnixNano())
	Rand := rand.New(source)
	//simply created our own rand.Intn so that its impossible to get same shuffle anytime.
	for i := range d {
		newPosition := Rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// deal function a function returning multiple values
func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}

// ----------------------------------
// save to File
func (d deck) toString() string {
	return strings.Join(d, ";")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666) //0666 perm so that anyone can read or write the file.
}

// ----------------------------------
// to read file
func (d deck) readFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//splits converts to []string
	str := strings.Split(string(byteSlice), ";")
	return deck(str)
}
