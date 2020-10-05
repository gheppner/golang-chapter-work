package main

import (
	"fmt"
	// "reflect"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	//  "net/url"
	"strings"
	//  "github.com/goombaio/namegenerator"
)

type Cards struct {
	Hand []Card
	Deck []Card
}

type Card struct {
	Suite string
	Value string
}

// new game creates a new player.
// games have N players associated with them.
type Games struct {
	Ids     []int
	Players Player
}

type Player struct {
	Username string
}

// list of players
// list of games (which contains list of players
// player is a user and uid
//

var cards Cards
var gameExist bool

//var games Games

func main() {
	cards := makeDeck()
	shuffledDeck := shuffleDeck(cards)
	games := initGames()
	fmt.Sprintln(shuffledDeck)
	//default game is id 0

	newGameHandler := func(w http.ResponseWriter, req *http.Request) {
		newGame(games)
		newId := games.Ids[len(games.Ids)-1]
		output := fmt.Sprintf("Game ID is %d \n", newId)
		io.WriteString(w, output)
	}

	dealOneCardHandler := func(w http.ResponseWriter, req *http.Request) {
		url := req.RequestURI
		// look for game
		for _, v := range games.Ids {
			if url == fmt.Sprintf("/deal/%d", v) {
				fmt.Printf("Found game %d", v)
				gameExist = true
				break
			}
		}

		if !gameExist {
			newGame(games)
		}

		dealOneCard(shuffledDeck)
		if len(cards.Hand) == 0 {
			io.WriteString(w, "Your Hand Is Empty")
		} else {
			io.WriteString(w, "Your Hand Now Contains: \n")
			yourHand := []string{}
			for i := 0; i < len(cards.Hand); i++ {
				card := cards.Hand[i]
				prettyCard := fmt.Sprintf("\n%s of %s", card.Value, card.Suite)
				yourHand = append(yourHand, prettyCard)
			}
			returnHand := strings.Join(yourHand, "\n")
			io.WriteString(w, returnHand)
		}
	}

	shuffleHandler := func(w http.ResponseWriter, req *http.Request) {
		shuffleDeck(cards)
		if len(cards.Hand) == 0 {
			io.WriteString(w, "Your Hand Is Empty")
		} else {
			io.WriteString(w, "Your Hand Now Contains: \n")
			yourHand := []string{}
			for i := 0; i < len(cards.Hand); i++ {
				card := cards.Hand[i]
				prettyCard := fmt.Sprintf("\n%s of %s", card.Value, card.Suite)
				yourHand = append(yourHand, prettyCard)
			}
			returnHand := strings.Join(yourHand, "\n")
			io.WriteString(w, returnHand)
		}
	}

	http.HandleFunc("/deal/", dealOneCardHandler)
	http.HandleFunc("/shuffle", shuffleHandler)
	http.HandleFunc("/newgame", newGameHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//func initPlayer() *PlayerList {
//  seed := time.Now().UTC().UnixNano()
//  nameGenerator := namegenerator.NewNameGenerator(seed)
//  name := nameGenerator.Generate()
//  playerlist := PlayerList{}
//  playerList.Players = append(playerList.Players, Player{Username: name})
//  return &playerlist
//}

func newGame(games *Games) *Games {
	currentGame := games.Ids[len(games.Ids)-1]
	currentGame++
	games.Ids = append(games.Ids, currentGame)
	fmt.Println(games.Ids)
	return games
}

// do i have to make a function to return this pointer?
func initGames() *Games {
	games := Games{}
	games.Ids = append(games.Ids, 0)
	return &games
}

func makeDeck() *Cards {
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
