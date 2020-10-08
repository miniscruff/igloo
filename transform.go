package igloo

// Transform holds all data associated to a location
type Transform struct {
	x        float64
	y        float64
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
	return t.x
}

// SetX will change the x value and mark us as dirty.
func (t *Transform) SetX(x float64) {
	t.x = x
	t.isDirty = true
}

// Y will return the y value
func (t *Transform) Y() float64 {
	return t.y
}

// SetY will change the y value and mark us as dirty.
func (t *Transform) SetY(y float64) {
	t.y = y
	t.isDirty = true
}

// Rotation will return the rotation value
func (t *Transform) Rotation() float64 {
	return t.rotation
}

// SetRotation will change the rotation value and mark us as dirty.
func (t *Transform) SetRotation(rotation float64) {
	t.rotation = rotation
	t.isDirty = true
}

// GetPosition will return both x and y
func (t *Transform) GetPosition() (float64, float64) {
	return t.x, t.y
}

// SetPosition will change both x and y values as well as marking us as dirty.
func (t *Transform) SetPosition(x, y float64) {
	t.x = x
	t.y = y
	t.isDirty = true
}

// Translate will move x and y by dx and dy as well as marking as dirty.
func (t *Transform) Translate(dx, dy float64) {
	t.x += dx
	t.y += dy
	t.isDirty = true
}

// NewTransform will create a new transform from x,y and rotation.
// Note that transforms start dirty.
func NewTransform(x, y, rotation float64) *Transform {
	return &Transform{
		x:        x,
		y:        y,
		rotation: rotation,
		isDirty:  true, // start dirty
	}
}
