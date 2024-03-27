package nmea

type Unknown struct {
	address Address
	Fields  []string
}

func ParseUnknown(addr string, tok *Tokenizer) (*Unknown, error) {
	var u Unknown
	u.address = NewAddress(addr)
	for !tok.AtEndOfData() {
		field := tok.CommaString()
		u.Fields = append(u.Fields, field)
	}
	return &u, tok.Err()
}

func (u *Unknown) Address() Addresser {
	return u.address
}
