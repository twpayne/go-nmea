package nmeagps

import "github.com/twpayne/go-nmea"

type DTM struct {
	address  Address
	Datum    string
	SubDatum string
	Lat      float64
	Lon      float64
	Alt      float64
	RefDatum string
}

func ParseDTM(addr string, tok *nmea.Tokenizer) (*DTM, error) {
	var dtm DTM
	dtm.address = NewAddress(addr)
	dtm.Datum = tok.CommaString()
	dtm.SubDatum = tok.CommaString()
	dtm.Lat = 60 * commaLatCommaHemi(tok)
	dtm.Lon = 60 * commaLonCommaHemi(tok)
	dtm.Alt = tok.CommaFloat()
	dtm.RefDatum = tok.CommaString()
	tok.EndOfData()
	return &dtm, tok.Err()
}

func (dtm DTM) Address() nmea.Address {
	return dtm.address
}
