package rockpaperscissors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWinnerCheck(t *testing.T) {
	// CheckWinner takes a value of 0,1,2 for rock, paper, scissors respectively
	assert.Contains(t, CheckWinner(0, 0), "Tie", "This should be a tie")         // Player: rock; Computer: rock
	assert.Contains(t, CheckWinner(0, 1), "You Lose", "Player should have lost") // Player: rock; Computer: paper
	assert.Contains(t, CheckWinner(0, 2), "You Win", "Player should have won")   // Player: rock; Computer: scissors

	assert.Contains(t, CheckWinner(1, 0), "You Win", "This should be a tie")    // Player: paper; Computer: rock
	assert.Contains(t, CheckWinner(1, 1), "Tie", "Player should have lost")     // Player: paper; Computer: paper
	assert.Contains(t, CheckWinner(1, 2), "You Lose", "Player should have won") // Player: paper; Computer: scissors

	assert.Contains(t, CheckWinner(2, 0), "You Lose", "This should be a tie")   // Player: scissors; Computer: rock
	assert.Contains(t, CheckWinner(2, 1), "You Win", "Player should have lost") // Player: scissors; Computer: paper
	assert.Contains(t, CheckWinner(2, 2), "Tie", "Player should have won")      // Player: scissors; Computer: scissors

}
