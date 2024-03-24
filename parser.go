package nmea

// FIXME add ignore trailing data discipline
// FIXME communicate whether checksum was valid, invalid, or absent

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
)

type ChecksumDiscipline int

const (
	ChecksumDisciplineStrict  ChecksumDiscipline = 0
	ChecksumDisciplineRequire ChecksumDiscipline = 1
	ChecksumDisciplineIgnore  ChecksumDiscipline = 2
	ChecksumDisciplineNever   ChecksumDiscipline = 3
)

type LineEndingDiscipline int

const (
	LineEndingDisciplineStrict  LineEndingDiscipline = 0
	LineEndingDisciplineRequire LineEndingDiscipline = 1
	LineEndingDisciplineIgnore  LineEndingDiscipline = 2
	LineEndingDisciplineNever   LineEndingDiscipline = 3
)

var (
	sentenceRx = regexp.MustCompile(`\A\$([^*]+)\*(.{2})?(\r?\n)?\z`)

	errFraming              = errors.New("framing error")
	errInvalidLineEnding    = errors.New("invalid line ending")
	errMissingChecksum      = errors.New("missing checksum")
	errMissingLineEnding    = errors.New("missing line ending")
	errUnexpectedChecksum   = errors.New("unexpected checksum")
	errUnexpectedLineEnding = errors.New("unexpected line ending")
)

type InvalidChecksumError struct {
	Expected byte
	Got      byte
}

func (e InvalidChecksumError) Error() string {
	return fmt.Sprintf("invalid checksum: expected %02X, got %02X", e.Expected, e.Got)
}

type UnknownAddressError struct {
	Address string
}

func (e UnknownAddressError) Error() string {
	return e.Address + ": unknown address"
}

type Parser struct {
	checksumDiscipline   ChecksumDiscipline
	lineEndingDiscipline LineEndingDiscipline
	sentenceParserFuncs  []func(string) SentenceParser
}

type ParserOption func(*Parser)

func WithChecksumDiscipline(checksumDiscipline ChecksumDiscipline) ParserOption {
	return func(p *Parser) {
		p.checksumDiscipline = checksumDiscipline
	}
}

func WithLineEndingDiscipline(lineEndingDiscipline LineEndingDiscipline) ParserOption {
	return func(p *Parser) {
		p.lineEndingDiscipline = lineEndingDiscipline
	}
}

func WithSentenceParserFunc(sentenceParserFunc func(string) SentenceParser) ParserOption {
	return func(p *Parser) {
		p.sentenceParserFuncs = append(p.sentenceParserFuncs, sentenceParserFunc)
	}
}

func WithSentenceParserMap(sentenceParserMap SentenceParserMap) ParserOption {
	return func(p *Parser) {
		sentenceParserFunc := func(addr string) SentenceParser {
			return sentenceParserMap[addr]
		}
		p.sentenceParserFuncs = append(p.sentenceParserFuncs, sentenceParserFunc)
	}
}

func NewParser(options ...ParserOption) *Parser {
	p := &Parser{
		checksumDiscipline: ChecksumDisciplineStrict,
	}
	for _, option := range options {
		option(p)
	}
	return p
}

func (p *Parser) Parse(data []byte) (Sentence, error) {
	m := sentenceRx.FindSubmatch(data)
	if m == nil {
		return nil, errFraming
	}

	var checksum Optional[byte]
	if len(m[2]) != 0 {
		hexDigit1, _ := hexDigitValue(m[2][0])
		hexDigit2, _ := hexDigitValue(m[2][1])
		checksum = NewOptional(16*byte(hexDigit1) + byte(hexDigit2))
	}
	// FIXME add lax checksum checking which returns any checksum error but also the parsed sentence
	// FIXME add hex check checksum that at least ensures that the checksum contains hex digits
	var calculatedChecksumValue byte
	for _, c := range m[1] {
		calculatedChecksumValue ^= c
	}
	switch p.checksumDiscipline {
	case ChecksumDisciplineStrict:
		switch {
		case !checksum.Valid:
			return nil, errMissingChecksum
		case checksum.Value != calculatedChecksumValue:
			return nil, InvalidChecksumError{
				Expected: calculatedChecksumValue,
				Got:      checksum.Value,
			}
		}
	case ChecksumDisciplineRequire:
		if !checksum.Valid {
			return nil, errMissingChecksum
		}
	case ChecksumDisciplineIgnore:
		// Do nothing.
	case ChecksumDisciplineNever:
		if checksum.Valid {
			return nil, errUnexpectedChecksum
		}
	}

	switch p.lineEndingDiscipline {
	case LineEndingDisciplineStrict:
		if !bytes.Equal(m[3], []byte{'\r', '\n'}) {
			return nil, errInvalidLineEnding
		}
	case LineEndingDisciplineRequire:
		if len(m[3]) == 0 {
			return nil, errMissingLineEnding
		}
	case LineEndingDisciplineIgnore:
		// Do nothing.
	case LineEndingDisciplineNever:
		if len(m[3]) != 0 {
			return nil, errUnexpectedLineEnding
		}
	}

	tokenizer := NewTokenizer(m[1])
	address := tokenizer.String()
	if err := tokenizer.Err(); err != nil {
		return nil, err
	}
	for _, sentenceParserFunc := range p.sentenceParserFuncs {
		if sentenceParser := sentenceParserFunc(address); sentenceParser != nil {
			return sentenceParser(address, tokenizer)
		}
	}
	return nil, UnknownAddressError{
		Address: address,
	}
}

func (p *Parser) ParseString(s string) (Sentence, error) {
	return p.Parse([]byte(s))
}
