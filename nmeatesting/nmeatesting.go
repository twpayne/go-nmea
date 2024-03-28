// Package nmeatesting provides common functionality for testing NMEA sentence parsers.
package nmeatesting

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea"
)

type TestCase struct {
	Skip        string
	Options     []nmea.ParserOption
	S           string
	ExpectedErr error
	Expected    nmea.Sentence
}

func TestSentenceParser(t *testing.T, sentenceParserFunc func(string) nmea.SentenceParser, testCases []TestCase) {
	t.Helper()
	for _, testCase := range testCases {
		t.Run(testCase.S, func(t *testing.T) {
			if testCase.Skip != "" {
				t.Skip(testCase.Skip)
			}
			options := append([]nmea.ParserOption{
				nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
				nmea.WithSentenceParserFunc(sentenceParserFunc),
			}, testCase.Options...)
			parser := nmea.NewParser(options...)
			actual, err := parser.ParseString(testCase.S)
			if testCase.ExpectedErr != nil {
				assert.IsError(t, err, testCase.ExpectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.Expected, actual)
			}
		})
	}
}
