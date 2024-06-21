package gordle

import (
	"errors"
	"slices"
	"testing"
)

func TestFeedbackString(t *testing.T) {
	type testCase struct {
		input []Hint
		want  string
	}

	testCases := map[string]testCase{
		"all correct": {
			input: []Hint{
				CorrectPosition,
				CorrectPosition,
				CorrectPosition,
			},
			want: "💚💚💚",
		},
		"all absent": {
			input: []Hint{
				AbsentCharacter,
				AbsentCharacter,
				AbsentCharacter,
			},
			want: "⬜⬜⬜",
		},
		"all incorrect position": {
			input: []Hint{
				WrongPosition,
				WrongPosition,
				WrongPosition,
			},
			want: "🟡🟡🟡",
		},
		"combined": {
			input: []Hint{
				AbsentCharacter,
				WrongPosition,
				CorrectPosition,
			},
			want: "⬜🟡💚",
		},
		"invalid hint": {
			input: []Hint{
				Hint(10),
			},
			want: "💔",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := Feedback(tc.input).String()
			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}

func TestProvideFeedback(t *testing.T) {
	type testCase struct {
		input    string
		solution string
		want     Feedback
		wantErr  error
	}

	testCases := map[string]testCase{
		"all correct": {
			input:    "HELLO",
			solution: "HELLO",
			want:     []Hint{CorrectPosition, CorrectPosition, CorrectPosition, CorrectPosition, CorrectPosition},
			wantErr:  nil,
		},
		"some correct with missing characters": {
			input:    "JELLO",
			solution: "HELLO",
			want:     []Hint{AbsentCharacter, CorrectPosition, CorrectPosition, CorrectPosition, CorrectPosition},
			wantErr:  nil,
		},
		"none correct": {
			input:    "AAAAA",
			solution: "BBBBB",
			want:     []Hint{AbsentCharacter, AbsentCharacter, AbsentCharacter, AbsentCharacter, AbsentCharacter},
			wantErr:  nil,
		},
		"wrong positions": {
			input:    "OLHEL",
			solution: "HELLO",
			want:     []Hint{WrongPosition, WrongPosition, WrongPosition, WrongPosition, WrongPosition},
			wantErr:  nil,
		},
		"user input shorter than solution": {
			input:    "SMOL",
			solution: "SMALL",
			want:     nil,
			wantErr:  ErrFeedbackInputSolutionLengthMismatch,
		},
		"solution shorter than user input": {
			input:    "SMALL",
			solution: "SMOL",
			want:     nil,
			wantErr:  ErrFeedbackInputSolutionLengthMismatch,
		},
		"empty input": {
			input:    "",
			solution: "HELLO",
			want:     nil,
			wantErr:  ErrFeedbackInputSolutionLengthMismatch,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := ProvideFeedback(tc.input, tc.solution)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("got: %v, want: %v", err, tc.wantErr)
			}

			if !slices.Equal(got, tc.want) {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
