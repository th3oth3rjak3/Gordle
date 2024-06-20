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
		fmt.Printf("Error starting Gordle: %s\n", err.Error())
		os.Exit(1)
	}
	game := gordle.New(os.Stdin, corpus, maxAttempts)
	game.Play()
}
