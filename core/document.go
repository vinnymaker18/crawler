package core

// All supported file types. We employ algorithms to extract natural
// language text from urls/files of these types.
const (
	DOC_TEXT = 0
	DOC_HTML = 1
	DOC_PDF  = 2
)

// RawDocument is any web page/file that has text content e.g., web pages,
// pdf files etc...
type RawDocument struct {
	// Document content stored as a string.
	Content string
}
