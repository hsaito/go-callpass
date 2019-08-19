package callpass

import (
	"strings"
)

func GetCallPass(callsign string) int16 {
	call := strings.ToUpper(callsign)

	hash := int16(0x73e2)
	i := 0

	for i < len(call) {
		hash = hash ^ int16(call[i])<<8
		if i+1 < len(call) {
			hash = hash ^ int16(call[i+1])
		}
		i += 2
	}
	return int16(hash & 0x7fff)
}

func CheckCallPass(callsign string, pass int16) bool {
	code := GetCallPass(callsign)
	return code == pass
}
