package Sky

import (
	"errors"
	"fmt"
	"math"
)

func HexDigitToDecimal(hexDigit byte) (int, error) {
	if hexDigit > 47 && hexDigit < 59 {
		return int(hexDigit) - 48, nil
	}
	if hexDigit > 64 && hexDigit < 71 {
		return int(hexDigit) - 55, nil
	}
	return 0, errors.New(fmt.Sprintf("invalid hex, got: %d", hexDigit))
}

func DecomposeDegreesMinutesSecs(degreesF float64) (degrees, minutes, secs int) {
	degreesF, minutesF := math.Modf(degreesF)
	minutesF *= 60
	minutesF, secsF := math.Modf(minutesF)
	secsF *= 60
	return int(degreesF), int(minutesF), int(math.Round(secsF))
}

func DecimalDigitToHex(decimalNum int) (byte, error) {
	if decimalNum > -1 && decimalNum < 10 {
		return byte(decimalNum + 48), nil
	}
	if decimalNum > 9 && decimalNum < 16 {
		return byte(decimalNum + 55), nil
	}
	return 0, errors.New(fmt.Sprintf("invalid decimal, got: %d", decimalNum))
}
