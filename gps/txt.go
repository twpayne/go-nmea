package gps

import "github.com/twpayne/go-nmea"

type TXT struct {
	address Address
	NumMsg  int
	MsgNum  int
	MsgType int
	Text    string
}

func ParseTXT(addr string, tok *nmea.Tokenizer) (*TXT, error) {
	var txt TXT
	txt.address = NewAddress(addr)
	txt.NumMsg = tok.CommaInt()
	txt.MsgNum = tok.CommaInt()
	txt.MsgType = tok.CommaInt()
	txt.Text = tok.CommaString()
	tok.EndOfData()
	return &txt, tok.Err()
}

func (txt TXT) Address() nmea.Addresser {
	return txt.address
}
