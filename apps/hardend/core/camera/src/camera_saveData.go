package src

import (
	"core/camera/constants"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
)

func (c *Camera) SaveFrame(destDirectory string) (filename string, err error) {
	switch c.ROI.imgType {
	case constants.ASI_IMG_RAW8:
		return c.saveRAW8(destDirectory)
	case constants.ASI_IMG_RAW16:
		return c.saveRAW16(destDirectory)
	case constants.ASI_IMG_RGB24:
		return c.saveRGB24(destDirectory)
	case constants.ASI_IMG_Y8:
		return c.saveRAW8(destDirectory)
	}
	return filename, nil
}

func (c *Camera) saveRAW8(destDirectory string) (filename string, err error) {
	img := image.NewGray(image.Rect(c.ROI.xStart, c.ROI.yStart,
		c.ROI.xStart+int(c.ROI.width),
		c.ROI.yStart+int(c.ROI.height),
	))
	img.Pix = c.buffer
	if filename, err = savePNG(img, destDirectory); err != nil {
		return "", err
	}
	//if filename, err = saveTXT(c.buffer, destDirectory); err != nil {
	//	return "", err
	//}
	return filename, nil
}

func (c *Camera) saveRAW16(destDirectory string) (filename string, err error) {
	img := image.NewGray16(image.Rect(c.ROI.xStart, c.ROI.yStart,
		c.ROI.xStart+int(c.ROI.width),
		c.ROI.yStart+int(c.ROI.height),
	))
	cBuffLen := len(c.buffer)
	for i := 0; i < cBuffLen; i += 2 {
		tmp := c.buffer[i]
		c.buffer[i] = c.buffer[i+1]
		c.buffer[i+1] = tmp
	}
	img.Pix = c.buffer
	img.Stride = int(c.ROI.width) * 2
	if filename, err = savePNG(img, destDirectory); err != nil {
		return "", err
	}
	//if filename, err = saveTXT(c.buffer, destDirectory); err != nil {
	//	return "", err
	//}
	return filename, nil
}

func (c *Camera) saveRGB24(destDirectory string) (filename string, err error) {
	img := image.NewRGBA(image.Rect(c.ROI.xStart, c.ROI.yStart,
		c.ROI.xStart+int(c.ROI.width),
		c.ROI.yStart+int(c.ROI.height),
	))
	cBuffLen := len(c.buffer)
	newBuff := make([]byte, cBuffLen/3*4)
	j := 0
	for i := 0; i < cBuffLen; i += 3 {
		newBuff[j] = c.buffer[i+2]
		newBuff[j+1] = c.buffer[i+1]
		newBuff[j+2] = c.buffer[i]
		newBuff[j+3] = 0xFF
		j += 4
	}
	img.Pix = newBuff
	img.Stride = int(c.ROI.width) * 4

	if filename, err = savePNG(img, destDirectory); err != nil {
		return "", err
	}

	//if filename, err = saveTXT(c.buffer, destDirectory); err != nil {
	//	return "", err
	//}
	return filename, nil
}

func savePNG(img image.Image, destDirectory string) (string, error) {
	filename := fmt.Sprintf("frame_%d.png", time.Now().UnixNano())
	filepath := destDirectory + "/" + filename
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}

	enc := png.Encoder{CompressionLevel: png.NoCompression}

	if err = enc.Encode(file, img); err != nil {
		return "", err
	}
	return filename, nil
}

func (c *Camera) SaveBufferToTXT(destDirectory string) (string, error) {
	return saveTXT(c.buffer, destDirectory)
}

func saveTXT(buff []byte, destDirectory string) (string, error) {
	filename := fmt.Sprintf("frame_%d.txt", time.Now().UnixNano())
	filepath := destDirectory + "/" + filename
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return "", err
	}

	if _, err = file.Write(buff); err != nil {
		return "", err
	}
	return filename, nil
}
