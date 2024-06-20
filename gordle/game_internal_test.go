package gordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	type testCase struct {
		input string
		want  []rune
	}

	testCases := map[string]testCase{
		"5 characters in English": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in Japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in Japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			game := New(strings.NewReader(tc.input), []string{string(tc.want)}, 0)

			got := game.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got: %v, want: %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	type testCase struct {
		input []rune
		want  error
	}

	testCases := map[string]testCase{
		"nominal": {
			input: []rune("HELLO"),
			want:  nil,
		},
		"too few characters": {
			input: []rune("HI"),
			want:  errInvalidWordLength,
		},
		"too many characters": {
			input: []rune("HELLO WORLD"),
			want:  errInvalidWordLength,
		},
		"input is empty": {
			input: []rune{},
			want:  errInvalidWordLength,
		},
		"input is nil": {
			input: nil,
			want:  errInvalidWordLength,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			game := New(nil, []string{"VALID"}, 0)
			got := game.validateGuess(tc.input)
			if !errors.Is(got, tc.want) {
				t.Errorf("got: %v, want: %v", got.Error(), tc.want.Error())
			}
		})
	}
}

func TestGameProvideFeedback(t *testing.T) {
	type testCase struct {
		input    []rune
		solution []string
		want     string
	}

	testCases := map[string]testCase{
		"all correct": {
			input:    []rune("HELLO"),
			solution: []string{"HELLO"},
			want:     "💚💚💚💚💚",
		},
		"some correct with missing characters": {
			input:    []rune("JELLO"),
			solution: []string{"HELLO"},
			want:     "⬜💚💚💚💚",
		},
		"none correct": {
			input:    []rune("AAAAA"),
			solution: []string{"BBBBB"},
			want:     "⬜⬜⬜⬜⬜",
		},
		"wrong positions": {
			input:    []rune("OLHEL"),
			solution: []string{"HELLO"},
			want:     "🟡🟡🟡🟡🟡",
		},
		"user input shorter than solution": {
			input:    []rune("SMOL"),
			solution: []string{"SMALL"},
			want:     "💚💚⬜💚⬜",
		},
		"solution shorter than user input": {
			input:    []rune("SMALL"),
			solution: []string{"SMOL"},
			want:     "💚💚⬜💚",
		},
		"empty input": {
			input:    []rune(""),
			solution: []string{"HELLO"},
			want:     "⬜⬜⬜⬜⬜",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			game := New(nil, tc.solution, 0)
			got := game.provideFeedback(tc.input)
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
