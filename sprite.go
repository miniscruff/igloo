package igloo

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Sprite represents a renderable element in the world.
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

// IsDirty returns whether or not our internal state has changed since last
// drawing. When dirty our next drawing attempt will refresh drawing values.
func (s *Sprite) IsDirty() bool {
	return s.isDirty
}

// Clean resets our dirty state, automatically called when drawing
func (s *Sprite) Clean() {
	s.isDirty = false
}

// Anchor determines our rotation point.
func (s *Sprite) Anchor() Vec2 {
	return s.anchor
}

// SetAnchor will change our rotation point.
// Will also mark the sprite as dirty.
// (0, 0) will rotate around the top left
// (0.5, 0.5) will rotate around the center
// (1, 1) will rotate around the bottom right
func (s *Sprite) SetAnchor(anchor Vec2) {
	if s.anchor == anchor {
		return
	}
	s.anchor = anchor
	s.isDirty = true
}

// Width returns our drawing width
func (s *Sprite) Width() float64 {
	return s.width
}

// SetWidth will change our drawing width.
// Will also mark the sprite as dirty.
func (s *Sprite) SetWidth(width float64) {
	if s.width == width {
		return
	}
	s.width = width
	s.isDirty = true
}

// Height returns our drawing height
func (s *Sprite) Height() float64 {
	return s.height
}

// SetHeight will change our drawing height.
// Will also mark the sprite as dirty.
func (s *Sprite) SetHeight(height float64) {
	if s.height == height {
		return
	}
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

// Draw will render the sprite onto the canvas.
// If our transform, sprite or camera are dirty then we will update internal
// values accordingly.
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

// NewSprite will create a basic sprite from image and transform.
// Anchor defaults to (0,0) and size will default to the image size.
func NewSprite(image *ebiten.Image, transform *Transform) *Sprite {
	w, h := image.Size()
	return NewSpriteAnchorSize(image, transform, Vec2Zero, float64(w), float64(h))
}

// NewSpriteAnchor will create a sprite with a custom anchor.
// Size will default to the image size.
func NewSpriteAnchor(image *ebiten.Image, transform *Transform, anchor Vec2) *Sprite {
	w, h := image.Size()
	return NewSpriteAnchorSize(image, transform, anchor, float64(w), float64(h))
}

// NewSpriteSize will create a sprite with a custom size.
// Anchor defaults to (0,0)
func NewSpriteSize(image *ebiten.Image, transform *Transform, width, height float64) *Sprite {
	return NewSpriteAnchorSize(image, transform, Vec2Zero, width, height)
}

// NewSpriteAnchorSize will create a sprite with a custom anchor and size
func NewSpriteAnchorSize(
	image *ebiten.Image, transform *Transform, anchor Vec2, width, height float64,
) *Sprite {
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
