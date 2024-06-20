package gordle

import "fmt"

// Game holds all the information we need to play a game of Gordle.
type Game struct{}

// New creates a new Gordle Game, which can be used to Play.
func New() *Game {
	return &Game{}
}

func (g *Game) Play() {
	// This is where the game will be played.
	// For now, we'll just print a message.
	fmt.Println("Welcome to Gordle!")
	fmt.Printf("Enter a guess:\n")
}
