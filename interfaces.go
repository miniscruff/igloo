package igloo

import "github.com/hajimehoshi/ebiten/v2"

// Dirtier allows structs to track when they are changed and only apply
// complex changes when there is something new to update.
type Dirtier interface {
	IsDirty() bool
	Clean()
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
// If you need to convert a screen position to world, such as a mouse position use
// ScreenToWorld.
//
// To avoid drawing anything off screen check if it is in view first.
//
// Camera also implements Dirter to avoid checks or calculation if the camera has not
// moved.
type Camera interface {
	Dirtier
	WorldToScreen(ebiten.GeoM) ebiten.GeoM
	ScreenToWorld(x, y float64) (float64, float64)
	IsInView(x, y, width, height float64) bool
}

// Canvaser represents a surface that images can be drawn to.
// Typically, this would be an *ebiten.Image.
type Canvaser interface {
	DrawImage(src *ebiten.Image, op *ebiten.DrawImageOptions)
}
