package reporters

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	_ "image/png" // Register PNG decoder
	"strings"

	"golang.org/x/image/draw"
)

// OptimizeAndEncodeScreenshot takes raw image data (expected to be PNG),
// resizes it if necessary, compresses it as JPEG, and returns a Base64 Data URI.
func OptimizeAndEncodeScreenshot(rawData []byte) (string, error) {
	src, _, err := image.Decode(bytes.NewReader(rawData))
	if err != nil {
		return "", err
	}

	dst := applyMaxWidthConstraint(src)

	encoded, err := encodeInJPEGWithCompression(dst)
	if err != nil {
		return "", err
	}

	return createURI(encoded), nil
}

func createURI(encoded string) string {
	var builder strings.Builder
	builder.WriteString("data:image/jpeg;base64,")
	builder.WriteString(encoded)

	return builder.String()
}

func encodeInJPEGWithCompression(dst image.Image) (string, error) {
	const imageQuality = 70
	var buf bytes.Buffer
	opts := &jpeg.Options{Quality: imageQuality}
	if errJPEGEncode := jpeg.Encode(&buf, dst, opts); errJPEGEncode != nil {
		return "", errJPEGEncode
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encoded, nil
}

func applyMaxWidthConstraint(src image.Image) image.Image {
	const maxWidth = 1280

	bounds := src.Bounds()

	width := bounds.Dx()
	if width <= maxWidth {
		return src
	}

	height := bounds.Dy()

	ratio := float64(maxWidth) / float64(width)
	newHeight := int(float64(height) * ratio)
	dstRect := image.Rect(0, 0, maxWidth, newHeight)
	resizedImg := image.NewRGBA(dstRect)
	// Use CatmullRom for high quality resizing
	draw.CatmullRom.Scale(resizedImg, dstRect, src, bounds, draw.Over, nil)

	return resizedImg
}
