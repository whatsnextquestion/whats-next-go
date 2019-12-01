package book

import "net/url"

// Book has a URL and a sequence of sentences
type Book struct {
	URL       url.URL
	Sentences []Sentence
}

// Sentence is collections of lower case string
type Sentence []string
