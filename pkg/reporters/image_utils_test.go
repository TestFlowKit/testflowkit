package reporters

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"strings"
	"testing"
)

func createTestImage(width, height int) []byte {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	// Fill with some color
	for y := range height {
		for x := range width {
			img.Set(x, y, color.RGBA{R: 100, G: 150, B: 200, A: 255})
		}
	}
	var buf bytes.Buffer
	_ = png.Encode(&buf, img)
	return buf.Bytes()
}

func TestOptimizeAndEncodeScreenshot(t *testing.T) {
	t.Run("Valid PNG image", func(t *testing.T) {
		// Create a small image
		rawData := createTestImage(100, 100)

		encoded, err := OptimizeAndEncodeScreenshot(rawData)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !strings.HasPrefix(encoded, "data:image/jpeg;base64,") {
			t.Errorf("Expected prefix 'data:image/jpeg;base64,', got %s", encoded)
		}
	})

	t.Run("Invalid image data", func(t *testing.T) {
		rawData := []byte("invalid data")
		_, err := OptimizeAndEncodeScreenshot(rawData)
		if err == nil {
			t.Error("Expected error for invalid image data, got nil")
		}
	})

	t.Run("Large image resizing", func(t *testing.T) {
		// Create an image larger than maxWidth (1280)
		width := 1300
		height := 1000
		rawData := createTestImage(width, height)

		encodedURI, err := OptimizeAndEncodeScreenshot(rawData)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Decode the result to check dimensions
		parts := strings.Split(encodedURI, ",")
		if len(parts) != 2 {
			t.Fatalf("Invalid data URI format")
		}

		decodedBytes, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			t.Fatalf("Failed to decode base64: %v", err)
		}

		img, _, err := image.Decode(bytes.NewReader(decodedBytes))
		if err != nil {
			t.Fatalf("Failed to decode result image: %v", err)
		}

		if img.Bounds().Dx() > 1280 {
			t.Errorf("Expected width <= 1280, got %d", img.Bounds().Dx())
		}
	})
}

func TestApplyMaxWidthConstraint(t *testing.T) {
	t.Run("No resize needed", func(t *testing.T) {
		img := image.NewRGBA(image.Rect(0, 0, 100, 100))
		result := applyMaxWidthConstraint(img)
		if result.Bounds().Dx() != 100 {
			t.Errorf("Expected width 100, got %d", result.Bounds().Dx())
		}
	})

	t.Run("Resize needed", func(t *testing.T) {
		img := image.NewRGBA(image.Rect(0, 0, 2000, 1000))
		result := applyMaxWidthConstraint(img)
		if result.Bounds().Dx() != 1280 {
			t.Errorf("Expected width 1280, got %d", result.Bounds().Dx())
		}
		// Check aspect ratio preservation
		// 2000 -> 1280 (ratio 0.64)
		// 1000 * 0.64 = 640
		if result.Bounds().Dy() != 640 {
			t.Errorf("Expected height 640, got %d", result.Bounds().Dy())
		}
	})
}
