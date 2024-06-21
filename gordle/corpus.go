package gordle

import (
	"math/rand"
	"strings"
)

// ErrCorpusIsEmpty is returned when the corpus is empty.
const ErrCorpusIsEmpty = corpusError("corpus is empty")

// ReadCorpus reads the file at the given path and returns a list of words.
func ReadCorpus(data []byte) ([]string, error) {
	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	// we expect the words to be a line or space-separated list.
	words := strings.Fields(string(data))
	return words, nil
}

// pickWord uses a pseudo-random number generator to select a random word from the corpus.
func pickWord(corpus []string) string {
	index := rand.Intn(len(corpus))
	return corpus[index]
}
