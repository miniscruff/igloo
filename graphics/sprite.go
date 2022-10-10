package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo"
	"github.com/miniscruff/igloo/content"
	"github.com/miniscruff/igloo/mathf"
)

type SpriteVisual struct {
	*igloo.Visualer

	sprite  *content.Sprite
	isDirty bool
}

func NewSpriteVisual(sprite *content.Sprite) *SpriteVisual {
	v := &SpriteVisual{
		sprite:  sprite,
		isDirty: true,
	}

	v.Visualer = &igloo.Visualer{
		Transform:   mathf.NewTransform(),
		Children:    make([]*igloo.Visualer, 0),
		Dirtier:     v,
		Drawer:      v,
		NativeSizer: v,
	}

	return v
}

func (v *SpriteVisual) Sprite() *content.Sprite {
	return v.sprite
}

func (v *SpriteVisual) SetSprite(sprite *content.Sprite) {
	if v.sprite == sprite {
		return
	}

	v.sprite = sprite
	v.isDirty = true
}

func (v *SpriteVisual) IsDirty() bool {
	return v.isDirty
}

func (v *SpriteVisual) Clean() {
	v.isDirty = false
}

func (v *SpriteVisual) NativeSize() (float64, float64) {
	pt := v.sprite.Image.Bounds().Size()
	return float64(pt.X), float64(pt.Y)
}

func (v *SpriteVisual) Draw(dest *ebiten.Image) {
	dest.DrawImage(v.sprite.Image, &ebiten.DrawImageOptions{
		GeoM:          v.Transform.GeoM(),
		ColorM:        v.sprite.ColorM,
		Filter:        v.sprite.Filter,
		CompositeMode: v.sprite.CompositeMode,
	})
}

/*

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
	// DrawOptions when drawing, note that the GeoM value is controlled
	// by the transform so changing it here will get overriden.
	DrawOptions *ebiten.DrawImageOptions

	lastVisible bool
	inView      bool
}

// Draw will render the sprite onto the canvas.
// If our transform, sprite or camera are dirty then we will update internal
// values accordingly.
func (s *Sprite) Draw(dest *ebiten.Image, camera Camera) {
	turnedOn := s.Visible && !s.lastVisible
	s.lastVisible = s.Visible

	if !s.Visible {
		return
	}

	if turnedOn || s.Transform.IsDirty() || camera.IsDirty() {
		s.inView = camera.IsInView(s.Transform.Bounds())
		if s.inView {
			screenGeom := camera.WorldToScreen(s.Transform.GeoM())
			s.DrawOptions.GeoM = screenGeom
		}
	}

	if s.inView {
		dest.DrawImage(s.Image, s.DrawOptions)
	}
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
		DrawOptions: &ebiten.DrawImageOptions{},
	}

	return sprite
}

func NewSpriteWithDrawOptions(
	image *ebiten.Image,
	drawOptions ebiten.DrawImageOptions,
	options ...mathf.TransformOption,
) *Sprite {
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
		DrawOptions: &drawOptions,
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
		DrawOptions: &ebiten.DrawImageOptions{},
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
*/
