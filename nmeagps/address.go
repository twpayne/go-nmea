package nmeagps

// FIXME should this be an opaque type? provides forward compatibility at the expense of fluency in tests

import "regexp"

type Address string

var addressRx = regexp.MustCompile(`\AG[A-Z]([A-Z]{3})\z`)

func (a Address) Constellation() byte {
	if len(a) < 2 {
		return 0
	}
	return a[1]
}

func (a Address) Formatter() string {
	if len(a) < 2 {
		return ""
	}
	return string(a[2:])
}

func (a Address) Proprietary() bool {
	return false
}

func (a Address) String() string {
	return string(a)
}

func (a Address) Talker() string {
	return string(a[:min(2, len(a))])
}

func (a Address) Valid() bool {
	return addressRx.MatchString(string(a))
}
