package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Cards struct {
  Hand []Card
  Deck []Card
}

type Card struct {
	Suite string
	Value string
}

var cards Cards

func main() {
	cards := makeDeck()
	shuffledDeck := shuffleDeck(cards)
	fmt.Sprintln(shuffledDeck)

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Ask Me to Deal, Shuffle or Exit\n ")
	for input.Scan() {
		if input.Text() == "Deal" || input.Text() == "deal" {
      fmt.Printf("\n ### You asked me to %s ### \n", input.Text())
			dealOneCard(shuffledDeck)

      showHand(shuffledDeck)
			fmt.Print("\nAsk Me to Deal, Shuffle or Exit\n \n ")

		} else if input.Text() == "Shuffle" || input.Text() == "shuffle" {
      fmt.Printf("\n ### You asked me to %s ### \n", input.Text())

			shuffledDeck := shuffleDeck(cards)
      showHand(shuffledDeck)

			fmt.Print("\nAsk Me to Deal, Shuffle or Exit\n \n ")

		} else if input.Text() == "Exit" || input.Text() == "exit" {
			os.Exit(0)
		} else {
      fmt.Printf("\n ### You asked me to %s ### \n", input.Text())
			fmt.Printf("I don't know how to %s\n", input.Text())
		}
	}

}

func showHand(cards *Cards) {
  if len(cards.Hand) == 0 {
     fmt.Printf("\nYour hand is empty. \n \n")
  } else {
    fmt.Printf("\nYour hand now contains: \n \n")
		for i := 0; i < len(cards.Hand); i++ {
      card := cards.Hand[i]
		  fmt.Printf("%s of %s \n", card.Value, card.Suite)
		}
  }
  return
}


func makeDeck() *Cards {

	// points really hung me up. didn't understand why i wasn't able to change the contents of the deck, realized it was because i needed to return a pointer, not a copy

	suites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "Jack", "Queen", "King", "Ace"}
	cards := Cards{}
	for _, suite := range suites {
		for _, value := range values {
			cards.Deck = append(cards.Deck, Card{Suite: suite, Value: value})
		}
	}

	return &cards
}

// next need to have this look for deck.Hand and merge it in first.
func shuffleDeck(cards *Cards) *Cards {

	rand.Seed(time.Now().Unix())
	// first munge stuff back together if necessary
	cards.Deck = append(cards.Deck, cards.Hand...)
	cards.Hand = append(cards.Hand[:0])
	rand.Shuffle(len(cards.Deck), func(i, j int) { cards.Deck[i], cards.Deck[j] = cards.Deck[j], cards.Deck[i] })
	return cards
}

// this should return a hand, not a card. that hand can be a pointer just like in makeDeck.

func dealOneCard(cards *Cards) *Cards {
	//somehow return the deck? keep passing it back and forth? Seems wrong
	//Deal card from top of deck (index 0).
  // This will fail if no cards left. Needs an if statement here to check length of the cards.Deck slice first
	topCard := cards.Deck[0]
	cards.Hand = append(cards.Hand, topCard)

	cards.Deck = append(cards.Deck[1:])
  fmt.Println("\nCards left in Deck:", len(cards.Deck))

	return cards
}
