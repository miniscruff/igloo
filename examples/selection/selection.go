package selection

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/miniscruff/igloo"
	"github.com/miniscruff/igloo/examples/boxes"
)

type SelectionScene struct {
	game igloo.Game
}

func (s *SelectionScene) Setup(game igloo.Game) {
	ebiten.SetWindowTitle("Igloo Demo (selection)")
	s.game = game
}

func (s *SelectionScene) Update(deltaTime float64) {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		s.game.LoadScene(&boxes.BoxesScene{})
	}
}

func (s *SelectionScene) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "1. Boxes", 20, 20)
}

func (s *SelectionScene) Dispose() {
}
