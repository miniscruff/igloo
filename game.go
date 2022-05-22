package igloo

import (
	"errors"

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
	game    *Game
	exit    bool
	errExit = errors.New("exiting game")
)

type Game struct {
	gameTime      *mathf.GameTime
	scenes        []Scene
	outsideWidth  int
	outsideHeight int
	screenWidth   int
	screenHeight  int
	windowWidth   int
	windowHeight  int
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	g.outsideWidth = outsideWidth
	g.outsideHeight = outsideHeight

	return g.screenWidth, g.screenHeight
}

func GetOutsideSize() (int, int) {
	return game.outsideWidth, game.outsideHeight
}

func GetWindowSize() (int, int) {
	return game.windowWidth, game.windowHeight
}

func GetScreenSize() (int, int) {
	return game.screenWidth, game.screenHeight
}

func Width() int {
	return game.screenWidth
}

func Height() int {
	return game.screenHeight
}

func SetScreenSize(w, h int) {
	game.screenWidth = w
	game.screenHeight = h
}

func SetWindowSize(w, h int) {
	game.windowWidth = w
	game.windowHeight = h
	ebiten.SetWindowSize(w, h)
}

// Update the top scene of the stack
func (g *Game) Update() error {
	lastScene := g.scenes[len(g.scenes)-1]
	lastScene.Update(g.gameTime)

	if exit {
		return errExit
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

func InitGame() {
	game = &Game{
		gameTime:     mathf.NewGameTime(),
		screenWidth:  800,
		screenHeight: 600,
	}
	exit = false
}

func Run() error {
	return ebiten.RunGame(game)
}
