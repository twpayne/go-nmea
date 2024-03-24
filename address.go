package nmea

import "fmt"

type Address interface {
	fmt.Stringer
	Formatter() string
	Proprietary() bool
	Talker() string
}
