package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game holds all the information we need to play a game of Gordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New creates a new Gordle Game, which can be used to Play.
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	game := &Game{
		// reader is a buffered reader that reads from the provided io.Reader interface.
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(solution),
		maxAttempts: maxAttempts,
	}

	return game
}

// Play starts Gordle gameplay.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	// ensure the player only has maxAttempts tries to guess the word
	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()
		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
	}

	fmt.Printf("😞 You've lost after %d guess(es)! The solution was: %s. \n", g.maxAttempts, string(g.solution))
}

// ask reads user input until a valid suggestion is made.
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt was invalid: %s\n", err.Error())
		} else {
			return guess
		}
	}
}

// errInvalidWordLength is returned when the guess has the wrong number of characters
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// validateGuess ensures that a player guess is valid.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected: %d, got: %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}

	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string into a list of characters.
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}
