package mathf_test

import (
	"testing"

	"github.com/miniscruff/igloo/mathf"
)

func TestBoundsContains(t *testing.T) {
	tests := map[string]struct {
		us mathf.Bounds
		point mathf.Vec2
		expected bool
	}{
		"left": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			point: mathf.Vec2{X: 0, Y: 20},
			expected: false,
		},
		"right": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			point: mathf.Vec2{X: 55, Y: 20},
			expected: false,
		},
		"above": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			point: mathf.Vec2{X: 20, Y: 0},
			expected: false,
		},
		"below": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			point: mathf.Vec2{X: 20, Y: 55},
			expected: false,
		},
		"in": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			point: mathf.Vec2{X: 17, Y: 18},
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.us.Contains(tc.point)
			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestBoundsOverlaps(t *testing.T) {
	tests := map[string]struct {
		us mathf.Bounds
		other mathf.Bounds
		expected bool
	}{
		"outside": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			other: mathf.NewBoundsWidthHeight(2, 2, 8, 8),
			expected: false,
		},
		"partially inside": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			other: mathf.NewBoundsWidthHeight(30, 30, 25, 25),
			expected: true,
		},
		"completely inside": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			other: mathf.NewBoundsWidthHeight(12, 15, 6, 6),
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.us.Overlaps(tc.other)
			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestBoundsInside(t *testing.T) {
	tests := map[string]struct {
		us mathf.Bounds
		other mathf.Bounds
		expected bool
	}{
		"outside": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			other: mathf.NewBoundsWidthHeight(2, 2, 8, 8),
			expected: false,
		},
		"partially inside": {
			us: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			other: mathf.NewBoundsWidthHeight(30, 30, 25, 25),
			expected: false,
		},
		"completely inside": {
			us: mathf.NewBoundsWidthHeight(12, 15, 6, 6),
			other: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.us.Inside(tc.other)
			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}

func TestBoundsSurrounds(t *testing.T) {
	tests := map[string]struct {
		us mathf.Bounds
		other mathf.Bounds
		expected bool
	}{
		"outside": {
			us: mathf.NewBoundsWidthHeight(2, 2, 8, 8),
			other: mathf.NewBoundsWidthHeight(15, 15, 25, 25),
			expected: false,
		},
		"partially outside": {
			us: mathf.NewBoundsWidthHeight(30, 30, 25, 25),
			other: mathf.NewBoundsWidthHeight(10, 10, 25, 25),
			expected: false,
		},
		"completely surrounded": {
			us: mathf.NewBoundsWidthHeight(5, 7, 35, 52),
			other: mathf.NewBoundsWidthHeight(12, 15, 6, 6),
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.us.Surrounds(tc.other)
			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
