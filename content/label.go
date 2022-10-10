package content

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type Label struct {
	font.Face
	ebiten.CompositeMode
	ebiten.Filter
}
