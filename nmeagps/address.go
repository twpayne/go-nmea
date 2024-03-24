package nmeagps

import "regexp"

type Address struct {
	address string
}

var addressRx = regexp.MustCompile(`\AG[A-Z]([A-Z]{3})\z`)

func NewAddress(addr string) Address {
	return Address{
		address: addr,
	}
}

func (a Address) Constellation() byte {
	if len(a.address) < 2 {
		return 0
	}
	return a.address[1]
}

func (a Address) Formatter() string {
	if len(a.address) < 2 {
		return ""
	}
	return string(a.address[2:])
}

func (a Address) Proprietary() bool {
	return false
}

func (a Address) String() string {
	return string(a.address)
}

func (a Address) Talker() string {
	return string(a.address[:min(2, len(a.address))])
}

func (a Address) Valid() bool {
	return addressRx.MatchString(string(a.address))
}
