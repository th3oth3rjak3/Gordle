package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/th3oth3rjak3/gordle/gordle"
)

// maxAttempts is the maximum number of guesses a player may attempt.
const maxAttempts = 20

var (
	//go:embed gordle/corpus/english.txt
	f embed.FS
)

func main() {
	englishCorpus, err := f.ReadFile("gordle/corpus/english.txt")
	if err != nil {
		exitGordle(err)
	}
	corpus, err := gordle.ReadCorpus(englishCorpus)
	if err != nil {
		exitGordle(err)
	}
	game, err := gordle.New(os.Stdin, corpus, maxAttempts)
	if err != nil {
		exitGordle(err)
	}

	if err = game.Play(); err != nil {
		exitGordle(err)
	}

}

func exitGordle(err error) {
	fmt.Printf("Unexpected error occurred: %s\n", err.Error())
	os.Exit(1)
}
