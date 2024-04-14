package flarm

import (
	"fmt"
	"math"
	"strconv"
)

var (
	AircraftTypes = map[int]string{
		0x0: "reserved",
		0x1: "glider/motor glider (turbo, self-launch, jet) / TMG",
		0x2: "tow plane/tug plane",
		0x3: "helicopter/gyrocopter/rotorcraft",
		0x4: "skydiver, parachute",
		0x5: "drop plane for skydivers",
		0x6: "hang glider (hard)",
		0x7: "paraglider (soft)",
		0x8: "aircraft with reciprocating engine(s)",
		0x9: "aircraft with jet/turboprop engine(s)",
		0xa: "unknown",
		0xb: "balloon (hot, gas, weather, static)",
		0xc: "airship, blimp, zeppelin",
		0xd: "unmanned aerial vehicle (UAV, RPAS, drone)",
		0xe: "(reserved)",
		0xf: "static obstacle",
	}

	AlarmLevels = map[int]string{
		0: "no alarm",
		1: "aircraft or obstacle alarm, 13-18 seconds to impact, Alert Zone alarm, or traffic advisory",
		2: "aircraft or obstacle alarm, 9-12 seconds to impact",
		3: "aircraft or obstacle alarm, 0-8 seconds to impact",
	}

	AlarmTypes = map[int]string{
		0: "no aircraft within range or no-alarm traffic information",
		2: "aircraft alarm",
		3: "obstacle/Alert Zone alarm",
		4: "traffic advisory",
	}

	AlarmLevelTimeToImpact = map[int]int{
		0: math.MaxInt,
		1: 18,
		2: 12,
		3: 8,
	}

	ErrorCodes = map[int]string{
		0x00:  "no error",
		0x11:  "firmware expired",
		0x12:  "firmware update error",
		0x21:  "power",
		0x22:  "UI error",
		0x23:  "audio error",
		0x24:  "ADC error",
		0x25:  "SD card error",
		0x26:  "USB error",
		0x27:  "LED error",
		0x28:  "EEPROM error",
		0x29:  "general hardware error",
		0x2a:  "transponder receiver Mode-C/S/ADS-B unserviceable",
		0x2b:  "EEPROM error",
		0x2c:  "GPIO error",
		0x31:  "GPS communication",
		0x32:  "configuration of GPS module",
		0x33:  "GPS antenna",
		0x41:  "RF communication",
		0x42:  "another FLARM device with the same radio ID is being received",
		0x43:  "wrong ICAO 24-bit address or radio ID",
		0x51:  "communication",
		0x61:  "flash memory",
		0x71:  "pressure sensor",
		0x81:  "obstacle database",
		0x91:  "flight recorder",
		0x93:  "engine-noise recording not possible",
		0x94:  "range analyzer",
		0xa1:  "configuration error",
		0xb1:  "invalid obstacle database license",
		0xb2:  "invalid IGC feature license",
		0xb3:  "invalid AUD feature license",
		0xb4:  "invalid ENL feature license",
		0xb5:  "invalid RFB feature license",
		0xb6:  "invalid TIS feature license",
		0x100: "generic error",
		0x101: "flash file system error",
		0x120: "device is operated outside the designated region",
		0xf1:  "other",
	}

	GPSStatus = map[int]string{
		0: "no GPS reception",
		1: "3d-fix on ground",
		2: "3d-fix when airborne",
	}

	IDTypes = map[int]string{
		0: "random ID",
		1: "official ICAO 24-bit aircraft address",
		2: "fixed FLARM id",
	}

	Severities = map[int]string{
		0: "no error",
		1: "information only",
		2: "functionality may be reduced",
		3: "fatal problem",
	}

	Sources = map[int]string{
		0: "FLARM",
		1: "ADS-B",
		3: "ADS-R",
		4: "TIS-B",
		6: "Mode-S",
	}

	ZoneTypes = map[int]string{
		0x41: "skydiver drop zone",
		0x42: "aerodrome traffic zone",
		0x43: "military firing area",
		0x44: "kite flying zone",
		0x45: "winch launching area",
		0x46: "RC flying area",
		0x47: "UAS flying area",
		0x48: "aerobatic box",
		0x7e: "generic danger area",
		0x7f: "generic prohibited area",
	}
)

func DescribeHex(value int, m map[int]string) string {
	if description, ok := m[value]; ok {
		return fmt.Sprintf("0x%x (%s)", value, description)
	}
	return "0x" + strconv.FormatInt(int64(value), 16)
}

func DescribeInt(value int, m map[int]string) string {
	if description, ok := m[value]; ok {
		return fmt.Sprintf("%d (%s)", value, description)
	}
	return strconv.Itoa(value)
}
