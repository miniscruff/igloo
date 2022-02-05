package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo/mathf"
	"image"
)

// NineSlice represents a renderable element in the world.
type NineSlice struct {
	Image     *ebiten.Image
	Transform *mathf.Transform
	// dirty flagging values
	anchor mathf.Vec2
	width  float64
	height float64
	// draw cache
	isDirty bool
	inView  bool

	// internal rendering values
	topLeft      *Sprite
	topCenter    *Sprite
	topRight     *Sprite
	middleLeft   *Sprite
	middleCenter *Sprite
	middleRight  *Sprite
	bottomLeft   *Sprite
	bottomCenter *Sprite
	bottomRight  *Sprite
	borders      SliceBorders
}

// IsDirty returns whether or not our internal state has changed since last
// drawing. When dirty our next drawing attempt will refresh drawing values.
func (s *NineSlice) IsDirty() bool {
	return s.isDirty
}

// Clean resets our dirty state, automatically called when drawing
func (s *NineSlice) Clean() {
	s.isDirty = false
}

// Anchor determines our rotation point.
func (s *NineSlice) Anchor() mathf.Vec2 {
	return s.anchor
}

// SetAnchor will change our rotation point.
// Will also mark the sprite as dirty.
// (0, 0) will rotate around the top left
// (0.5, 0.5) will rotate around the center
// (1, 1) will rotate around the bottom right
func (s *NineSlice) SetAnchor(anchor mathf.Vec2) {
	if s.anchor == anchor {
		return
	}

	s.anchor = anchor
	s.isDirty = true
}

// Width returns our drawing width
func (s *NineSlice) Width() float64 {
	return s.width
}

// SetWidth will change our drawing width.
// Will also mark the sprite as dirty.
func (s *NineSlice) SetWidth(width float64) {
	if s.width == width {
		return
	}

	s.width = width
	s.isDirty = true
}

// Height returns our drawing height
func (s *NineSlice) Height() float64 {
	return s.height
}

// SetHeight will change our drawing height.
// Will also mark the sprite as dirty.
func (s *NineSlice) SetHeight(height float64) {
	if s.height == height {
		return
	}

	s.height = height
	s.isDirty = true
}

func (s *NineSlice) Size() (float64, float64) {
	return s.width, s.height
}

func (s *NineSlice) positionAndScaleImages() {
	centerWidth := s.width - float64(s.borders.Left-s.borders.Right)
	middleHeight := s.height - float64(s.borders.Top-s.borders.Bottom)
	halfCenterWidth := centerWidth / 2
	halfMiddleHeight := middleHeight / 2
	centerX := s.Transform.X() - s.width*s.anchor.X + s.width*0.5
	middleY := s.Transform.Y() - s.height*s.anchor.Y + s.height*0.5

	// center
	s.middleCenter.Transform.SetX(centerX)
	s.middleCenter.Transform.SetY(middleY)
	s.middleCenter.SetWidth(centerWidth)
	s.middleCenter.SetHeight(middleHeight)

	// corners
	s.topLeft.Transform.SetX(centerX - halfCenterWidth)
	s.topLeft.Transform.SetY(middleY - halfMiddleHeight)

	s.topRight.Transform.SetX(centerX + halfCenterWidth)
	s.topRight.Transform.SetY(middleY - halfMiddleHeight)

	s.bottomLeft.Transform.SetX(centerX - halfCenterWidth)
	s.bottomLeft.Transform.SetY(middleY + halfMiddleHeight)

	s.bottomRight.Transform.SetX(centerX + halfCenterWidth)
	s.bottomRight.Transform.SetY(middleY + halfMiddleHeight)

	// edges
	s.topCenter.Transform.SetX(centerX)
	s.topCenter.Transform.SetY(middleY - halfMiddleHeight)
	s.topCenter.SetWidth(centerWidth)

	s.bottomCenter.Transform.SetX(centerX)
	s.bottomCenter.Transform.SetY(middleY + halfMiddleHeight)
	s.bottomCenter.SetWidth(centerWidth)

	s.middleLeft.Transform.SetX(centerX - halfCenterWidth)
	s.middleLeft.Transform.SetY(middleY)
	s.middleLeft.SetHeight(middleHeight)

	s.middleRight.Transform.SetX(centerX + halfCenterWidth)
	s.middleRight.Transform.SetY(middleY)
	s.middleRight.SetHeight(middleHeight)
}

// Draw will render the sprite onto the canvas.
// If our transform, sprite or camera are dirty then we will update internal
// values accordingly.
func (s *NineSlice) Draw(dest *ebiten.Image, camera Camera) {
	transformDirty := s.Transform.IsDirty()
	if transformDirty || s.IsDirty() {
		s.positionAndScaleImages()
	}

	if transformDirty || s.IsDirty() || camera.IsDirty() {
		anchorOffset := s.anchor.Mul(mathf.Vec2{X: s.width, Y: s.height})
		topLeft := s.Transform.Position().Sub(anchorOffset)
		s.inView = camera.IsInView(topLeft, s.width, s.height)
	}

	if s.inView {
		s.topLeft.Draw(dest, camera)
		s.topCenter.Draw(dest, camera)
		s.topRight.Draw(dest, camera)
		s.middleLeft.Draw(dest, camera)
		s.middleCenter.Draw(dest, camera)
		s.middleRight.Draw(dest, camera)
		s.bottomLeft.Draw(dest, camera)
		s.bottomCenter.Draw(dest, camera)
		s.bottomRight.Draw(dest, camera)
	}

	s.Clean()
}

// NineSliceOption allows setting the starting options for a sprite.
// You can also just set the values yourself after using the proper
// set methods.
type NineSliceOption func(s *NineSlice)

func NineSliceAtPosition(position mathf.Vec2) NineSliceOption {
	return func(s *NineSlice) {
		s.Transform.SetPosition(position)
	}
}

func NineSliceAtXY(x, y float64) NineSliceOption {
	return func(s *NineSlice) {
		s.Transform.SetPosition(mathf.Vec2{X: x, Y: y})
	}
}

func NineSliceWithRotation(rotation float64) NineSliceOption {
	return func(s *NineSlice) {
		s.Transform.SetRotation(rotation)
	}
}

func NineSliceWithWidth(width float64) NineSliceOption {
	return func(s *NineSlice) {
		s.width = width
	}
}

func NineSliceWithHeight(height float64) NineSliceOption {
	return func(s *NineSlice) {
		s.height = height
	}
}

func NineSliceWithSize(width, height float64) NineSliceOption {
	return func(s *NineSlice) {
		s.width = width
		s.height = height
	}
}

func NineSliceScaled(factor float64) NineSliceOption {
	return func(s *NineSlice) {
		s.width *= factor
		s.height *= factor
	}
}

func NineSliceWithAnchor(anchor mathf.Vec2) NineSliceOption {
	return func(s *NineSlice) {
		s.anchor = anchor
	}
}

// SliceBorders organizes the sizes for our nine slice to generate the internal sprites.
type SliceBorders struct {
	// Width of the left column
	Left int
	// Width of the center column
	Center int
	// Width of the right column
	Right int
	// Height of the top row
	Top int
	// Height of the middle row
	Middle int
	// Height of the bottom row
	Bottom int
}

func SubImageWidthHeight(img *ebiten.Image, x, y, width, height int) *ebiten.Image {
	sx := img.Bounds().Min.X
	sy := img.Bounds().Min.Y
	rect := image.Rect(sx+x, sy+y, sx+x+width, sy+y+height)
	return img.SubImage(rect).(*ebiten.Image)
}

func (b *SliceBorders) TopLeftSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, 0, 0, b.Left, b.Top),
		SpriteWithAnchor(mathf.AnchorBottomRight),
	)
}

func (b *SliceBorders) TopCenterSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left, 0, b.Center, b.Top),
		SpriteWithAnchor(mathf.AnchorBottomCenter),
	)
}

func (b *SliceBorders) TopRightSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left+b.Center, 0, b.Right, b.Top),
		SpriteWithAnchor(mathf.AnchorBottomLeft),
	)
}

func (b *SliceBorders) MiddleLeftSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, 0, b.Top, b.Left, b.Middle),
		SpriteWithAnchor(mathf.AnchorMiddleRight),
	)
}

func (b *SliceBorders) MiddleCenterSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left, b.Top, b.Center, b.Middle),
		SpriteWithAnchor(mathf.AnchorMiddleCenter),
	)
}

func (b *SliceBorders) MiddleRightSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left+b.Center, b.Top, b.Right, b.Middle),
		SpriteWithAnchor(mathf.AnchorMiddleLeft),
	)
}

func (b *SliceBorders) BottomLeftSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, 0, b.Top+b.Middle, b.Left, b.Bottom),
		SpriteWithAnchor(mathf.AnchorTopRight),
	)
}

func (b *SliceBorders) BottomCenterSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left, b.Top+b.Middle, b.Center, b.Bottom),
		SpriteWithAnchor(mathf.AnchorTopCenter),
	)
}

func (b *SliceBorders) BottomRightSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left+b.Center, b.Top+b.Middle, b.Right, b.Bottom),
		SpriteWithAnchor(mathf.AnchorTopLeft),
	)
}

// NewNineSlice will create a nine slice from an image and border values
// Defaults include:
// * Width and height of the image
// * Positon at 0,0
// * Rotation of 0
// * Anchor in the middle center
func NewNineSlice(img *ebiten.Image, borders SliceBorders, options ...NineSliceOption) *NineSlice {
	w, h := img.Size()

	nineSlice := &NineSlice{
		Image:        img,
		Transform:    mathf.NewTransform(),
		width:        float64(w),
		height:       float64(h),
		anchor:       mathf.AnchorMiddleCenter,
		isDirty:      true,
		inView:       false,
		borders:      borders,
		topLeft:      borders.TopLeftSprite(img),
		topCenter:    borders.TopCenterSprite(img),
		topRight:     borders.TopRightSprite(img),
		middleLeft:   borders.MiddleLeftSprite(img),
		middleCenter: borders.MiddleCenterSprite(img),
		middleRight:  borders.MiddleRightSprite(img),
		bottomLeft:   borders.BottomLeftSprite(img),
		bottomCenter: borders.BottomCenterSprite(img),
		bottomRight:  borders.BottomRightSprite(img),
	}

	for _, o := range options {
		o(nineSlice)
	}

	return nineSlice
}
