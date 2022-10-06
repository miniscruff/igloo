package graphics
/*

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/miniscruff/igloo/mathf"
)

// NineSlice represents a renderable slicable element in the world.
type NineSlice struct {
	Image     *ebiten.Image
	Transform *mathf.Transform
	Visible   bool

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
	lastVisible  bool
}

func (s *NineSlice) positionAndScaleImages() {
	width := s.Transform.Width()
	height := s.Transform.Height()
	anchor := s.Transform.Anchor()

	centerWidth := width - float64(s.borders.Left+s.borders.Right)
	if centerWidth <= 0 {
		centerWidth = 0
	}

	middleHeight := height - float64(s.borders.Top+s.borders.Bottom)
	if middleHeight <= 0 {
		middleHeight = 0
	}

	halfCenterWidth := centerWidth / 2
	halfMiddleHeight := middleHeight / 2
	centerX := s.Transform.X() - width*anchor.X + width*0.5
	middleY := s.Transform.Y() - height*anchor.Y + height*0.5

	// center
	s.middleCenter.Transform.SetX(centerX)
	s.middleCenter.Transform.SetY(middleY)
	s.middleCenter.Transform.SetWidth(centerWidth)
	s.middleCenter.Transform.SetHeight(middleHeight)

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
	s.topCenter.Transform.SetWidth(centerWidth)

	s.bottomCenter.Transform.SetX(centerX)
	s.bottomCenter.Transform.SetY(middleY + halfMiddleHeight)
	s.bottomCenter.Transform.SetWidth(centerWidth)

	s.middleLeft.Transform.SetX(centerX - halfCenterWidth)
	s.middleLeft.Transform.SetY(middleY)
	s.middleLeft.Transform.SetHeight(middleHeight)

	s.middleRight.Transform.SetX(centerX + halfCenterWidth)
	s.middleRight.Transform.SetY(middleY)
	s.middleRight.Transform.SetHeight(middleHeight)
}

// Draw will render the sprite onto the canvas.
// If our transform, sprite or camera are dirty then we will update internal
// values accordingly.
func (s *NineSlice) Draw(dest *ebiten.Image, camera Camera) {
	turnedOn := s.Visible && !s.lastVisible
	turnedOff := !s.Visible && s.lastVisible

	if s.Transform.IsDirty() || turnedOn {
		s.positionAndScaleImages()
		s.Transform.Clean()
	}

	if turnedOn || turnedOff {
		s.topLeft.Visible = s.Visible
		s.topCenter.Visible = s.Visible
		s.topRight.Visible = s.Visible
		s.middleLeft.Visible = s.Visible
		s.middleCenter.Visible = s.Visible
		s.middleRight.Visible = s.Visible
		s.bottomLeft.Visible = s.Visible
		s.bottomCenter.Visible = s.Visible
		s.bottomRight.Visible = s.Visible
	}

	s.topLeft.Draw(dest, camera)
	s.topCenter.Draw(dest, camera)
	s.topRight.Draw(dest, camera)
	s.middleLeft.Draw(dest, camera)
	s.middleCenter.Draw(dest, camera)
	s.middleRight.Draw(dest, camera)
	s.bottomLeft.Draw(dest, camera)
	s.bottomCenter.Draw(dest, camera)
	s.bottomRight.Draw(dest, camera)

	s.lastVisible = s.Visible
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
		mathf.TransformWithAnchor(mathf.AnchorBottomRight),
	)
}

func (b *SliceBorders) TopCenterSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left, 0, b.Center, b.Top),
		mathf.TransformWithAnchor(mathf.AnchorBottomCenter),
	)
}

func (b *SliceBorders) TopRightSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left+b.Center, 0, b.Right, b.Top),
		mathf.TransformWithAnchor(mathf.AnchorBottomLeft),
	)
}

func (b *SliceBorders) MiddleLeftSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, 0, b.Top, b.Left, b.Middle),
		mathf.TransformWithAnchor(mathf.AnchorMiddleRight),
	)
}

func (b *SliceBorders) MiddleCenterSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left, b.Top, b.Center, b.Middle),
		mathf.TransformWithAnchor(mathf.AnchorMiddleCenter),
	)
}

func (b *SliceBorders) MiddleRightSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left+b.Center, b.Top, b.Right, b.Middle),
		mathf.TransformWithAnchor(mathf.AnchorMiddleLeft),
	)
}

func (b *SliceBorders) BottomLeftSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, 0, b.Top+b.Middle, b.Left, b.Bottom),
		mathf.TransformWithAnchor(mathf.AnchorTopRight),
	)
}

func (b *SliceBorders) BottomCenterSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left, b.Top+b.Middle, b.Center, b.Bottom),
		mathf.TransformWithAnchor(mathf.AnchorTopCenter),
	)
}

func (b *SliceBorders) BottomRightSprite(img *ebiten.Image) *Sprite {
	return NewSprite(
		SubImageWidthHeight(img, b.Left+b.Center, b.Top+b.Middle, b.Right, b.Bottom),
		mathf.TransformWithAnchor(mathf.AnchorTopLeft),
	)
}

// NewNineSlice will create a nine slice from an image and border values
// Defaults include:
// * Width and height of the image
// * Position at 0,0
// * Rotation of 0
// * Anchor in the middle center
func NewNineSlice(
	image *ebiten.Image,
	borders SliceBorders,
	options ...mathf.TransformOption,
) *NineSlice {
	w, h := image.Size()
	// prepend our natural size option
	options = append([]mathf.TransformOption{
		mathf.TransformWithNaturalSize(float64(w), float64(h)),
	}, options...)
	transform := mathf.NewTransform(options...)

	nineSlice := &NineSlice{
		Image:        image,
		Transform:    transform,
		Visible:      true,
		borders:      borders,
		topLeft:      borders.TopLeftSprite(image),
		topCenter:    borders.TopCenterSprite(image),
		topRight:     borders.TopRightSprite(image),
		middleLeft:   borders.MiddleLeftSprite(image),
		middleCenter: borders.MiddleCenterSprite(image),
		middleRight:  borders.MiddleRightSprite(image),
		bottomLeft:   borders.BottomLeftSprite(image),
		bottomCenter: borders.BottomCenterSprite(image),
		bottomRight:  borders.BottomRightSprite(image),
		lastVisible:  true,
	}

	return nineSlice
}
*/
