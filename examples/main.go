package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo"
	"github.com/miniscruff/igloo/examples/selection"
	"math"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

type Game struct {
	scene igloo.Scene
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	if g.scene == nil {
		return nil
	}

	deltaTime := 1 / ebiten.CurrentTPS()
	if math.IsInf(deltaTime, 0) {
		return nil
	}

	g.scene.Update(deltaTime)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.scene == nil {
		return
	}
	deltaTime := 1.0 / ebiten.CurrentTPS()
	if math.IsInf(deltaTime, 0) {
		return
	}
	g.scene.Draw(screen)
}

func (g *Game) LoadScene(newScene igloo.Scene) {
	if g.scene != nil {
		g.scene.Dispose()
	}
	g.scene = newScene
	g.scene.Setup(g)
}

func main() {
	ebiten.SetWindowTitle("Igloo Demo")
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowSize(screenWidth, screenHeight)

	g := &Game{}
	g.LoadScene(&selection.SelectionScene{})
	ebiten.RunGame(g)
}
