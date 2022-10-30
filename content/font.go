package content

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Font struct {
	font.Face
	ebiten.CompositeMode
	ebiten.Filter

	ascent     float64
	lineHeight float64
}

func (f *Font) Ascent() float64 {
	if f.ascent == 0 {
		rect := text.BoundString(f.Face, "A")
		f.ascent = float64(rect.Dy())
	}

	return f.ascent
}

func (f *Font) LineHeight() float64 {
	if f.lineHeight == 0 {
		rect := text.BoundString(f.Face, "Aj")
		f.lineHeight = float64(rect.Dy())
	}

	return f.lineHeight
}
