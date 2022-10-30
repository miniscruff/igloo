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

func NewSpriteVisual() *SpriteVisual {
	v := &SpriteVisual{
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
