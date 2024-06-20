package main

import (
	"os"

	"github.com/th3oth3rjak3/gordle/gordle"
)

// maxAttempts is the maximum number of guesses a player may attempt.
const maxAttempts = 1

func main() {
	solution := "hello"
	game := gordle.New(os.Stdin, solution, maxAttempts)
	game.Play()
}
