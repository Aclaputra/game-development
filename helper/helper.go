package helper

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(path string) (*ebiten.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	ebitenImage := ebiten.NewImageFromImage(img)
	return ebitenImage, nil
}

func LoadAndCropImage(path string, x, y, width, height int) (*ebiten.Image, error) {
	// Open the file at the given path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the image from the file
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	// Define the rectangle for cropping
	rect := image.Rect(x, y, x+width, y+height)

	// Crop the image by sub-imaging it
	croppedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(rect)

	// Create an Ebiten image from the cropped image
	ebitenImage := ebiten.NewImageFromImage(croppedImg)
	return ebitenImage, nil
}
