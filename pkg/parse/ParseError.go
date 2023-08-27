package parse

type ParseError struct {
	message string
}

func NewParseError(message string) *ParseError {
	return &ParseError{
		message: message,
	}
}

func (e *ParseError) Message() string {
	return e.message
}
