package constants

type ASI_CAMERA_MODE uint8

const (
	ASI_MODE_NORMAL = iota
	ASI_MODE_TRIG_SOFT_EDGE
	ASI_MODE_TRIG_RISE_EDGE
	ASI_MODE_TRIG_FALL_EDGE
	ASI_MODE_TRIG_SOFT_LEVEL
	ASI_MODE_TRIG_HIGH_LEVEL
	ASI_MODE_TRIG_LOW_LEVEL
	ASI_MODE_END = -1
)

func (mode ASI_CAMERA_MODE) ToString() string {
	switch mode {
	case ASI_MODE_NORMAL:
		return "Normal"
	case ASI_MODE_TRIG_SOFT_EDGE:
		return "TrigSoftEdge"
	case ASI_MODE_TRIG_RISE_EDGE:
		return "TrigRiseEdge"
	case ASI_MODE_TRIG_FALL_EDGE:
		return "TrigFallEdge"
	case ASI_MODE_TRIG_SOFT_LEVEL:
		return "TrigSoftLevel"
	case ASI_MODE_TRIG_HIGH_LEVEL:
		return "TrigHighLevel"
	case ASI_MODE_TRIG_LOW_LEVEL:
		return "TrigLowLevel"
	}
	return "mode not found"
}
