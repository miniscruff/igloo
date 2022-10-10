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

	label   *content.Label
	text    string
	isDirty bool
}

func NewLabelVisual(label *content.Label) *LabelVisual {
	v := &LabelVisual{
		label:   label,
		isDirty: false,
	}

	lineHeight := float64(text.BoundString(label, "A").Dy())

	v.Visualer = &igloo.Visualer{
		Transform: mathf.NewTransform(
			mathf.TransformWithFixedOffset(lineHeight),
		),
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

func (v *LabelVisual) Text() string {
	return v.text
}

func (v *LabelVisual) SetText(newText string) {
	if v.text == newText {
		return
	}

	v.text = newText
	v.isDirty = true

	// not ideal but best for now
	nw, nh := v.NativeSize()
	v.Transform.SetNaturalWidth(nw)
	v.Transform.SetNaturalHeight(nh)
	v.Transform.ResetScale()
}

func (v *LabelVisual) NativeSize() (float64, float64) {
	rect := text.BoundString(v.label, v.text)
	return float64(rect.Dx()), float64(rect.Dy())
}

func (v *LabelVisual) Draw(dest *ebiten.Image) {
	text.DrawWithOptions(dest, v.text, v.label, &ebiten.DrawImageOptions{
		GeoM:          v.Transform.GeoM(),
		ColorM:        v.ColorM,
		Filter:        v.label.Filter,
		CompositeMode: v.label.CompositeMode,
	})
}
