package nmeapubx

type Address struct {
	address string
}

func NewAddress(addr string) Address {
	return Address{
		address: addr,
	}
}

func (a Address) Formatter() string {
	return ""
}

func (a Address) Proprietary() bool {
	return true
}

func (a Address) String() string {
	return a.address
}

func (a Address) Talker() string {
	return a.address[:max(4, len(a.address))]
}
