package graphics

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/miniscruff/igloo/mathf"
)

// Sprite represents a renderable element in the world.
type Sprite struct {
	Image     *ebiten.Image
	Transform *mathf.Transform
	Visible   bool

	lastVisible bool
	inView      bool
	options     *ebiten.DrawImageOptions
}

// Draw will render the sprite onto the canvas.
// If our transform, sprite or camera are dirty then we will update internal
// values accordingly.
func (s *Sprite) Draw(dest *ebiten.Image, camera Camera) {
	turnedOn := s.Visible && !s.lastVisible

	if turnedOn || (s.Visible && (s.Transform.IsDirty() || camera.IsDirty())) {
		s.inView = camera.IsInView(s.Transform.Bounds())
		if s.inView {
			screenGeom := camera.WorldToScreen(s.Transform.GeoM())
			s.options.GeoM = screenGeom
		}
	}

	if s.inView && s.Visible {
		dest.DrawImage(s.Image, s.options)
	}

	s.lastVisible = s.Visible
}

// NewSprite will create a sprite with image and transform options.
// Will configure transform to use the size of our image as the natural size.
func NewSprite(image *ebiten.Image, options ...mathf.TransformOption) *Sprite {
	w, h := image.Size()
	// prepend our natural size option
	options = append([]mathf.TransformOption{
		mathf.TransformWithNaturalSize(float64(w), float64(h)),
	}, options...)
	transform := mathf.NewTransform(options...)

	sprite := &Sprite{
		Image:       image,
		Transform:   transform,
		Visible:     true,
		lastVisible: true,
		inView:      false,
		options:     &ebiten.DrawImageOptions{},
	}

	return sprite
}

// NewSprite will create a sprite with image and transform.
// Will configure transform to use the size of our image as the natural size.
func NewSpriteWithTransform(image *ebiten.Image, transform *mathf.Transform) *Sprite {
	w, h := image.Size()
	transform.SetNaturalWidth(float64(w))
	transform.SetNaturalHeight(float64(h))

	sprite := &Sprite{
		Image:       image,
		Transform:   transform,
		Visible:     true,
		lastVisible: true,
		inView:      false,
		options:     &ebiten.DrawImageOptions{},
	}

	return sprite
}

// SpriteSheet represents a collection of images
type SpriteSheet []*ebiten.Image

func (ss *SpriteSheet) FrameAt(percent float64) *ebiten.Image {
	count := float64(len(*ss))
	return (*ss)[int(mathf.Lerp(0, count-1, percent))]
}

func SheetFromGrid(sheet *ebiten.Image, columns, rows, frames int) *SpriteSheet {
	var images SpriteSheet = make([]*ebiten.Image, frames)

	w, h := sheet.Size()
	fw := w / columns
	fh := h / rows

	for y := 0; y < rows; y++ {
		for x := 0; x < columns; x++ {
			i := y*rows + x
			images[i] = sheet.SubImage(
				image.Rect(x*fw, y*fw, (x+1)*fw, (y+1)*fh),
			).(*ebiten.Image)

			if i == frames-1 {
				goto done
			}
		}
	}

done:
	return &images
}

func NewFrameClip(
	target *Sprite,
	sheet *SpriteSheet,
	duration float64,
	options ...mathf.TweenOption,
) *mathf.Tween {
	t := mathf.NewTween(
		duration,
		mathf.TweenUpdateFunc(func(value float64) {
			target.Image = sheet.FrameAt(value)
		}),
	)

	for _, opt := range options {
		opt(t)
	}

	return t
}
