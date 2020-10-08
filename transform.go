package igloo

// Transform holds all data associated to our location
type Transform struct {
	x        float64
	y        float64
	rotation float64
	isDirty  bool
}

func (t *Transform) IsDirty() bool {
	return t.isDirty
}

func (t *Transform) Clean() {
	t.isDirty = false
}

func (t *Transform) X() float64 {
	return t.x
}

func (t *Transform) SetX(x float64) {
	t.x = x
	t.isDirty = true
}

func (t *Transform) Y() float64 {
	return t.y
}

func (t *Transform) SetY(y float64) {
	t.y = y
	t.isDirty = true
}

func (t *Transform) Rotation() float64 {
	return t.rotation
}

func (t *Transform) SetRotation(rotation float64) {
	t.rotation = rotation
	t.isDirty = true
}

func (t *Transform) GetPosition() (float64, float64) {
	return t.x, t.y
}

func (t *Transform) SetPosition(x, y float64) {
	t.x = x
	t.y = y
	t.isDirty = true
}

func (t *Transform) Translate(dx, dy float64) {
	t.x += dx
	t.y += dy
	t.isDirty = true
}

func NewTransform(x, y, rotation float64) *Transform {
	return &Transform{
		x:        x,
		y:        y,
		rotation: rotation,
		isDirty:  true, // start dirty
	}
}
