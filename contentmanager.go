package igloo

import (
	"fmt"
	"image"
	"io/fs"

	// import png for image loading
	_ "image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"

	"github.com/hajimehoshi/ebiten/v2"
)

// ContentManager handles content loading, unloading, caching
// and sharing.
type ContentManager struct {
	fsys    fs.FS
	rootDir string

	images map[string]*ebiten.Image
	fonts  map[string]*opentype.Font
	faces  []font.Face
}

func NewContentManager(fsys fs.FS, rootDir string) *ContentManager {
	return &ContentManager{
		fsys:    fsys,
		rootDir: rootDir,
		images:  make(map[string]*ebiten.Image),
		fonts:   make(map[string]*sfnt.Font),
		faces:   make([]font.Face, 0),
	}
}

func (cm *ContentManager) fullPath(path string) string {
	return cm.rootDir + "/" + path
}

func (cm *ContentManager) readFSFile(path string) ([]byte, error) {
	fullPath := cm.fullPath(path)
	fileBytes, err := fs.ReadFile(cm.fsys, fullPath)

	if err != nil {
		return nil, fmt.Errorf("reading file %v: %w", fullPath, err)
	}

	if len(fileBytes) == 0 {
		return nil, fmt.Errorf("file empty %v", fullPath)
	}

	return fileBytes, nil
}

func (cm *ContentManager) LoadImage(path string) (*ebiten.Image, error) {
	cached, found := cm.images[path]
	if found {
		return cached, nil
	}

	fullPath := cm.fullPath(path)

	file, err := cm.fsys.Open(fullPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	ebiImage := ebiten.NewImageFromImage(img)
	cm.images[path] = ebiImage

	return ebiImage, nil
}

func (cm *ContentManager) LoadOpenType(path string) (*sfnt.Font, error) {
	cached, found := cm.fonts[path]
	if found {
		return cached, nil
	}

	fontBytes, err := cm.readFSFile(path)
	if err != nil {
		return nil, err
	}

	openType, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing opentype font %v: %w", path, err)
	}

	cm.fonts[path] = openType

	return openType, nil
}

// LoadFontFace will load a font face from font and options.
// Note that no caching is done for this, so only load once  per unique font and options.
func (cm *ContentManager) LoadFontFace(
	font *sfnt.Font,
	options *opentype.FaceOptions,
) (font.Face, error) {
	face, err := opentype.NewFace(font, options)
	if err != nil {
		return nil, fmt.Errorf("loading font face for %v: %w", font, err)
	}

	cm.faces = append(cm.faces, face)

	return face, nil
}

func (cm *ContentManager) Dispose() {
	for _, i := range cm.images {
		i.Dispose()
	}

	for _, f := range cm.faces {
		f.Close()
	}

	cm.images = make(map[string]*ebiten.Image)
	cm.fonts = make(map[string]*sfnt.Font)
	cm.faces = make([]font.Face, 0)
}
