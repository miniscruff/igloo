package igloo

import (
	"errors"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/miniscruff/igloo/mathf"
)

var (
	game    *Game
	exit    bool
	errExit = errors.New("exiting game")
)

type SceneContext struct {
	Scene  Scene
	Ticker *mathf.Ticker
}

// GameConfig contains values you should set when initializing
// that can only be configured at start.
type GameConfig struct {
	Fsys       fs.FS
	AssetsPath string
}

type Game struct {
	scenes      []*SceneContext
	assetLoader *AssetLoader

	// window values
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
	lastScene.Ticker.Tick()
	lastScene.Scene.Update()

	if exit {
		return errExit
	}

	return nil
}

// Draw all the game scenes, bottom up
func (g *Game) Draw(dest *ebiten.Image) {
	for _, s := range g.scenes {
		s.Scene.Draw(dest)
	}
}

// Push a new scene to the top of the stack
func Push(scene Scene) {
	context := &SceneContext{
		Scene:  scene,
		Ticker: mathf.NewTicker(),
	}

	scene.Setup(game.assetLoader)

	// force an update as well as it will be the newest scene
	scene.Update()

	game.scenes = append(game.scenes, context)
}

// Pop a scene off the stack
func Pop() {
	lastScene := game.scenes[len(game.scenes)-1]
	lastScene.Scene.Dispose()
	lastScene.Ticker = nil

	game.scenes = game.scenes[:len(game.scenes)-1]
}

// Exit the game at the end of the next update
func Exit() {
	exit = true
}

func InitGame(config GameConfig) {
	game = &Game{
		screenWidth:  800,
		screenHeight: 600,
		assetLoader:  NewAssetLoader(config.Fsys, config.AssetsPath),
	}
	exit = false
}

func Run() error {
	return ebiten.RunGame(game)
}
