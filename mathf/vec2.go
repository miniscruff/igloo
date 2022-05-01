package mathf

import (
	"fmt"
	"image"
	"math"
)

// Vec2 describes a 2D vector or point in floats
type Vec2 struct {
	X, Y float64
}

var (
	// Vec2Zero is a Vec2 of (0, 0)
	Vec2Zero = Vec2{0, 0}
	// Vec2One is a Vec2 of (1, 1)
	Vec2One = Vec2{1, 1}
)

// String returns vec2 as a string
func (v *Vec2) String() string {
	return fmt.Sprintf("Vec2(%v, %v)", v.X, v.Y)
}

// Add other to us
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: v.X + other.X, Y: v.Y + other.Y}
}

// AddScalar adds scalar to both elements
func (v Vec2) AddScalar(scalar float64) Vec2 {
	return Vec2{X: v.X + scalar, Y: v.Y + scalar}
}

// Mul multiplies other to us
func (v Vec2) Mul(other Vec2) Vec2 {
	return Vec2{X: v.X * other.X, Y: v.Y * other.Y}
}

// MulScalar multiplies both elements by a scalar
func (v Vec2) MulScalar(scalar float64) Vec2 {
	return Vec2{X: v.X * scalar, Y: v.Y * scalar}
}

// Sub other from us
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{X: v.X - other.X, Y: v.Y - other.Y}
}

// SubScalar subtracts both elements by a scalar
func (v Vec2) SubScalar(scalar float64) Vec2 {
	return Vec2{X: v.X - scalar, Y: v.Y - scalar}
}

// Div other from us
func (v Vec2) Div(other Vec2) Vec2 {
	return Vec2{X: v.X / other.X, Y: v.Y / other.Y}
}

// SubScalar subtracts both elements by a scalar
func (v Vec2) DivScalar(scalar float64) Vec2 {
	return Vec2{X: v.X / scalar, Y: v.Y / scalar}
}

// Unit is a 1 unit vector in the same direction as v.
// Unless v is (0,0) in which case it returns (0,0).
func (v Vec2) Unit() Vec2 {
	mag := v.Mag()
	if mag == 0 {
		return Vec2Zero
	}

	return Vec2{v.X / mag, v.Y / mag}
}

// Mag returns the magnitude of our vector
func (v Vec2) Mag() float64 {
	return math.Hypot(v.X, v.Y)
}

// SqrMag returns the Square Magnitude of our vector
func (v Vec2) SqrMag() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Dist returns the distance between two vectors
func (v Vec2) Dist(other Vec2) float64 {
	return math.Sqrt(v.SqrDist(other))
}

// SqrDist returns the square distance between us and another vector
func (v Vec2) SqrDist(other Vec2) float64 {
	return math.Pow(v.X-other.X, 2) + math.Pow(v.Y-other.Y, 2)
}

// XY returns the X and Y components separately
func (v Vec2) XY() (float64, float64) {
	return v.X, v.Y
}

// Angle returns the angle in radians of our vector
func (v Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// Normal returns a vectors normal, same as rotating 90 degress
func (v Vec2) Normal() Vec2 {
	return Vec2{X: -v.Y, Y: v.X}
}

// Dot returns the dot product of vectors v and other
func (v Vec2) Dot(other Vec2) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Cross returns the cross product of vectors v and other
func (v Vec2) Cross(other Vec2) float64 {
	return v.X*other.X - v.Y*other.Y
}

func (v Vec2) ToPoint() image.Point {
	return image.Point{X: int(v.X), Y: int(v.Y)}
}

// Map applies a function to both X and Y components and
// returns a new Vec2 of the result
func (v Vec2) Map(fun func(float64) float64) Vec2 {
	return Vec2{
		X: fun(v.X),
		Y: fun(v.Y),
	}
}

// Vec2FromAngle returns a Vec2 from an angle in radians
func Vec2FromAngle(angle float64) Vec2 {
	sin, cos := math.Sincos(angle)
	return Vec2{X: cos, Y: sin}
}

// Vec2FromPoint returns a Vec2 from an image point
func Vec2FromPoint(pt image.Point) Vec2 {
	return Vec2{X: float64(pt.X), Y: float64(pt.Y)}
}

// Vec2Lerp returns a Vec2 as a linear interpolation between two vectors
func Vec2Lerp(start, end Vec2, percent float64) Vec2 {
	return Vec2{
		X: Lerp(start.X, end.X, percent),
		Y: Lerp(start.Y, end.Y, percent),
	}
}
