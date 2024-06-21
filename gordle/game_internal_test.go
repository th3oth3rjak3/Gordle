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
			game, err := New(strings.NewReader(tc.input), []string{string(tc.want)}, 0)
			if err != nil {
				t.Errorf("didn't expect an error, got: %s", err.Error())
			}
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
			game, err := New(nil, []string{"VALID"}, 0)
			if err != nil {
				t.Errorf("didn't expect an error, got: %s", err.Error())
			}
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
		wantErr  error
	}

	testCases := map[string]testCase{
		"all correct": {
			input:    []rune("HELLO"),
			solution: []string{"HELLO"},
			want:     "💚💚💚💚💚",
			wantErr:  nil,
		},
		"some correct with missing characters": {
			input:    []rune("JELLO"),
			solution: []string{"HELLO"},
			want:     "⬜💚💚💚💚",
			wantErr:  nil,
		},
		"none correct": {
			input:    []rune("AAAAA"),
			solution: []string{"BBBBB"},
			want:     "⬜⬜⬜⬜⬜",
			wantErr:  nil,
		},
		"wrong positions": {
			input:    []rune("OLHEL"),
			solution: []string{"HELLO"},
			want:     "🟡🟡🟡🟡🟡",
			wantErr:  nil,
		},
		"user input shorter than solution": {
			input:    []rune("SMOL"),
			solution: []string{"SMALL"},
			want:     "",
			wantErr:  ErrFeedbackInputSolutionLengthMismatch,
		},
		"solution shorter than user input": {
			input:    []rune("SMALL"),
			solution: []string{"SMOL"},
			want:     "",
			wantErr:  ErrFeedbackInputSolutionLengthMismatch,
		},
		"empty input": {
			input:    []rune(""),
			solution: []string{"HELLO"},
			want:     "",
			wantErr:  ErrFeedbackInputSolutionLengthMismatch,
		},
		"one letter with correct placement near the end": {
			input:    []rune("SALSA"),
			solution: []string{"PALSY"},
			want:     "⬜💚💚💚⬜",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			game, err := New(nil, tc.solution, 0)
			if err != nil {
				t.Errorf("didn't expect an error, got: %s", err.Error())
			}
			got, err := game.provideFeedback(tc.input)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("got: %v, want: %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
