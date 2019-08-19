package callpass

import (
	"strings"
)

func GetCallpass(callsign string) int16 {
	call := strings.ToUpper(callsign)

	hash := int16(0x73e2)
	i := 0

	for i < len(call) {
		hash = hash ^ int16(call[i])<<8
		hash = hash ^ int16(call[i+1])
		i += 2
	}
	return int16(hash & 0x7fff)
}

func CheckCallpass(callsign string, pass int16) bool {
	code := GetCallpass(callsign)
	return code == pass
}
