package global

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

const SCREEN_WIDTH = 800
const SCREEN_HEIGHT = 600

//go:embed img/sprite.png
//go:embed img/bricks.png
//go:embed img/bkg.png

var assets embed.FS

func LoadImg(path string) *ebiten.Image {
	f, err := assets.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)

}

// var bkg = LoadImg("img/bkg.png")
// var bricks = LoadImg("imt/bricks.png")
