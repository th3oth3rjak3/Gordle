package main

import (
	"fmt"
	"os"

	"github.com/th3oth3rjak3/gordle/gordle"
)

// maxAttempts is the maximum number of guesses a player may attempt.
const maxAttempts = 20

func main() {
	corpus, err := gordle.ReadCorpus("./gordle/corpus/english.txt")
	if err != nil {
		exitGordle(err)
	}
	game, err := gordle.New(os.Stdin, corpus, maxAttempts)
	if err != nil {
		exitGordle(err)
	}
	game.Play()
}

func exitGordle(err error) {
	fmt.Printf("Error starting Gordle: %s\n", err.Error())
	os.Exit(1)
}
