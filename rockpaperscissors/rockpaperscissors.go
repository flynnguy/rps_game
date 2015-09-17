// A reddit post got me thinking about rocks, paper, scissors...
package rockpaperscissors

import (
	"fmt"
	"math/rand"
	"strings"
)

var RPSInput = map[string]int{
	"r": 0,
	"p": 1,
	"s": 2,
}

// CheckWinner compares two int values that a user has entered based on RPSInput var:
// "r": 0,
// "p": 1,
// "s": 2,
func CheckWinner(player1, player2 int) string {
	winner := (3 + (player1 - player2)) % 3
	result := "Something odd happened"
	switch winner {
	case 1:
		result = "You Win!"
	case 2:
		result = "You Lose."
	case 0:
		result = "It's a Tie!"
	}
	return result
}

// GetHand prompts the user to pick rock, paper or scissors
func GetHand() int {
	fmt.Print("Pick [r]ock, [p]aper, or [s]cissors:   ")
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	return RPSInput[input]
}

// GetComputerHand just gets a random number between 1-3.
func GetComputerHand() int {
	return rand.Intn(3 - 0)
}

// Turn get's a user's prompt as well as get a random value from the computer and outputs the winner
func Turn() {
	player1 := GetHand()
	computer := GetComputerHand()
	player1Hand := ""
	computerHand := ""
	for key, value := range RPSInput {
		if value == player1 {
			player1Hand = key
		}
		if value == computer {
			computerHand = key
		}
	}
	fmt.Printf("You picked: %s, the computer picked: %s\n", player1Hand, computerHand)
	fmt.Println(CheckWinner(player1, computer))
}

/*
// Example main()
package main

import (
	"flynnguy.com/rockpaperscissors"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for true {
		rockpaperscissors.Turn()
	}
}
*/
