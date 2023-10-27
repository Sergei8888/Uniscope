package celestron

import (
	Sky "core/mount/src/Coordinates/Sky"
	"strconv"
)

func (m *MountCelestron) GetRaDec() (*Sky.EquatorialCS, error) {
	_, degrees, err := m.getCoordinates('E')
	if err != nil {
		return nil, err
	}

	ra, err := Sky.NewRightAscension(Sky.DecomposeDegreesMinutesSecs(degrees[0]))
	if err != nil {
		return nil, err
	}

	dec, err := Sky.NewDeclination(Sky.DecomposeDegreesMinutesSecs(degrees[1]))
	if err != nil {
		return nil, err
	}
	return Sky.NewEquatorialCS(ra, dec), nil
}

func (m *MountCelestron) GetAzmAlt() (*Sky.HorizontalCS, error) {
	_, degrees, err := m.getCoordinates('Z')
	if err != nil {
		return nil, err
	}

	azm, err := Sky.NewAzimuth(Sky.DecomposeDegreesMinutesSecs(degrees[0]))
	if err != nil {
		return nil, err
	}

	alt, err := Sky.NewAltitude(Sky.DecomposeDegreesMinutesSecs(degrees[1]))
	if err != nil {
		return nil, err
	}

	return Sky.NewHorizontalCS(azm, alt), nil
}

func (m *MountCelestron) getCoordinates(_type byte) (percentage []float64, degrees []float64, err error) {
	_, err = m.Communicator().Write([]byte{_type})
	if err != nil {
		return nil, nil, err
	}
	answer, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return nil, nil, err
	}

	//var firstPosition float64
	//for i := 1; i < 5; i++ {
	//    num, err := Sky.HexDigitToDecimal(answer[i])
	//    if err != nil {
	//        return nil, nil, err
	//    }
	//    firstPosition += float64(num)
	//    firstPosition *= 16
	//}
	//firstPosition = firstPosition / 16 / 65536

	//var secondPosition float64
	//for i := 6; i < 10; i++ {
	//    num, err := Sky.HexDigitToDecimal(answer[i])
	//    if err != nil {
	//        return nil, nil, err
	//    }
	//    secondPosition += float64(num)
	//    secondPosition *= 16
	//}

	firstPosition, err := strconv.ParseInt(string(answer[1:5]), 16, 32)
	if err != nil {
		return nil, nil, err
	}
	firstPositionF := float64(firstPosition) / 65536

	secondPosition, err := strconv.ParseInt(string(answer[6:10]), 16, 32)
	if err != nil {
		return nil, nil, err
	}
	secondPositionF := float64(secondPosition) / 65536

	percentage = []float64{firstPositionF, secondPositionF}
	degrees = []float64{firstPositionF * 360, secondPositionF * 360}
	return
}

func (m *MountCelestron) GetPreciseRaDec() (*Sky.EquatorialCS, error) {
	_, degrees, err := m.getPreciseCoordinates('e')
	if err != nil {
		return nil, err
	}

	ra, err := Sky.NewRightAscension(Sky.DecomposeDegreesMinutesSecs(degrees[0]))
	if err != nil {
		return nil, err
	}

	dec, err := Sky.NewDeclination(Sky.DecomposeDegreesMinutesSecs(degrees[1]))
	if err != nil {
		return nil, err
	}

	return Sky.NewEquatorialCS(ra, dec), nil
}

func (m *MountCelestron) GetPreciseAzmAlt() (*Sky.HorizontalCS, error) {
	_, degrees, err := m.getPreciseCoordinates('z')
	if err != nil {
		return nil, err
	}

	azm, err := Sky.NewAzimuth(Sky.DecomposeDegreesMinutesSecs(degrees[0]))
	if err != nil {
		return nil, err
	}

	alt, err := Sky.NewAltitude(Sky.DecomposeDegreesMinutesSecs(degrees[1]))
	if err != nil {
		return nil, err
	}

	return Sky.NewHorizontalCS(azm, alt), nil
}

func (m *MountCelestron) getPreciseCoordinates(_type byte) (percentage []float64, degrees []float64, err error) {
	_, err = m.Communicator().Write([]byte{_type})
	if err != nil {
		return nil, nil, err
	}
	answer, err := m.Communicator().ReadAnswerBytes()
	if err != nil {
		return nil, nil, err
	}

	firstPosition, err := strconv.ParseInt(string(answer[1:7]), 16, 32)
	if err != nil {
		return nil, nil, err
	}
	firstPositionF := float64(firstPosition) / 16777216

	secondPosition, err := strconv.ParseInt(string(answer[10:16]), 16, 32)
	if err != nil {
		return nil, nil, err
	}
	secondPositionF := float64(secondPosition) / 16777216

	percentage = []float64{firstPositionF, secondPositionF}
	degrees = []float64{firstPositionF * 360, secondPositionF * 360}
	return
}
