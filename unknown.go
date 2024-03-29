package nmea

type Unknown struct {
	Address
	Fields []string
}

func ParseUnknown(addr string, tok *Tokenizer) (*Unknown, error) {
	var u Unknown
	u.Address = NewAddress(addr)
	for !tok.AtEndOfData() {
		field := tok.CommaString()
		u.Fields = append(u.Fields, field)
	}
	return &u, tok.Err()
}
