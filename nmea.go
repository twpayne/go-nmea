// Package nmea parses NMEA sentences.
package nmea

type Sentence interface {
	Address() Addresser
}

type SentenceParser func(string, *Tokenizer) (Sentence, error)

type SentenceParserMap map[string]SentenceParser

func MakeSentenceParser[S Sentence](f func(string, *Tokenizer) (S, error)) SentenceParser {
	return func(addr string, tok *Tokenizer) (Sentence, error) {
		return f(addr, tok)
	}
}
