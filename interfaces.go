package igloo

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

// Dirtier allows structs to track when they are changed and only apply
// complex changes when there is something new to update.
type Dirtier interface {
	// IsDirty returns whether or not our object is dirty
	IsDirty() bool
	// Clean the object back to a fresh state
	Clean()
}

// Game adds a LoadScene method to ebitens Game interface allowing us to go from
// one scene to another.
type Game interface {
	ebiten.Game

	// LoadScene will unload the current scene and load in a new one.
	// It should dipose of the previous scene before setting up the next.
	// This is to allow a clean transition from one scene to another by allowing the
	// previous scene to be disposed before the new scene is setup.
	// Game Logic > Dispose current scene > Setup new scene.
	LoadScene(Scene) error
}

// Scene is a core component of logic and rendering.
// The core methods are update and draw to handle game logic and rendering.
type Scene interface {
	// Setup is used to handle all the loading of content and elements.
	// You should load all content and data here instead of a constructor function.
	// This way the LoadScene on the Game can transition between scenes.
	Setup(Game, fs.FS) error

	// Update all game elements.
	Update(deltaTime float64)

	// Draw all game elements.
	Draw(screen *ebiten.Image)

	// Dispose of game content or data.
	Dispose()
}

// Camera allows you to move through a large scene without having to modify each
// individual element. To allow for any type of camera ( or none at all ) instead
// of defining specifics on how the camera works or behaves in relation to elements
// we instead define how the elements behave with the camera.
//
// That is, if your element should interact with a possible moving camera have its
// world position modified to use the screen position from WorldToScreen. If your
// element should not be affected by a camera, such as a UI element, do not use the
// world to screen function.
//
// To avoid drawing anything off screen check if it is in view first.
type Camera interface {
	Dirtier

	// WorldToScreen will convert a world matrix into a screen matrix for rendering
	WorldToScreen(ebiten.GeoM) ebiten.GeoM

	// ScreenToWorld will convert an X,Y coordinate from the screen to its world position
	ScreenToWorld(x, y float64) (float64, float64)

	// IsInView returns whether or not a rectangle is within the cameras view
	IsInView(x, y, width, height float64) bool
}

// Canvaser represents a surface that images can be drawn to.
// Typically, this would be an *ebiten.Image.
type Canvaser interface {
	DrawImage(src *ebiten.Image, op *ebiten.DrawImageOptions)
}
