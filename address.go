package nmea

// FIXME add Valid method?

import (
	"fmt"
	"strings"
)

type Addresser interface {
	fmt.Stringer
	Formatter() string
	Proprietary() bool
	Talker() string
}

type Address struct {
	address string
}

func NewAddress(addr string) Address {
	return Address{
		address: addr,
	}
}

func (a Address) Formatter() string {
	switch {
	case a.Proprietary():
		return a.address
	case len(a.address) <= 2:
		return ""
	default:
		return a.address[2:]
	}
}

func (a Address) Proprietary() bool {
	return strings.HasPrefix(a.address, "P")
}

func (a Address) String() string {
	return a.address
}

func (a Address) Talker() string {
	if a.Proprietary() {
		return a.address[:min(4, len(a.address))]
	}
	return a.address[:min(2, len(a.address))]
}
