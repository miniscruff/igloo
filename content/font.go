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

	fixedOffset float64
}

func (f *Font) FixedOffset() float64 {
	if f.fixedOffset == 0 {
		f.fixedOffset = float64(text.BoundString(f.Face, "A").Dy())
	}

	return f.fixedOffset
}
