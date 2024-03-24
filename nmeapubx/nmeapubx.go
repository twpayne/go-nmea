// Package nmeapubx parses u-blox NMEA sentences.
//
// See https://content.u-blox.com/sites/default/files/products/documents/u-blox8-M8_ReceiverDescrProtSpec_UBX-13003221.pdf.
package nmeapubx

import (
	"fmt"

	"github.com/twpayne/go-nmea"
)

var sentenceParserMap = map[int]nmea.SentenceParser{
	0:  nmea.MakeSentenceParser(ParsePosition),
	3:  nmea.MakeSentenceParser(ParseStatus),
	4:  nmea.MakeSentenceParser(ParseTime),
	40: nmea.MakeSentenceParser(ParseRate),
}

type UnknownMsgIDError struct {
	MsgID int
}

func (e UnknownMsgIDError) Error() string {
	return fmt.Sprintf("%d: unknown message id", e.MsgID)
}

func ParseSentence(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	msgID := tok.CommaInt()
	if err := tok.Err(); err != nil {
		return nil, err
	}
	sentenceParser := sentenceParserMap[msgID]
	if sentenceParser != nil {
		return sentenceParser(addr, tok)
	}
	return nil, &UnknownMsgIDError{
		MsgID: msgID,
	}
}

func SentenceParser(addr string) nmea.SentenceParser {
	if addr != "PUBX" {
		return nil
	}
	return ParseSentence
}
