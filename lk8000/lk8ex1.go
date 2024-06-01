package lk8000

import (
	"github.com/twpayne/go-nmea"
)

type LK8EX1 struct {
	nmea.Address
	RawPressure    nmea.Optional[float64]
	Alt            nmea.Optional[float64]
	VarioCMS       nmea.Optional[float64]
	Temperature    nmea.Optional[float64]
	BatteryPercent nmea.Optional[int]
	BatteryVoltage nmea.Optional[float64]
}

func ParseLK8EX1(addr string, tok *nmea.Tokenizer) (*LK8EX1, error) {
	var lk8ex1 LK8EX1
	lk8ex1.Address = nmea.NewAddress(addr)
	if rawPressure := tok.CommaUnsignedInt(); rawPressure != 999999 {
		lk8ex1.RawPressure = nmea.NewOptional(float64(rawPressure) / 100)
	}
	if alt := tok.CommaFloat(); lk8ex1.RawPressure.Valid && alt != 99999 {
		lk8ex1.Alt = nmea.NewOptional(alt)
	}
	if vario := tok.CommaFloat(); vario != 9999 {
		lk8ex1.VarioCMS = nmea.NewOptional(vario)
	}
	if temperature := tok.CommaFloat(); temperature != 99 {
		lk8ex1.Temperature = nmea.NewOptional(temperature)
	}
	switch battery := tok.CommaFloat(); {
	case battery == 999:
	case battery < 1000:
		lk8ex1.BatteryVoltage = nmea.NewOptional(battery)
	default:
		lk8ex1.BatteryPercent = nmea.NewOptional(int(battery) - 1000)
	}
	if !tok.AtEndOfData() {
		tok.Comma()
	}
	tok.EndOfData()
	return &lk8ex1, tok.Err()
}
