// Package nmeatest provides common functionality for testing NMEA sentence parsers.
package nmeatest

import (
	"slices"
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

func TestSentenceParserFunc(t *testing.T, options []nmea.ParserOption, testCases []TestCase) {
	t.Helper()
	for _, testCase := range testCases {
		t.Run(testCase.S, func(t *testing.T) {
			if testCase.Skip != "" {
				t.Skip(testCase.Skip)
			}
			testCaseOptions := slices.Clone(options)
			testCaseOptions = append(testCaseOptions, testCase.Options...)
			parser := nmea.NewParser(testCaseOptions...)
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
