package igloo

import (
	"fmt"
	"math"
)

// Vec2f describes a 2D vector or point in floats
type Vec2f struct {
	X, Y float64
}

var (
	// Vec2fZero is a Vec2f of (0, 0)
	Vec2fZero = Vec2f{0, 0}
	// Vec2fOne is a Vec2f of (1, 1)
	Vec2fOne = Vec2f{1, 1}
)

// String returns vec2 as a string
func (v *Vec2f) String() string {
	return fmt.Sprintf("Vec2f(%v, %v)", v.X, v.Y)
}

// Add other to us
func (v Vec2f) Add(other Vec2f) Vec2f {
	return Vec2f{X: v.X + other.X, Y: v.Y + other.Y}
}

// MulScalar multiplies both elements by a scalar
func (v Vec2f) MulScalar(scalar float64) Vec2f {
	return Vec2f{X: v.X * scalar, Y: v.Y * scalar}
}

// Sub other from us
func (v Vec2f) Sub(other Vec2f) Vec2f {
	return Vec2f{X: v.X - other.X, Y: v.Y - other.Y}
}

// SubScalar subtracts both elements by a scalar
func (v Vec2f) SubScalar(scalar float64) Vec2f {
	return Vec2f{X: v.X - scalar, Y: v.Y - scalar}
}

// Unit is a 1 unit vector in the same direction as v.
// Unless v is (0,0) in which case it returns (0,0).
func (v Vec2f) Unit() Vec2f {
	mag := v.Mag()
	if mag == 0 {
		return Vec2fZero
	}

	return Vec2f{v.X / mag, v.Y / mag}
}

// Mag returns the magnitude of our vector
func (v Vec2f) Mag() float64 {
	return math.Hypot(v.X, v.Y)
}

// SqrMag returns the Square Magnitude of our vector
func (v Vec2f) SqrMag() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Dist returns the distance between two vectors
func (v Vec2f) Dist(other Vec2f) float64 {
	return math.Sqrt(v.SqrDist(other))
}

// SqrDist returns the square distance between us and another vector
func (v Vec2f) SqrDist(other Vec2f) float64 {
	return math.Pow(v.X-other.X, 2) + math.Pow(v.Y-other.Y, 2)
}

// XY returns the X and Y components separately
func (v Vec2f) XY() (float64, float64) {
	return v.X, v.Y
}

// Angle returns the angle in radians of our vector
func (v Vec2f) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

// Normal returns a vectors normal, same as rotating 90 degress
func (v Vec2f) Normal() Vec2f {
	return Vec2f{X: -v.Y, Y: v.X}
}

// Dot returns the dot product of vectors v and other
func (v Vec2f) Dot(other Vec2f) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Cross returns the cross product of vectors v and other
func (v Vec2f) Cross(other Vec2f) float64 {
	return v.X*other.X - v.Y*other.Y
}

// ToVec2i converts to integers
func (v Vec2f) ToVec2i() Vec2i {
	return Vec2i{X: int(v.X), Y: int(v.Y)}
}

// Map applies a function to both X and Y components and
// returns a new Vec2f of the result
func (v Vec2f) Map(fun func(float64) float64) Vec2f {
	return Vec2f{
		X: fun(v.X),
		Y: fun(v.Y),
	}
}

// Vec2fFromAngle returns a Vec2f from an angle in radians
func Vec2fFromAngle(angle float64) Vec2f {
	sin, cos := math.Sincos(angle)
	return Vec2f{X: cos, Y: sin}
}

// Vec2i describes a 2D vector or point in ints
type Vec2i struct {
	X, Y int
}

var (
	// Vec2iZero is a Vec2i of (0, 0)
	Vec2iZero = Vec2i{0, 0}
	// Vec2iOne is a Vec2i of (1, 1)
	Vec2iOne = Vec2i{1, 1}
)

// String returns vec2 as a string
func (v *Vec2i) String() string {
	return fmt.Sprintf("Vec2i(%v, %v)", v.X, v.Y)
}

// Add other to us
func (v Vec2i) Add(other Vec2i) Vec2i {
	return Vec2i{X: v.X + other.X, Y: v.Y + other.Y}
}

// MulScalar multiplies both elements by a scalar
func (v Vec2i) MulScalar(scalar int) Vec2i {
	return Vec2i{X: v.X * scalar, Y: v.Y * scalar}
}

// Sub other from us
func (v Vec2i) Sub(other Vec2i) Vec2i {
	return Vec2i{X: v.X - other.X, Y: v.Y - other.Y}
}

// SubScalar subtracts both elements by a scalar
func (v Vec2i) SubScalar(scalar int) Vec2i {
	return Vec2i{X: v.X - scalar, Y: v.Y - scalar}
}

// Mag returns the magnitude of our vector
func (v Vec2i) Mag() float64 {
	return math.Hypot(float64(v.X), float64(v.Y))
}

// SqrMag returns the Square Magnitude of our vector
func (v Vec2i) SqrMag() int {
	return v.X*v.X + v.Y*v.Y
}

// Dist returns the distance between two vectors
func (v Vec2i) Dist(other Vec2i) float64 {
	return math.Sqrt(v.SqrDist(other))
}

// SqrDist returns the square distance between us and another vector
func (v Vec2i) SqrDist(other Vec2i) float64 {
	return math.Pow(float64(v.X-other.X), 2) + math.Pow(float64(v.Y-other.Y), 2)
}

// XY returns the X and Y components separately
func (v Vec2i) XY() (int, int) {
	return v.X, v.Y
}

// Map applies a function to both X and Y components and
// returns a new Vec2i of the result
func (v Vec2i) Map(fun func(int) int) Vec2i {
	return Vec2i{
		X: fun(v.X),
		Y: fun(v.Y),
	}
}

// ToVec2f converts to floats
func (v Vec2i) ToVec2f() Vec2f {
	return Vec2f{X: float64(v.X), Y: float64(v.Y)}
}
