package mathf

import (
	"math"
)

var twoPi = math.Pi * 2

// RotateTowards calculates the angle going from current to target
// not exceeding max delta and respecting the jump from +pi to -pi.
// All angles are in radians.
func RotateTowards(current, target, maxDelta float64) float64 {
	deltaAngle := DeltaAngle(current, target)
	if deltaAngle > -maxDelta && deltaAngle < maxDelta {
		return target
	}

	target = current + deltaAngle
	return BindAngle(MoveTowards(current, target, maxDelta))
}

func MoveTowards(current, target, maxDelta float64) float64 {
	if math.Abs(target-current) <= maxDelta {
		return target
	}

	return current + Sign(target-current)*maxDelta
}

func Sign(value float64) float64 {
	if value > 0 {
		return 1
	}

	return -1
}

func DeltaAngle(current, target float64) float64 {
	return BindAngle(target - current)
}

func BindAngle(angle float64) float64 {
	bound := Repeat(angle, twoPi)
	if bound > math.Pi {
		bound -= twoPi
	}
	return bound
}

func Repeat(value, length float64) float64 {
	return Clamp(value-math.Floor(value/length)*length, 0, length)
}

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}

	if value > max {
		return max
	}

	return value
}

func Lerp(start, end, percent float64) float64 {
	return start + (end-start)*percent
}
