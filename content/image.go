package content

import (
	"image"
	"io/fs"

	// import png for image loading
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImage(fs fs.FS, path string) (*ebiten.Image, error) {
	file, err := fs.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
