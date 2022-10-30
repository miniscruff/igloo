package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/miniscruff/igloo"
	"github.com/miniscruff/igloo/content"
	"github.com/miniscruff/igloo/mathf"
)

// probably want a text render mode like scale, word wrap or something
// otherwise the size is just based on the text input

type LabelVisual struct {
	*igloo.Visualer
	ebiten.ColorM

	font    *content.Font
	text    string
	isDirty bool
}

func NewLabelVisual() *LabelVisual {
	v := &LabelVisual{
		isDirty: false,
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

func (v *LabelVisual) IsDirty() bool {
	return v.isDirty
}

func (v *LabelVisual) Clean() {
	v.isDirty = false
}

func (v *LabelVisual) Font() *content.Font {
	return v.font
}

func (v *LabelVisual) SetFont(f *content.Font) {
	if v.font == f {
		return
	}

	v.Transform.SetFixedOffset(f.Ascent())
	v.font = f
	v.isDirty = true
}

func (v *LabelVisual) Text() string {
	return v.text
}

func (v *LabelVisual) SetText(newText string) {
	if v.text == newText {
		return
	}

	v.text = newText
	v.isDirty = true
}

func (v *LabelVisual) NativeSize() (float64, float64) {
	rect := text.BoundString(v.font, v.text)
	w, h := float64(rect.Dx()), float64(rect.Dy())

	if h < v.font.LineHeight() {
		h = v.font.LineHeight()
	}

	return w, h
}

func (v *LabelVisual) Draw(dest *ebiten.Image) {
	text.DrawWithOptions(dest, v.text, v.font, &ebiten.DrawImageOptions{
		GeoM:          v.Transform.GeoM(),
		ColorM:        v.ColorM,
		Filter:        v.font.Filter,
		CompositeMode: v.font.CompositeMode,
	})
}
