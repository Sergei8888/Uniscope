package constants

type ExposureStatus uint8

const (
	ASI_EXP_IDLE    = iota //idle, ready to start exposure
	ASI_EXP_WORKING        //exposure in progress
	ASI_EXP_SUCCESS        // exposure completed successfully, image can be read out
	ASI_EXP_FAILED         // exposure failure, need to restart exposure
)
