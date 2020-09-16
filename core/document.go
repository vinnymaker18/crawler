package core

// All supported file types. We use algorithms to mine natural
// language text from urls/files of these types.
const (
	DOC_TEXT = 0
	DOC_HTML = 1
	DOC_PDF  = 2
)

// RawDocument is any fetched web page/file. Such documents are used
// to mine interesting text.
type RawDocument struct {
	// Detected mime type for this document. Mime type is sniffed using
	// upto first 512 bytes of data.
	mimeType string

	// Raw stream of bytes that make up a file/doc
	content []byte
}

// Article is the important(and interesting to analyse) textual content
// mined from a raw doc.
type Article struct {
}
