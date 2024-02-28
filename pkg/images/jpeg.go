package images

import (
	"bytes"
	"image/jpeg"
	"io"
)

func ResizeJPEG(input io.Reader, height int, width int) (io.Reader, error) {
	// Decode the image (from JPG to image.Image):
	src, err := jpeg.Decode(input)
	if err != nil {
		return nil, err
	}

	resizeImageIfNecessary(src, height, width)

	// Encode to output
	output := bytes.NewBuffer([]byte{})
	if err := jpeg.Encode(output, src, &jpeg.Options{Quality: 90}); err != nil {
		return nil, err
	}

	return output, nil
}
