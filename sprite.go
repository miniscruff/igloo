package igloo

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteConfig determines the starting state of our sprite
type SpriteConfig struct {
	Image     *ebiten.Image
	Transform *Transform
	// Value of 0 will use the image width
	Width float64
	// Value of 0 will use the image height
	Height float64
	// Defaults to top left
	Anchor Vec2f
}

// Sprite represents a renderable element in the world.
type Sprite struct {
	Image     *ebiten.Image
	Transform *Transform
	// dirty flagging values
	anchor Vec2f
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
func (s *Sprite) Anchor() Vec2f {
	return s.anchor
}

// SetAnchor will change our rotation point.
// Will also mark the sprite as dirty.
// (0, 0) will rotate around the top left
// (0.5, 0.5) will rotate around the center
// (1, 1) will rotate around the bottom right
func (s *Sprite) SetAnchor(anchor Vec2f) {
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

func (s *Sprite) Size() (float64, float64) {
	return s.width, s.height
}

func (s *Sprite) createGeoM() ebiten.GeoM {
	geom := ebiten.GeoM{}

	intWidth, intHeight := s.Image.Size()
	imageWidth := float64(intWidth)
	imageHeight := float64(intHeight)

	if imageWidth != s.width || imageHeight != s.height {
		geom.Scale(s.width/imageWidth, s.height/imageHeight)
	}

	if s.anchor != Vec2fZero {
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
func NewSprite(config SpriteConfig) *Sprite {
	w, h := config.Image.Size()

	if config.Width <= 0 {
		config.Width = float64(w)
	}

	if config.Height <= 0 {
		config.Height = float64(h)
	}

	return &Sprite{
		Image:     config.Image,
		Transform: config.Transform,
		anchor:    config.Anchor,
		width:     config.Width,
		height:    config.Height,
		isDirty:   true, // start dirty
		inView:    false,
		geom:      ebiten.GeoM{},
	}
}
