package mathf

type Bounds struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewBoundsWidthHeight(x, y, width, height float64) Bounds {
	return Bounds{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (b Bounds) Right() float64 {
	return b.X + b.Width
}

func (b Bounds) Bottom() float64 {
	return b.Y + b.Height
}

// Contains returns true if the point is inside our bounds
func (b Bounds) Contains(p Vec2) bool {
	return (b.X < p.X &&
		p.X < b.Right() &&
		b.Y < p.Y &&
		p.Y < b.Bottom())
}

// Overlaps returns true if we are even partially overlapping with the other bounds
func (b Bounds) Overlaps(o Bounds) bool {
	return (b.X < o.Right() &&
		o.X < b.Right() &&
		b.Y < o.Bottom() &&
		o.Y < b.Bottom())
}

// Inside returns true if our bounds is entirely inside the other bounds
func (b Bounds) Inside(o Bounds) bool {
	return (b.X > o.X &&
		b.Right() < o.Right() &&
		b.Y > o.Y &&
		b.Bottom() < o.Bottom())
}

// Surrounds returns true if our bounds entirely surrounds the other bounds
func (b Bounds) Surrounds(o Bounds) bool {
	return (b.X < o.X &&
		b.Right() > o.Right() &&
		b.Y < o.Y &&
		b.Bottom() > o.Bottom())
}

