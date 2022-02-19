package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo/mathf"
)

// Sprite represents a renderable element in the world.
type Sprite struct {
	Image     *ebiten.Image
	Transform *mathf.Transform
	// dirty flagging values
	anchor mathf.Vec2
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
func (s *Sprite) Anchor() mathf.Vec2 {
	return s.anchor
}

// SetAnchor will change our rotation point.
// Will also mark the sprite as dirty.
// (0, 0) will rotate around the top left
// (0.5, 0.5) will rotate around the center
// (1, 1) will rotate around the bottom right
func (s *Sprite) SetAnchor(anchor mathf.Vec2) {
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

func (s *Sprite) Bounds() mathf.Bounds {
	anchorOffset := s.anchor.Mul(mathf.Vec2{X: s.width, Y: s.height})
	topLeft := s.Transform.Position().Sub(anchorOffset)
	return mathf.NewBoundsWidthHeight(
		topLeft.X,
		topLeft.Y,
		s.width,
		s.height,
	)
}

func (s *Sprite) createGeoM() ebiten.GeoM {
	geom := ebiten.GeoM{}

	intWidth, intHeight := s.Image.Size()
	imageWidth := float64(intWidth)
	imageHeight := float64(intHeight)

	if imageWidth != s.width || imageHeight != s.height {
		geom.Scale(s.width/imageWidth, s.height/imageHeight)
	}

	if s.anchor != mathf.Vec2Zero {
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
func (s *Sprite) Draw(dest *ebiten.Image, camera Camera) {
	transformDirty := s.Transform.IsDirty()
	if transformDirty || s.IsDirty() {
		s.geom = s.createGeoM()
	}

	if transformDirty || s.IsDirty() || camera.IsDirty() {
		anchorOffset := s.anchor.Mul(mathf.Vec2{X: s.width, Y: s.height})
		topLeft := s.Transform.Position().Sub(anchorOffset)
		s.inView = camera.IsInView(topLeft, s.width, s.height)

		if s.inView {
			screenGeom := camera.WorldToScreen(s.geom)
			s.options = &ebiten.DrawImageOptions{
				GeoM: screenGeom,
			}
		}
	}

	if s.inView {
		dest.DrawImage(s.Image, s.options)
	}

	s.Clean()
}

// SpriteOption allows setting the starting options for a sprite.
// You can also just set the values yourself after using the proper
// set methods.
type SpriteOption func(s *Sprite)

func SpriteAtPosition(position mathf.Vec2) SpriteOption {
	return func(s *Sprite) {
		s.Transform.SetPosition(position)
	}
}

func SpriteAtXY(x, y float64) SpriteOption {
	return func(s *Sprite) {
		s.Transform.SetPosition(mathf.Vec2{X: x, Y: y})
	}
}

func SpriteWithRotation(rotation float64) SpriteOption {
	return func(s *Sprite) {
		s.Transform.SetRotation(rotation)
	}
}

func SpriteWithWidth(width float64) SpriteOption {
	return func(s *Sprite) {
		s.width = width
	}
}

func SpriteWithHeight(height float64) SpriteOption {
	return func(s *Sprite) {
		s.height = height
	}
}

func SpriteWithSize(width, height float64) SpriteOption {
	return func(s *Sprite) {
		s.width = width
		s.height = height
	}
}

func SpriteScaled(factor float64) SpriteOption {
	return func(s *Sprite) {
		s.width *= factor
		s.height *= factor
	}
}

func SpriteWithAnchor(anchor mathf.Vec2) SpriteOption {
	return func(s *Sprite) {
		s.anchor = anchor
	}
}

// NewSprite will create a sprite from image.
// Defaults include:
// * Width and height of the image
// * Positon at 0,0
// * Rotation of 0
// * Anchor in the middle center
func NewSprite(image *ebiten.Image, options ...SpriteOption) *Sprite {
	w, h := image.Size()
	sprite := &Sprite{
		Image:     image,
		Transform: mathf.NewTransform(),
		width:     float64(w),
		height:    float64(h),
		anchor:    mathf.AnchorMiddleCenter,
		isDirty:   true,
		inView:    false,
		geom:      ebiten.GeoM{},
	}

	for _, o := range options {
		o(sprite)
	}

	return sprite
}
