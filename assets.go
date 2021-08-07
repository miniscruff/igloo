package igloo

import (
	"fmt"
	"image"
	"io/fs"

	// import png for image loading
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func ReadFSFile(fsys fs.FS, path string) ([]byte, error) {
	fileBytes, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("reading font file %v: %w", path, err)
	}

	if len(fileBytes) == 0 {
		return nil, fmt.Errorf("font file empty %v", path)
	}

	return fileBytes, nil
}

func LoadOpenType(fs fs.FS, path string, options *opentype.FaceOptions) (font.Face, error) {
	fontBytes, err := ReadFSFile(fs, path)
	if err != nil {
		return nil, err
	}

	openType, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing opentype font %v: %w", path, err)
	}

	face, err := opentype.NewFace(openType, options)
	if err != nil {
		return nil, fmt.Errorf("loading font face for %v: %w", path, err)
	}

	return face, nil
}

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
