package igloo

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo/mathf"
)

// Dirtier allows structs to track when they are changed and only apply
// complex changes when there is something new to update.
type Dirtier interface {
	// IsDirty returns whether or not our object is dirty
	IsDirty() bool
	// Clean the object back to a fresh state
	Clean()
}

// Scene is a core component of logic and rendering.
// The core methods are update and draw to handle game logic and rendering.
type Scene interface {
	Updater
	Disposer

	// Draw all game elements.
	Draw(screen *ebiten.Image)
}

// Updater is the default updating method
type Updater interface {
	Update(*mathf.GameTime)
}

// Disposer lets content be disposed of properly
type Disposer interface {
	Dispose()
}

var (
	game      *Game
	exit      bool
	exitError = errors.New("exiting game")
)

func init() {
	game = &Game{
		gameTime: mathf.NewGameTime(),
	}
	exit = false
}

type Game struct {
	ebiten.Game
	gameTime *mathf.GameTime
	scenes   []Scene
}

// Todo: we will want to set and update stuff on this func...
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return 800, 600
}

// Update the top scene of the stack
func (g *Game) Update() error {
	g.gameTime.Tick()

	// do not update if our game is running slowly
	if g.gameTime.IsSlow() {
		fmt.Printf("game is running slowly: %v\n", g.gameTime.DeltaTime())
		return nil
	}

	lastScene := g.scenes[len(g.scenes)-1]
	lastScene.Update(g.gameTime)

	if exit {
		return exitError
	}

	return nil
}

// Draw all the game scenes, bottom up
func (g *Game) Draw(dest *ebiten.Image) {
	for _, s := range g.scenes {
		s.Draw(dest)
	}
}

// Push a new scene to the top of the stack
func Push(scene Scene) {
	game.scenes = append(game.scenes, scene)
}

// Pop a scene off the stack
func Pop() {
	lastScene := game.scenes[len(game.scenes)-1]
	lastScene.Dispose()

	game.scenes = game.scenes[:len(game.scenes)-1]
}

// Exit the game at the end of the next update
func Exit() {
	exit = true
}

func Run() error {
	return ebiten.RunGame(game)
}
