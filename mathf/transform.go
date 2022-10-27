package mathf

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Transform holds all data associated to a location
type Transform struct {
	position      Vec2    // x y of our pivot if not stretching
	width         float64 // width before scaling if not stretched
	height        float64 // height before scaling if not stretched
	rotation      float64 // in radians
	pivot         Vec2    // normalized position where our transform is rotated and positioned around
	anchors       Sides   // normalized position in the parent Transform of our corners
	offsets       Sides   // absolute offsets of the corners relataive to ouor parent
	fixedOffset   float64 // text is drawn from the bottom, this option adjusts the pivot
	naturalWidth  float64 // native width of our source graphic
	naturalHeight float64 // native height of our source graphic

	bounds  Bounds      // calculated bounding box
	geom    ebiten.GeoM // calculated geom matrix
	isDirty bool
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

func (t *Transform) Pivot() Vec2 {
	return t.pivot
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

func (t *Transform) Anchors() Sides {
	return t.anchors
}

func (t *Transform) Offsets() Sides {
	return t.offsets
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

func (t *Transform) SetPivot(pivot Vec2) {
	if t.pivot == pivot {
		return
	}

	t.pivot = pivot
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

func (t *Transform) SetSize(width, height float64) {
	t.SetWidth(width)
	t.SetHeight(height)
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

func (t *Transform) SetFixedOffset(offset float64) {
	if t.fixedOffset == offset {
		return
	}

	t.fixedOffset = offset
	t.isDirty = true
}

// SetPosition will set our position to a different one
func (t *Transform) SetPosition(pos Vec2) {
	t.SetX(pos.X)
	t.SetY(pos.Y)
}

func (t *Transform) SetAnchors(sides Sides) {
	if t.anchors == sides {
		return
	}

	t.anchors = sides
	t.isDirty = true
}

func (t *Transform) SetOffsets(sides Sides) {
	if t.offsets == sides {
		return
	}

	t.offsets = sides
	t.isDirty = true
}

func (t *Transform) SetLeftAnchor(value float64) {
	if t.anchors.Left == value {
		return
	}

	t.anchors.Left = value
	t.isDirty = true
}

func (t *Transform) SetRightAnchor(value float64) {
	if t.anchors.Right == value {
		return
	}

	t.anchors.Right = value
	t.isDirty = true
}

func (t *Transform) SetTopAnchor(value float64) {
	if t.anchors.Top == value {
		return
	}

	t.anchors.Top = value
	t.isDirty = true
}

func (t *Transform) SetBottomAnchor(value float64) {
	if t.anchors.Bottom == value {
		return
	}

	t.anchors.Bottom = value
	t.isDirty = true
}

func (t *Transform) SetLeftOffset(value float64) {
	if t.offsets.Left == value {
		return
	}

	t.offsets.Left = value
	t.isDirty = true
}

func (t *Transform) SetRightOffset(value float64) {
	if t.offsets.Right == value {
		return
	}

	t.offsets.Right = value
	t.isDirty = true
}

func (t *Transform) SetTopOffset(value float64) {
	if t.offsets.Top == value {
		return
	}

	t.offsets.Top = value
	t.isDirty = true
}

func (t *Transform) SetBottomOffset(value float64) {
	if t.offsets.Bottom == value {
		return
	}

	t.offsets.Bottom = value
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

func (t *Transform) Size() (float64, float64) {
	// TODO: should use anchor values if stretched
	return t.width, t.height
}

func (t *Transform) NaturalSize() (float64, float64) {
	return t.naturalWidth, t.naturalHeight
}

// Bounds of our transform
// NOTE: does not take into account rotation yet...
func (t *Transform) Bounds() Bounds {
	return t.bounds
}

func (t *Transform) InView(other *Transform) bool {
	return t.Bounds().Overlaps(other.Bounds())
}

// Build updates our ebiten.GeoM with updated values.
// Be sure to only build a new GeoM if we need to.
// Offset is the current world offset we should apply to our local values.
func (t *Transform) Build() {
	t.geom.Reset()

	/*
		if t.anchorMin == t.anchorMax {
			// TODO: world offset scale
			if t.width != t.naturalWidth || t.height != t.naturalHeight {
				t.geom.Scale(t.width/t.naturalWidth, t.height/t.naturalHeight)
			}

			t.geom.Translate(-t.width*t.pivot.X, -t.height*t.pivot.Y+t.fixedOffset)

			if t.rotation != 0 {
				t.geom.Rotate(t.rotation)
			}

			// log.Printf("xy, offset: %v, %v at %v", t.position, offset.position, t.position.Add(offset.position))
			t.geom.Translate(t.X(), t.Y())
			// t.geom.Translate(offset.X(), offset.Y())
		} else {
			pivotOffset := t.pivot.Mul(Vec2{X: t.width, Y: t.height})
			topLeft := t.position.Sub(pivotOffset)

			t.bounds = NewBoundsWidthHeight(
				topLeft.X,
				topLeft.Y,
				t.width,
				t.height,
			)
		}
	*/
}

// NewTransform will create a new transform with:
// Note that transforms start dirty.
func NewTransform() *Transform {
	t := &Transform{
		position:    Vec2Zero,
		rotation:    0,
		pivot:       Vec2TopLeft,
		anchors:     Sides{},
		offsets:     Sides{},
		fixedOffset: 0,
		geom:        ebiten.GeoM{},
		isDirty:     true,
	}

	return t
}

// Transform tweens

func NewRotationTween(
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

func NewPositionTween(
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

func NewWidthTween(
	target *Transform,
	start, end float64,
	duration float64,
	options ...TweenOption,
) *Tween {
	t := NewTween(
		duration,
		TweenUpdateFunc(func(value float64) {
			target.SetWidth(Lerp(start, end, value))
		}),
	)

	for _, opt := range options {
		opt(t)
	}

	return t
}
