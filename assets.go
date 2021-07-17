package igloo

import (
	"fmt"
	"io/fs"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func ReadFontFile(fsys fs.FS, path string) ([]byte, error) {
	fontBytes, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("reading font file %v: %w", path, err)
	}
	if len(fontBytes) <= 0 {
		return nil, fmt.Errorf("font file empty %v", path)
	}
	return fontBytes, nil
}

func LoadOpenType(fs fs.FS, path string, options *opentype.FaceOptions) (font.Face, error) {
	fontBytes, err := ReadFontFile(fs, path)
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
