package nmea

// FIXME add MakeSentenceParserMap that uses reflection instead of requiring use of MakeSentenceParser all the time

// FIXME check NMEA specs on whitespace handling; should we skip/ignore it or
// include it in parsing? sparkfun has a lot of whitespace, but maybe they had
// an inattentive documentation writer

type Sentence interface {
	Address() Address
}

type SentenceParser func(string, *Tokenizer) (Sentence, error)

type SentenceParserMap map[string]SentenceParser

func MakeSentenceParser[S Sentence](f func(string, *Tokenizer) (S, error)) SentenceParser {
	return func(address string, tokenizer *Tokenizer) (Sentence, error) {
		return f(address, tokenizer)
	}
}
