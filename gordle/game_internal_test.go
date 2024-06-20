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
			game := New(strings.NewReader(tc.input), string(tc.want), 0)

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
			game := New(nil, "VALID", 0)
			got := game.validateGuess(tc.input)
			if !errors.Is(got, tc.want) {
				t.Errorf("got: %v, want: %v", got.Error(), tc.want.Error())
			}
		})
	}
}
