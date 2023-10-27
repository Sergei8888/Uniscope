package constants

type ASI_CONTROL_TYPE int

const (
	ASI_GAIN                = iota //gain
	ASI_EXPOSURE                   //exposure time (microsecond)
	ASI_GAMMA                      //gamma with range 1 to 100 (nominally 50)
	ASI_WB_R                       //red component of white balance
	ASI_WB_B                       //blue component of white balance
	ASI_BRIGHTNESS                 //pixel value offset (a bias not a scale factor)
	ASI_BANDWIDTHOVERLOAD          //The total data transfer rate percentage
	ASI_OVERCLOCK                  //over clock
	ASI_TEMPERATURE                //sensor temperature，10 times the actual temperature
	ASI_FLIP                       //image flip
	ASI_AUTO_MAX_GAIN              //maximum gain when auto adjust
	ASI_AUTO_MAX_EXP               //maximum exposure time when auto adjust，unit is micro seconds
	ASI_AUTO_MAX_BRIGHTNESS        //target brightness when auto adjust
	ASI_HARDWARE_BIN               //hardware binning of pixels
	ASI_HIGH_SPEED_MODE            //high speed mode
	ASI_COOLER_POWER_PERC          //cooler power percent(only cool src)
	ASI_TARGET_TEMP                //sensor's target temperature(only cool src)，don't multiply by 10
	ASI_COOLER_ON                  //open cooler (only cool src)
	ASI_MONO_BIN                   //lead to a smaller grid at software bin mode for color src
	ASI_FAN_ON                     //only cooled src has fan
	ASI_PATTERN_ADJUST             //currently only supported by 1600 mono src
	ASI_ANTI_DEW_HEATER
)
