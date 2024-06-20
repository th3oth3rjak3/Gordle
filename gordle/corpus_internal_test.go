package gordle

import "testing"

// inCorpus is a helper method to determine if a word is in the corpus.
func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}

	return false
}

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "HALLO", "XAIPE"}
	word := pickWord(corpus)

	if !inCorpus(corpus, word) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}
