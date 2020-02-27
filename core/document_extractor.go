package core

// An error encountered when trying to parse a document. For now, it's just
// a dumb type, doesn't hold any info about the actual error.
type ParseError struct {
}

func (err *ParseError) Error() string {
	return "Parsing error"
}

// Parses a raw document and extracts natural language text from it.
// The extracted text will have an encoding
func ParseRawDocument(rawDoc RawDocument) (string, error) {
	// TODO - Implement this.
	return "", &ParseError{}
}
