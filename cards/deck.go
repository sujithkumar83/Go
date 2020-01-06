package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Crate a new type of 'deck' which is a slice of strings

type deck []string

// Reciever function. This enables the method print() to every deck type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// _ can be used if index is not going to be used in a for loop

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Examples of slice function. A slice[:x] references a range [0,x)
func deal(d deck, dealsize int) (deck, deck) {
	return d[:dealsize], d[dealsize:]
}

// Conversion example. []string(d) simply refers to converting an array of string to a single string
// stirngs pkg has many useful methods like Join (concatnates with a separator)

func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

// Conversion example. []byte(d) simply refers to converting an string to an array of bytes
func (d deck) savetoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// in case of no error Go returns the value nil against error.
// os pkg has non-platform (windows, linux) utilities like save, delete or read from file available
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1: Log the error and return a call to newDeck
		// Optiom 2: Log the error and exit the profgram
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Note the nifty swap betwen two variables
// Randomizing the seed value everytime: An exampl below
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
