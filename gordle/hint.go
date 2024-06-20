package gordle

import (
	"slices"
	"strings"
)

// hint describes the validity of a character in a word
type hint byte

const (
	// absentCharacter indicates that the character is not present in the solution
	absentCharacter hint = iota
	// wrongPosition indicates that the character is present in the solution, but not at the current location
	wrongPosition
	// correctPosition indicates that the character in the current position is correct.
	correctPosition
)

// String implements the Stringer interface.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜" // White square
	case wrongPosition:
		return "🟡" // Yellow circle
	case correctPosition:
		return "💚" // Green heart
	default:
		// This should never happen.
		return "💔" // Red broken heart
	}
}

// feedback is a list of hints, one per character of a word
type feedback []hint

// String implements the Stringer interface for a slice of hints.
func (fb feedback) String() string {
	sb := &strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

// getRunePositions creates a map of the runes in a string and a slice of the
// index positions of each rune in the input.
// Case-insensitivity is the responsibility of the caller.
func getRunePositions(input string) map[rune][]int {
	counts := make(map[rune][]int, len(input))
	for i, r := range input {
		counts[r] = append(counts[r], i)
	}
	return counts
}

// filter removes any element that matches the predicate
func filter[T any](input []T, predicate func(T) bool) []T {
	return slices.DeleteFunc(input, predicate)
}

// ProvideFeedback checks a userInput against the solution and creates a visual
// representation of correct/incorrect character placement.
// Case-insensitivity is the responsibility of the caller.
func ProvideFeedback(userInput string, solution string) string {
	runePositions := getRunePositions(solution)
	// fb is initialized with absentCharacter values (0)
	fb := make(feedback, len(solution))
	for i, r := range userInput {
		if slices.Contains(runePositions[r], i) {
			fb[i] = correctPosition
			filtered := filter(
				runePositions[r],
				func(input int) bool { return input == i })
			runePositions[r] = filtered
		} else if len(runePositions[r]) > 0 {
			fb[i] = wrongPosition
			runePositions[r] = runePositions[r][1:]
		}
	}

	return fb.String()
}
