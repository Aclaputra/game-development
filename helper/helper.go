package helper

import (
	"image"
	"os"

	"github.com/Aclaputra/game_development/model"

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

func LoadAndCropImage(req *model.LoadAndCropImageRequest) (*ebiten.Image, error) {
	// Open the file at the given path
	file, err := os.Open(req.Path)
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
	rect := image.Rect(req.X, req.Y, req.X+req.Width, req.Y+req.Height)

	// Crop the image by sub-imaging it
	croppedImg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(rect)

	// Create an Ebiten image from the cropped image
	ebitenImage := ebiten.NewImageFromImage(croppedImg)
	return ebitenImage, nil
}
