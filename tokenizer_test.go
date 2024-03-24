package nmea_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea"
)

func TestTokenizer(t *testing.T) {
	tok := nmea.NewTokenizer([]byte("GPGGA,123"))
	assert.Equal(t, "GPGGA", tok.String())
	tok.Comma()
	assert.Equal(t, 123, tok.UnsignedInt())
	assert.NoError(t, tok.Err())
}

func TestTokenizer_Float(t *testing.T) {
	for _, tc := range []struct {
		s           string
		expectedErr error
		expected    float64
	}{
		{
			s:        "1",
			expected: 1,
		},
		{
			s:        "1.23",
			expected: 1.23,
		},
		{
			s:        "-1.23",
			expected: -1.23,
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			tok := nmea.NewTokenizer([]byte(tc.s))
			actual := tok.Float()
			err := tok.Err()
			if tc.expectedErr != nil {
				assert.IsError(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestTokenizer_Hex(t *testing.T) {
	for _, tc := range []struct {
		s           string
		expectedErr error
		expected    int
	}{
		{
			s:        "a",
			expected: 0xa,
		},
		{
			s:        "abcdef",
			expected: 0xabcdef,
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			tok := nmea.NewTokenizer([]byte(tc.s))
			actual := tok.Hex()
			err := tok.Err()
			if tc.expectedErr != nil {
				assert.IsError(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestTokenizer_HexBytes(t *testing.T) {
	for _, tc := range []struct {
		s           string
		expectedErr error
		expected    []byte
	}{
		{
			s:        "0123456789ABCDEFabcdef",
			expected: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0xab, 0xcd, 0xef},
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			tok := nmea.NewTokenizer([]byte(tc.s))
			actual := tok.HexBytes()
			err := tok.Err()
			if tc.expectedErr != nil {
				assert.IsError(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestTokenizer_OptionalInt(t *testing.T) {
	for _, tc := range []struct {
		s           string
		expectedErr error
		expected    int
	}{
		{
			s: "",
		},
		{
			s: ",",
		},
		{
			s:        "123",
			expected: 123,
		},
		{
			s:        "-123",
			expected: -123,
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			tok := nmea.NewTokenizer([]byte(tc.s))
			actual := tok.OptionalInt()
			err := tok.Err()
			if tc.expectedErr != nil {
				assert.IsError(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
