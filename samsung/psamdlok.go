package samsung

import "github.com/twpayne/go-nmea"

type PSAMDLOK struct {
	nmea.Address
	Unknown1 int
	Unknown2 int
	Unknown3 int
}

func ParsePSAMDLOK(addr string, tok *nmea.Tokenizer) (*PSAMDLOK, error) {
	var psamdlok PSAMDLOK
	psamdlok.Address = nmea.NewAddress(addr)
	psamdlok.Unknown1 = tok.CommaInt()
	psamdlok.Unknown2 = comma0xHex(tok)
	psamdlok.Unknown3 = comma0xHex(tok)
	tok.EndOfData()
	return &psamdlok, tok.Err()
}

func comma0xHex(tok *nmea.Tokenizer) int {
	if tok.Err() != nil {
		return 0
	}
	tok.Comma()
	tok.LiteralString("0x")
	return tok.Hex()
}
