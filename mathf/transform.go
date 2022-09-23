package mathf

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Transform holds all data associated to a location
type Transform struct {
	position Vec2
	rotation float64
	anchor   Vec2
	// text is drawn from the bottom, this option adjusts the anchor
	fromBottom    bool
	width         float64
	height        float64
	naturalWidth  float64
	naturalHeight float64
	geom          ebiten.GeoM
	isDirty       bool
}

// IsDirty returns whether or not we have changed since the last update
func (t *Transform) IsDirty() bool {
	return t.isDirty
}

// Clean will clean up the dirty status back to false
func (t *Transform) Clean() {
	t.isDirty = false
}

func (t *Transform) GeoM() ebiten.GeoM {
	if t.isDirty {
		t.build()
		t.Clean()
	}

	return t.geom
}

// X will return the x value
func (t *Transform) X() float64 {
	return t.position.X
}

// Y will return the y value
func (t *Transform) Y() float64 {
	return t.position.Y
}

// Rotation will return the rotation value
func (t *Transform) Rotation() float64 {
	return t.rotation
}

func (t *Transform) Anchor() Vec2 {
	return t.anchor
}

func (t *Transform) Width() float64 {
	return t.width
}

func (t *Transform) Height() float64 {
	return t.height
}

func (t *Transform) NaturalWidth() float64 {
	return t.naturalWidth
}

func (t *Transform) NaturalHeight() float64 {
	return t.naturalHeight
}

// Position will return our position vector
func (t *Transform) Position() Vec2 {
	return t.position
}

// SetX will change the x value and mark us as dirty.
func (t *Transform) SetX(x float64) {
	if t.position.X == x {
		return
	}

	t.position.X = x
	t.isDirty = true
}

// SetY will change the y value and mark us as dirty.
func (t *Transform) SetY(y float64) {
	if t.position.Y == y {
		return
	}

	t.position.Y = y
	t.isDirty = true
}

// SetRotation will change the rotation value and mark us as dirty.
func (t *Transform) SetRotation(rotation float64) {
	if t.rotation == rotation {
		return
	}

	t.rotation = rotation
	t.isDirty = true
}

func (t *Transform) SetAnchor(anchor Vec2) {
	if t.anchor == anchor {
		return
	}

	t.anchor = anchor
	t.isDirty = true
}

func (t *Transform) SetWidth(width float64) {
	if t.width == width {
		return
	}
	if width < 0 {
		width = 0
	}

	t.width = width
	t.isDirty = true
}

func (t *Transform) SetHeight(height float64) {
	if t.height == height {
		return
	}
	if height < 0 {
		height = 0
	}

	t.height = height
	t.isDirty = true
}

// SetNaturalWidth will set the natural width
func (t *Transform) SetNaturalWidth(naturalWidth float64) {
	if t.naturalWidth == naturalWidth {
		return
	}

	t.naturalWidth = naturalWidth
	t.isDirty = true
}

// SetNaturalHeight will set the natural height
func (t *Transform) SetNaturalHeight(naturalHeight float64) {
	if t.naturalHeight == naturalHeight {
		return
	}

	t.naturalHeight = naturalHeight
	t.isDirty = true
}

// ResetScale will set the width and height values to match the natural width
// and height values.
func (t *Transform) ResetScale() {
	if t.width == t.naturalWidth && t.height == t.naturalHeight {
		return
	}

	t.width = t.naturalWidth
	t.height = t.naturalHeight
	t.isDirty = true
}

// SetPosition will set our position to a different one
func (t *Transform) SetPosition(pos Vec2) {
	t.SetX(pos.X)
	t.SetY(pos.Y)
}

// Translate will move x and y by our vec2
func (t *Transform) Translate(delta Vec2) {
	if delta == Vec2Zero {
		return
	}

	t.position.X += delta.X
	t.position.Y += delta.Y
	t.isDirty = true
}

// TranslateX moves us in the X axis
func (t *Transform) TranslateX(x float64) {
	if x == 0 {
		return
	}

	t.position.X += x
	t.isDirty = true
}

// TranslateY moves us in the Y axis
func (t *Transform) TranslateY(y float64) {
	if y == 0 {
		return
	}

	t.position.Y += y
	t.isDirty = true
}

// Size of our transform
func (t *Transform) Size() (float64, float64) {
	return t.width, t.height
}

// Bounds of our transform
// NOTE: does not take into account rotation yet...
func (t *Transform) Bounds() Bounds {
	anchorOffset := t.anchor.Mul(Vec2{X: t.width, Y: t.height})
	topLeft := t.position.Sub(anchorOffset)

	return NewBoundsWidthHeight(
		topLeft.X,
		topLeft.Y,
		t.width,
		t.height,
	)
}

// build updates our ebiten.GeoM with updated values if
// our transform is dirty.
func (t *Transform) build() {
	t.geom.Reset()

	if t.width != t.naturalWidth || t.height != t.naturalHeight {
		t.geom.Scale(t.width/t.naturalWidth, t.height/t.naturalHeight)
	}

	// handle flipping the y anchor here
	ay := t.anchor.Y
	if t.fromBottom {
		ay--
	}

	t.geom.Translate(-t.width*t.anchor.X, -t.height*ay)

	if t.rotation != 0 {
		t.geom.Rotate(t.rotation)
	}

	t.geom.Translate(t.X(), t.Y())
}

type TransformOption func(t *Transform)

func TransformAtPosition(position Vec2) TransformOption {
	return func(t *Transform) {
		t.position = position
	}
}

func TransformAtXY(x, y float64) TransformOption {
	return func(t *Transform) {
		t.position.X = x
		t.position.Y = y
	}
}

func TransformWithRotation(rotation float64) TransformOption {
	return func(t *Transform) {
		t.rotation = rotation
	}
}

func TransformWithWidth(width float64) TransformOption {
	return func(t *Transform) {
		t.width = width
	}
}

func TransformWithHeight(height float64) TransformOption {
	return func(t *Transform) {
		t.height = height
	}
}

// TransformWithNaturalWidth will set the transform natural width value
// this will also override the current height to match
func TransformWithNaturalWidth(naturalWidth float64) TransformOption {
	return func(t *Transform) {
		t.naturalWidth = naturalWidth
		t.width = naturalWidth
	}
}

// TransformWithNaturalHeight will set the transform natural height value
// this will also override the current height to match
func TransformWithNaturalHeight(naturalHeight float64) TransformOption {
	return func(t *Transform) {
		t.naturalHeight = naturalHeight
		t.height = naturalHeight
	}
}

func TransformWithSize(width, height float64) TransformOption {
	return func(t *Transform) {
		t.width = width
		t.height = height
	}
}

func TransformWithNaturalSize(naturalWidth, naturalHeight float64) TransformOption {
	return func(t *Transform) {
		t.naturalWidth = naturalWidth
		t.naturalHeight = naturalHeight
		t.width = naturalWidth
		t.height = naturalHeight
	}
}

func TransformWithAnchor(anchor Vec2) TransformOption {
	return func(t *Transform) {
		t.anchor = anchor
	}
}

func TransformDrawFromBottom() TransformOption {
	return func(t *Transform) {
		t.fromBottom = true
	}
}

// NewTransform will create a new transform with:
// * position at 0,0,
// * rotation of 0 degress
// * anchor at top left
// * width, height, natural width and natural height are all 0
// * you will have to set these before the first build
// Note that transforms start dirty.
func NewTransform(options ...TransformOption) *Transform {
	t := &Transform{
		position:   Vec2Zero,
		rotation:   0,
		anchor:     AnchorTopLeft,
		fromBottom: false,
		geom:       ebiten.GeoM{},
		isDirty:    true,
	}

	for _, o := range options {
		o(t)
	}

	return t
}

// Transform tweens

func NewRotationClip(
	target *Transform,
	start, end, duration float64,
	useRelative bool,
	options ...TweenOption,
) *Tween {
	startRotation := start
	endRotation := end

	t := NewTween(
		duration,
		TweenUpdateFunc(func(value float64) {
			target.SetRotation(Lerp(startRotation, endRotation, value))
		}),
	)

	if useRelative {
		TweenOnStart(func(t *Tween) {
			startRotation = target.Rotation() + start
			endRotation = target.Rotation() + end
		})(t)
	}

	for _, opt := range options {
		opt(t)
	}

	return t
}

func NewPositionClip(
	target *Transform,
	start, end Vec2,
	duration float64,
	useRelative bool,
	options ...TweenOption,
) *Tween {
	startPosition := start
	endPosition := end

	t := NewTween(
		duration,
		TweenUpdateFunc(func(value float64) {
			target.SetPosition(Vec2Lerp(startPosition, endPosition, value))
		}),
	)

	if useRelative {
		TweenOnStart(func(t *Tween) {
			startPosition = target.Position().Add(start)
			endPosition = target.Position().Add(end)
		})(t)
	}

	for _, opt := range options {
		opt(t)
	}

	return t
}
