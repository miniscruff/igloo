package igloo

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite represents an renderable element in the world.
type Sprite struct {
	Image     *ebiten.Image
	Transform *Transform
	// dirty flagging values
	anchor Vec2
	width  float64
	height float64
	// draw cache
	isDirty bool
	inView  bool
	options *ebiten.DrawImageOptions
	geom    ebiten.GeoM
}

func (s *Sprite) IsDirty() bool {
	return s.isDirty
}

func (s *Sprite) Clean() {
	s.isDirty = false
}

func (s *Sprite) Anchor() Vec2 {
	return s.anchor
}

func (s *Sprite) SetAnchor(anchor Vec2) {
	s.anchor = anchor
	s.isDirty = true
}

func (s *Sprite) Width() float64 {
	return s.width
}

func (s *Sprite) SetWidth(width float64) {
	s.width = width
	s.isDirty = true
}

func (s *Sprite) Height() float64 {
	return s.height
}

func (s *Sprite) SetHeight(height float64) {
	s.height = height
	s.isDirty = true
}

func (s *Sprite) createGeoM() ebiten.GeoM {
	geom := ebiten.GeoM{}

	intWidth, intHeight := s.Image.Size()
	imageWidth := float64(intWidth)
	imageHeight := float64(intHeight)
	if imageWidth != s.width || imageHeight != s.height {
		geom.Scale(s.width/imageWidth, s.height/imageHeight)
	}
	if s.anchor != Vec2Zero {
		geom.Translate(
			-s.width*s.anchor.X,
			-s.height*s.anchor.Y,
		)
	}
	if s.Transform.Rotation() != 0 {
		geom.Rotate(s.Transform.Rotation())
	}

	geom.Translate(s.Transform.X(), s.Transform.Y())
	return geom
}

func (s *Sprite) Draw(canvas Canvaser, camera Camera) {
	transformDirty := s.Transform.IsDirty()
	if transformDirty || s.IsDirty() {
		s.geom = s.createGeoM()
	}

	if transformDirty || s.IsDirty() || camera.IsDirty() {
		s.inView = camera.IsInView(s.Transform.X(), s.Transform.Y(), s.width, s.height)
		if s.inView {
			screenGeom := camera.WorldToScreen(s.geom)
			s.options = &ebiten.DrawImageOptions{
				GeoM: screenGeom,
			}
		}
	}

	if s.inView {
		canvas.DrawImage(s.Image, s.options)
	}
	s.Clean()
}

func NewSprite(image *ebiten.Image, transform *Transform) *Sprite {
	w, h := image.Size()
	return NewSpriteAnchorSize(image, transform, Vec2Zero, float64(w), float64(h))
}

func NewSpriteAnchor(image *ebiten.Image, transform *Transform, anchor Vec2) *Sprite {
	w, h := image.Size()
	return NewSpriteAnchorSize(image, transform, anchor, float64(w), float64(h))
}

func NewSpriteSize(image *ebiten.Image, transform *Transform, width, height float64) *Sprite {
	return NewSpriteAnchorSize(image, transform, Vec2Zero, width, height)
}

func NewSpriteAnchorSize(image *ebiten.Image, transform *Transform, anchor Vec2, width, height float64) *Sprite {
	return &Sprite{
		Image:     image,
		Transform: transform,
		anchor:    anchor,
		width:     width,
		height:    height,
		isDirty:   true, // start dirty
		geom:      ebiten.GeoM{},
	}
}
