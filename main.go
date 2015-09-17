// A reddit post got me thinking about rocks, paper, scissors...
package main

import (
	"./rockpaperscissors"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for true {
		rockpaperscissors.Turn()
	}
}
