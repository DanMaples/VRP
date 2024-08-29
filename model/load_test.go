package model_test

import (
	"testing"

	"github.com/DanMaples/VRP/model"
)

func TestDistanceToComplete(t *testing.T) {
	load := model.Load{1, model.Point{X: 3.0, Y: 4.0}, model.Point{X: 6.0, Y: 8.0}}

	tests := map[string]struct {
		currentPoint     model.Point
		expectedDistance float64
	}{
		"fromOrigin": {
			currentPoint:     model.Point{X: 0.0, Y: 0.0},
			expectedDistance: 10.0,
		},
		"fromOneAway": {
			currentPoint:     model.Point{X: 3.0, Y: 5.0},
			expectedDistance: 6.0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualDistance := load.DistanceToComplete(tc.currentPoint)
			if tc.expectedDistance != actualDistance {
				t.Fatalf("Expected %.20f, Actual %.20f", tc.expectedDistance, actualDistance)
			}
		})
	}
}
