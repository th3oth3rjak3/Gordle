package gordle_test

import (
	"os"
	"testing"

	"github.com/th3oth3rjak3/gordle/gordle"
)

func TestReadCorpus(t *testing.T) {
	type testCase struct {
		path   string
		length int
		err    error
	}

	testCases := map[string]testCase{
		"English corpus": {
			path:   "./corpus/english.txt",
			length: 2309,
			err:    nil,
		},
		"empty corpus": {
			path:   "./corpus/empty.txt",
			length: 0,
			err:    gordle.ErrCorpusIsEmpty,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(tc.path)
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}
			words, err := gordle.ReadCorpus(data)
			if tc.err != err {
				t.Errorf("got: %v, want: %v", err, tc.err)
			}
			if tc.length != len(words) {
				t.Errorf("got %d words, want %d words", len(words), tc.length)
			}
		})
	}
}
