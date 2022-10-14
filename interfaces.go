package igloo

import (
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

// Updater is the default updating method
type Updater interface {
	Update()
}

// Disposer lets content be disposed of properly
type Disposer interface {
	Dispose()
}

// Drawer will draw some visual to the destination
type Drawer interface {
	Draw(dest *ebiten.Image)
}

type NativeSizer interface {
	NativeSize() (float64, float64)
}

// Scene is a core component of logic and rendering.
// The core methods are update and draw to handle game logic and rendering.
type Scene interface {
	Updater
	Disposer
	Drawer

	Setup(assetLoader *AssetLoader) error
}

// PostSetup is an optional interface for scenes that will trigger after setup
// but before any update or draws so you can further refine the scene.
type PostSetup interface {
	PostSetup() error
}

// PreDispose is an optional interface for scenes that will trigger before the
// normal dispose when a scene is popped.
type PreDispose interface {
	PreDispose() error
}
