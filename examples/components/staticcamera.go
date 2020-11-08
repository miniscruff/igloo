package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// StaticCamera is a fixed camera that does not move or effect positions.
// Useful for demos or user interface rendering.
type StaticCamera struct {
}

func (m *StaticCamera) IsDirty() bool {
	return false
}

func (m *StaticCamera) Clean() {
}

func (m *StaticCamera) WorldToScreen(geom ebiten.GeoM) ebiten.GeoM {
	return geom
}

func (m *StaticCamera) ScreenToWorld(x, y float64) (float64, float64) {
	return x, y
}

func (m *StaticCamera) IsInView(x, y, width, height float64) bool {
	return true
}
