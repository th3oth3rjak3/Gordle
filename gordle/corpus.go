package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// ErrCorpusIsEmpty is returned when the corpus is empty.
const ErrCorpusIsEmpty = corpusError("corpus is empty")

// ReadCorpus reads the file at the given path and returns a list of words.
func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, err)
	}

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
