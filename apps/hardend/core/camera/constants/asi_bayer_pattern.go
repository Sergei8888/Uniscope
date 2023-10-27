package constants

type ASI_BAYER_PATTERN uint8

const (
	ASI_BAYER_RG = iota
	ASI_BAYER_BG
	ASI_BAYER_GR
	ASI_BAYER_GB
)

func (bp ASI_BAYER_PATTERN) ToString() string {
	switch bp {
	case 0:
		return "RG"
	case 1:
		return "BG"
	case 2:
		return "GR"
	case 3:
		return "GB"
	}
	return "bayer pattern not found"
}
