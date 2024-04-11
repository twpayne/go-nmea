package flarm

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type PFLAO struct {
	nmea.Address
	AlarmLevel    int
	Inside        int
	Lat           int
	Lon           int
	Radius        int
	Bottom        int
	Top           int
	ActivityLimit time.Time
	ID            int
	IDType        int
	ZoneType      int
}

func ParsePFLAO(addr string, tok *nmea.Tokenizer) (*PFLAO, error) {
	var pflao PFLAO
	pflao.Address = nmea.NewAddress(addr)
	pflao.AlarmLevel = tok.CommaUnsignedInt()
	pflao.Inside = tok.CommaUnsignedInt()
	pflao.Lat = tok.CommaInt()
	pflao.Lon = tok.CommaInt()
	pflao.Radius = tok.CommaUnsignedInt()
	pflao.Bottom = tok.CommaInt()
	pflao.Top = tok.CommaInt()
	pflao.ActivityLimit = time.Unix(int64(tok.CommaUnsignedInt()), 0).UTC()
	pflao.ID = tok.CommaHex()
	pflao.IDType = tok.CommaUnsignedInt()
	pflao.ZoneType = tok.CommaHex()
	return &pflao, tok.Err()
}
