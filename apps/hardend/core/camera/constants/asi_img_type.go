package constants

type ASI_IMG_TYPE int8

const (
	ASI_IMG_RAW8  = iota // Each pixel is an 8-bit (1 byte) gray level
	ASI_IMG_RGB24        // Each pixel consists of RGB, 3 bytes totally (color cameras only)
	ASI_IMG_RAW16        // 2 bytes for every pixel with 65536 gray levels
	ASI_IMG_Y8           // monochrome mode，1 byte every pixel (color cameras only)
	ASI_IMG_END   = -1
)

func (img ASI_IMG_TYPE) ToString() string {
	switch img {
	case ASI_IMG_RAW8:
		return "RAW8"
	case ASI_IMG_RAW16:
		return "RAW16"
	case ASI_IMG_RGB24:
		return "RGB24"
	case ASI_IMG_Y8:
		return "Y8"
	}
	return "type not found"
}
