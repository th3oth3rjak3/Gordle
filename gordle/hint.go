package gordle

import (
	"strings"
)

// Hint describes the validity of a character in a word
type Hint byte

const (
	// AbsentCharacter indicates that the character is not present in the solution
	AbsentCharacter Hint = iota
	// WrongPosition indicates that the character is present in the solution, but not at the current location
	WrongPosition
	// CorrectPosition indicates that the character in the current position is correct.
	CorrectPosition
)

// ErrFeedbackInputSolutionLengthMismatch is returned when the user input and solution are not the same length.
const ErrFeedbackInputSolutionLengthMismatch = feedbackError("user input and solution must be the same length")

// String implements the Stringer interface.
func (h Hint) String() string {
	switch h {
	case AbsentCharacter:
		return "⬜" // White square
	case WrongPosition:
		return "🟡" // Yellow circle
	case CorrectPosition:
		return "💚" // Green heart
	default:
		// This should never happen.
		return "💔" // Red broken heart
	}
}

// feedback is a list of hints, one per character of a word
type Feedback []Hint

// String implements the Stringer interface for a slice of hints.
func (fb Feedback) String() string {
	sb := &strings.Builder{}
	for _, h := range fb {
		_, _ = sb.WriteString(h.String())
	}
	return sb.String()
}

// ProvideFeedback checks a userInput against the solution and creates a visual
// representation of correct/incorrect character placement.
// Case-insensitivity is the responsibility of the caller.
func ProvideFeedback(userInput string, solution string) (Feedback, error) {
	if len(userInput) != len(solution) {
		return nil, ErrFeedbackInputSolutionLengthMismatch
	}

	// zero value for hints is absentCharacter, so no need to set those explicitly
	fb := make(Feedback, len(userInput))
	visited := make([]bool, len(solution))

	// make a first pass to set all the correct positions as priority
	for i, r := range userInput {
		if []rune(solution)[i] == r {
			fb[i] = CorrectPosition
			visited[i] = true
		}
	}

	// make a second pass to set all positions that are in the solution but not in the correct position
	// this will be done from left to right.
	for i, r := range userInput {
		if fb[i] != AbsentCharacter {
			continue
		}

		for j, rn := range solution {
			if !visited[j] && r == rn {
				fb[i] = WrongPosition
				visited[j] = true
				break
			}
		}
	}

	return fb, nil
}
