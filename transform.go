package igloo

// Transform holds all data associated to a location
type Transform struct {
	position Vec2
	rotation float64
	isDirty  bool
}

// IsDirty returns whether or not we have changed since the last update
func (t *Transform) IsDirty() bool {
	return t.isDirty
}

// Clean will clean up the dirty status back to false
func (t *Transform) Clean() {
	t.isDirty = false
}

// X will return the x value
func (t *Transform) X() float64 {
	return t.position.X
}

// SetX will change the x value and mark us as dirty.
func (t *Transform) SetX(x float64) {
	if t.position.X == x {
		return
	}

	t.position.X = x
	t.isDirty = true
}

// Y will return the y value
func (t *Transform) Y() float64 {
	return t.position.Y
}

// SetY will change the y value and mark us as dirty.
func (t *Transform) SetY(y float64) {
	if t.position.Y == y {
		return
	}

	t.position.Y = y
	t.isDirty = true
}

// Rotation will return the rotation value
func (t *Transform) Rotation() float64 {
	return t.rotation
}

// SetRotation will change the rotation value and mark us as dirty.
func (t *Transform) SetRotation(rotation float64) {
	if t.rotation == rotation {
		return
	}

	t.rotation = rotation
	t.isDirty = true
}

// Position will return our position vector
func (t *Transform) Position() Vec2 {
	return t.position
}

// SetPosition will set our position to a different one
func (t *Transform) SetPosition(pos Vec2) {
	t.SetX(pos.X)
	t.SetY(pos.Y)
}

// Translate will move x and y by our vec2
func (t *Transform) Translate(delta Vec2) {
	t.position.X += delta.X
	t.position.Y += delta.Y
	t.isDirty = true
}

// TranslateX moves us in the X axis
func (t *Transform) TranslateX(x float64) {
	t.position.X += x
	t.isDirty = true
}

// TranslateY moves us in the Y axis
func (t *Transform) TranslateY(y float64) {
	t.position.Y += y
	t.isDirty = true
}

// NewTransform will create a new transform from x,y and rotation.
// Note that transforms start dirty.
func NewTransform(position Vec2, rotation float64) *Transform {
	return &Transform{
		position: position,
		rotation: rotation,
		isDirty:  true, // start dirty
	}
}
