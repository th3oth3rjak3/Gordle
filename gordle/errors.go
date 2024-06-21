package gordle

// corpusError is a sentinal error type used when loading or using the corpus is invalid.
type corpusError string

// Error is the implementation of the Error interface by corpusError.
func (e corpusError) Error() string {
	return string(e)
}

// feedbackError defines a sentinal error type used when feedback is invalid.
type feedbackError string

// Error is the implementation of the Error interface by feedbackError.
func (e feedbackError) Error() string {
	return string(e)
}
