package gordle

import (
	"slices"
	"testing"
)

func TestFeedbackString(t *testing.T) {
	type testCase struct {
		input []hint
		want  string
	}

	testCases := map[string]testCase{
		"all correct": {
			input: []hint{
				correctPosition,
				correctPosition,
				correctPosition,
			},
			want: "💚💚💚",
		},
		"all absent": {
			input: []hint{
				absentCharacter,
				absentCharacter,
				absentCharacter,
			},
			want: "⬜⬜⬜",
		},
		"all incorrect position": {
			input: []hint{
				wrongPosition,
				wrongPosition,
				wrongPosition,
			},
			want: "🟡🟡🟡",
		},
		"combined": {
			input: []hint{
				absentCharacter,
				wrongPosition,
				correctPosition,
			},
			want: "⬜🟡💚",
		},
		"invalid hint": {
			input: []hint{
				hint(10),
			},
			want: "💔",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := feedback(tc.input).String()
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}

func TestGetRunePositions(t *testing.T) {
	type testCase struct {
		input string
		want  map[rune][]int
	}

	testCases := map[string]testCase{
		"no repeating characters": {
			input: "WORLD",
			want: map[rune][]int{
				'W': {0},
				'O': {1},
				'R': {2},
				'L': {3},
				'D': {4},
			},
		},
		"all same character": {
			input: "AAA",
			want: map[rune][]int{
				'A': {0, 1, 2},
			},
		},
		"empty input": {
			input: "",
			want:  map[rune][]int{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := getRunePositions(tc.input)
			for k, v := range tc.want {
				if !slices.Equal(got[k], v) {
					t.Errorf("got: %v, want: %v", got, tc.want)
				}
			}
		})
	}
}

func TestProvideFeedback(t *testing.T) {
	type testCase struct {
		input    string
		solution string
		want     string
	}

	testCases := map[string]testCase{
		"all correct": {
			input:    "HELLO",
			solution: "HELLO",
			want:     "💚💚💚💚💚",
		},
		"some correct with missing characters": {
			input:    "JELLO",
			solution: "HELLO",
			want:     "⬜💚💚💚💚",
		},
		"none correct": {
			input:    "AAAAA",
			solution: "BBBBB",
			want:     "⬜⬜⬜⬜⬜",
		},
		"wrong positions": {
			input:    "OLHEL",
			solution: "HELLO",
			want:     "🟡🟡🟡🟡🟡",
		},
		"user input shorter than solution": {
			input:    "SMOL",
			solution: "SMALL",
			want:     "💚💚⬜💚⬜",
		},
		"solution shorter than user input": {
			input:    "SMALL",
			solution: "SMOL",
			want:     "💚💚⬜💚",
		},
		"empty input": {
			input:    "",
			solution: "HELLO",
			want:     "⬜⬜⬜⬜⬜",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := ProvideFeedback(tc.input, tc.solution)
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
