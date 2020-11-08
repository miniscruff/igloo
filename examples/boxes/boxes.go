package boxes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/miniscruff/igloo"
	"github.com/miniscruff/igloo/examples/components"
	"image/color"
)

type BoxesScene struct {
	game        igloo.Game
	camera      igloo.Camera
	blankImage  *ebiten.Image
	boxes       []*igloo.Sprite
	isDragging  bool
	draggingBox *igloo.Sprite
}

func (s *BoxesScene) Setup(game igloo.Game) {
	ebiten.SetWindowTitle("Igloo Demo (boxes)")
	s.game = game
	s.camera = &components.StaticCamera{}
	s.blankImage = ebiten.NewImage(1, 1)
	s.blankImage.Set(0, 0, color.RGBA{255, 255, 255, 255})

	s.draggingBox = igloo.NewSprite(s.blankImage, igloo.NewTransform(0, 0, 0))
}

// Update all game elements
func (s *BoxesScene) Update(deltaTime float64) {
	cx, cy := ebiten.CursorPosition()
	mouseX := float64(cx)
	mouseY := float64(cy)
	if !s.isDragging && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.isDragging = true
		s.draggingBox.Transform.SetPosition(mouseX, mouseY)
		s.draggingBox.SetWidth(1)
		s.draggingBox.SetHeight(1)
	} else if s.isDragging && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		s.isDragging = false
		s.boxes = append(s.boxes, s.draggingBox)
		s.draggingBox = igloo.NewSprite(s.blankImage, igloo.NewTransform(0, 0, 0))
	} else if s.isDragging {
		if mouseX < s.draggingBox.Transform.X() {
			s.draggingBox.SetWidth(s.draggingBox.Transform.X() - mouseX + s.draggingBox.Width())
		}
		if mouseY < s.draggingBox.Transform.Y() {
			s.draggingBox.SetHeight(s.draggingBox.Transform.Y() - mouseY - s.draggingBox.Height())
		}
		s.draggingBox.SetWidth(mouseX - s.draggingBox.Transform.X())
		s.draggingBox.SetHeight(mouseY - s.draggingBox.Transform.Y())
	}
}

// Draw all game elements
func (s *BoxesScene) Draw(screen *ebiten.Image) {
	if s.isDragging {
		s.draggingBox.Draw(screen, s.camera)
	}
	for _, b := range s.boxes {
		b.Draw(screen, s.camera)
	}
	ebitenutil.DebugPrintAt(screen, "left click and drag to create boxes", 20, 20)
}

// Dispose of game content or data
func (s *BoxesScene) Dispose() {
	s.blankImage.Dispose()
}
