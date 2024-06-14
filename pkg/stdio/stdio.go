package stdio

import (
	"os"
)

func IsStdinPresent() (present bool) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		present = false

		return
	}

	isPipedFromChrDev := (stat.Mode() & os.ModeCharDevice) == 0
	isPipedFromFIFO := (stat.Mode() & os.ModeNamedPipe) != 0

	present = isPipedFromChrDev || isPipedFromFIFO

	return
}
