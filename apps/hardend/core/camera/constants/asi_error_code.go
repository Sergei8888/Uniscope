package constants

type ASI_ERROR_CODE uint8

const (
	ASI_SUCCESS                    = iota // operation was successful
	ASI_ERROR_INVALID_INDEX               //no src connected or index value out of boundary
	ASI_ERROR_INVALID_ID                  //invalid ID
	ASI_ERROR_INVALID_CONTROL_TYPE        //invalid control type
	ASI_ERROR_CAMERA_CLOSED               //src didn't open
	ASI_ERROR_CAMERA_REMOVED              //failed to find the src, maybe the src has been removed
	ASI_ERROR_INVALID_PATH                //cannot find the path of the file
	ASI_ERROR_INVALID_FILEFORMAT
	ASI_ERROR_INVALID_SIZE     //wrong video format size
	ASI_ERROR_INVALID_IMGTYPE  //unsupported image format
	ASI_ERROR_OUTOF_BOUNDARY   //the startpos is outside the image boundary
	ASI_ERROR_TIMEOUT          //timeout
	ASI_ERROR_INVALID_SEQUENCE //stop capture first
	ASI_ERROR_BUFFER_TOO_SMALL //buffer size is not big enough
	ASI_ERROR_VIDEO_MODE_ACTIVE
	ASI_ERROR_EXPOSURE_IN_PROGRESS
	ASI_ERROR_GENERAL_ERROR //general error, eg: value is out of valid range
	ASI_ERROR_END
)
