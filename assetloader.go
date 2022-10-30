package igloo

import (
	"fmt"
	"image"
	"io/fs"

	// import png for image loading
	_ "image/png"

	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
)

// ContentManager handles content loading, unloading, caching
// and sharing.
type AssetLoader struct {
	fsys    fs.FS
	rootDir string
}

func NewAssetLoader(fsys fs.FS, rootDir string) *AssetLoader {
	return &AssetLoader{
		fsys:    fsys,
		rootDir: rootDir,
	}
}

func (a *AssetLoader) fullPath(path string) string {
	return a.rootDir + "/" + path
}

func (a *AssetLoader) readFSFile(path string) ([]byte, error) {
	fullPath := a.fullPath(path)
	fileBytes, err := fs.ReadFile(a.fsys, fullPath)

	if err != nil {
		return nil, fmt.Errorf("reading file %v: %w", fullPath, err)
	}

	if len(fileBytes) == 0 {
		return nil, fmt.Errorf("file empty %v", fullPath)
	}

	return fileBytes, nil
}

func (a *AssetLoader) LoadImage(path string) (*ebiten.Image, error) {
	fullPath := a.fullPath(path)

	file, err := a.fsys.Open(fullPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	ebiImage := ebiten.NewImageFromImage(img)

	return ebiImage, nil
}

func (a *AssetLoader) LoadOpenType(path string) (*opentype.Font, error) {
	fontBytes, err := a.readFSFile(path)
	if err != nil {
		return nil, err
	}

	openType, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing opentype font %v: %w", path, err)
	}

	return openType, nil
}
