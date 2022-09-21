package graphics

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/miniscruff/igloo"
	"github.com/miniscruff/igloo/mathf"
)

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
	igloo.Dirtier

	// WorldToScreen will convert a world matrix into a screen matrix for rendering
	WorldToScreen(ebiten.GeoM) ebiten.GeoM

	// ScreenToWorld will convert an X,Y coordinate from the screen to its world position
	ScreenToWorld(screen image.Point) mathf.Vec2

	// IsInView returns whether or not a rectangle is within the cameras view.
	IsInView(bounds mathf.Bounds) bool
}

// Drawer is the default drawing method
type Drawer interface {
	Draw(dest *ebiten.Image, camera Camera)
}
