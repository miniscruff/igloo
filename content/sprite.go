package content

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	*ebiten.Image
	ebiten.ColorM
	ebiten.CompositeMode
	ebiten.Filter
}
