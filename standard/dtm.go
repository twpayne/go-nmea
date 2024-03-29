package standard

import "github.com/twpayne/go-nmea"

type DTM struct {
	nmea.Address
	Datum    string
	SubDatum string
	Lat      float64
	Lon      float64
	Alt      float64
	RefDatum string
}

func ParseDTM(addr string, tok *nmea.Tokenizer) (*DTM, error) {
	var dtm DTM
	dtm.Address = nmea.NewAddress(addr)
	dtm.Datum = tok.CommaString()
	dtm.SubDatum = tok.CommaString()
	dtm.Lat = 60 * tok.CommaLatCommaHemi()
	dtm.Lon = 60 * tok.CommaLonCommaHemi()
	dtm.Alt = tok.CommaFloat()
	dtm.RefDatum = tok.CommaString()
	tok.EndOfData()
	return &dtm, tok.Err()
}
