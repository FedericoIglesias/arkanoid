package img

import (
	"image"
	"image/png"
	_ "image/png"
	"os"
)

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func GetSprite(pathPNG string, sizeX, sizeY, posX, posY int) image.Image {
	file, err := os.Open(pathPNG)

	if err != nil {
		panic("error to read file")
	}

	originalImage, err := png.Decode(file)
	if err != nil {
		panic("error to read file")
	}
	cropSize := image.Rect(0, 0, sizeX, sizeY)
	cropSize = cropSize.Add(image.Point{posX, posY})
	croppedImage := originalImage.(SubImager).SubImage(cropSize)

	// croppedImageFile, err := os.Create("cropped.png")
	// if err != nil {
	// 	panic(err)
	// }
	// defer croppedImageFile.Close()
	// if err := png.Encode(croppedImageFile, croppedImage); err != nil {
	// 	panic(err)
	// }
	return croppedImage
}
