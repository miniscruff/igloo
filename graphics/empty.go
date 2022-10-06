package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo"
	"github.com/miniscruff/igloo/mathf"
)

type EmptyVisual struct {
	*igloo.Visualer
}

func NewEmptyVisual() *EmptyVisual {
	ev := &EmptyVisual{}
	ev.Visualer = &igloo.Visualer{
		Transform:   mathf.NewTransform(),
		Children:    make([]*igloo.Visualer, 0),
		Dirtier:     ev,
		Drawer:      ev,
		NativeSizer: ev,
	}
	return ev
}

func (v *EmptyVisual) IsDirty() bool {
	return false
}

func (v *EmptyVisual) Clean() {
}

func (v *EmptyVisual) NativeSizer() (float64, float64) {
	return 0, 0
}

func (v *EmptyVisual) Draw(dest *ebiten.Image) {
}
