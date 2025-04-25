// Package nmea parses NMEA sentences.
package nmea

type Sentence interface {
	GetAddress() Address
}

type SentenceParser func(string, *Tokenizer) (Sentence, error)

type SentenceParserMap map[string]SentenceParser

func MakeSentenceParser[S Sentence](f func(string, *Tokenizer) (S, error)) SentenceParser {
	return func(addr string, tok *Tokenizer) (Sentence, error) {
		return f(addr, tok)
	}
}

// Frame returns data framed with NMEA characters, a checksum, and end of line
// characters.
func Frame(data []byte) []byte {
	result := make([]byte, len(data)+6)
	result[0] = '$'
	copy(result[1:], data)
	var checksum byte
	for _, b := range data {
		checksum ^= b
	}
	result[len(data)+1] = '*'
	result[len(data)+2] = hexDigit(checksum >> 4)
	result[len(data)+3] = hexDigit(checksum & 0xf)
	result[len(data)+4] = '\r'
	result[len(data)+5] = '\n'
	return result
}

// FrameString returns data framed with NMEA characters, a checksum, and end of
// line characters.
func FrameString(data string) string {
	return string(Frame([]byte(data)))
}

func hexDigit(value byte) byte {
	switch {
	case value < 0xa:
		return '0' + value
	case value < 0x10:
		return 'A' + value - 0xa
	default:
		panic("digit out of range")
	}
}
