package mathf_test

import (
	"math"
	"testing"

	"github.com/miniscruff/igloo/mathf"
)

func TestRotateTowards(t *testing.T) {
	tests := map[string]struct {
		current  float64
		target   float64
		maxDelta float64
		expected float64
	}{
		"greater": {
			current:  0,
			target:   0.15,
			maxDelta: 0.3,
			expected: 0.15,
		},
		"less than": {
			current:  0.25,
			target:   0.15,
			maxDelta: 1,
			expected: 0.15,
		},
		"greater and more then delta": {
			current:  0.1,
			target:   0.5,
			maxDelta: 0.15,
			expected: 0.25,
		},
		"less and more then delta": {
			current:  0.5,
			target:   0.1,
			maxDelta: 0.2,
			expected: 0.3,
		},
		"pi to -pi": {
			current: math.Pi-0.1,
			target: -math.Pi+0.1,
			maxDelta: 0.3,
			expected: -math.Pi+0.1,
		},
		"-pi to pi": {
			current: -math.Pi+0.1,
			target: math.Pi-0.1,
			maxDelta: 0.3,
			expected: math.Pi-0.1,
		},
		"pi to -pi hit max": {
			current: math.Pi-0.2,
			target: -math.Pi+0.2,
			maxDelta: 0.3,
			expected: -math.Pi+0.1,
		},
		"-pi to pi hit max": {
			current: -math.Pi+0.2,
			target: math.Pi-0.2,
			maxDelta: 0.3,
			expected: math.Pi-0.1,
		},
		"pi to -pi stop before bounce": {
			current: math.Pi-0.3,
			target: -math.Pi+0.3,
			maxDelta: 0.2,
			expected: math.Pi-0.1,
		},
		"-pi to pi stop before bounce": {
			current: -math.Pi+0.3,
			target: math.Pi-0.3,
			maxDelta: 0.15,
			expected: -math.Pi+0.15,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := mathf.RotateTowards(tc.current, tc.target, tc.maxDelta)
			if math.Abs(tc.expected - got) > 0.005 {
				t.Fatalf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
