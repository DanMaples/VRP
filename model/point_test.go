package model_test

import (
	"reflect"
	"testing"

	"github.com/DanMaples/VRP/model"
)

func TestNewPoint(t *testing.T) {
	tests := map[string]struct {
		input         string
		expectedPoint model.Point
	}{
		"basic": {
			input:         "(-9.100071078494038,-48.89301103772511)",
			expectedPoint: model.Point{X: -9.100071078494038, Y: -48.89301103772511},
		},
		"noParens": {
			input:         "-19.100071078494038,-4.89301103772511",
			expectedPoint: model.Point{X: -19.100071078494038, Y: -4.89301103772511},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualPoint := model.NewPoint(tc.input)
			if actualPoint.X != tc.expectedPoint.X || actualPoint.Y != tc.expectedPoint.Y {
				t.Fatalf("Expected (%.20f,%.20f), Actual (%.20f,%.20f)", tc.expectedPoint.X, tc.expectedPoint.Y, actualPoint.X, actualPoint.Y)
			}
		})
	}
}

func TestDistance(t *testing.T) {
	tests := map[string]struct {
		pointA           model.Point
		pointB           model.Point
		expectedDistance float64
	}{
		"basic": {
			pointA:           model.Point{X: 0.0, Y: 0.0},
			pointB:           model.Point{X: 3.0, Y: 4.0},
			expectedDistance: 5.0,
		},
		"offset": {
			pointA:           model.Point{X: 1.0, Y: 1.0},
			pointB:           model.Point{X: 4.0, Y: 5.0},
			expectedDistance: 5.0,
		},
		"negativeQuadrant": {
			pointA:           model.Point{X: -1.0, Y: -1.0},
			pointB:           model.Point{X: -4.0, Y: -5.0},
			expectedDistance: 5.0,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualDistance := tc.pointA.Distance(tc.pointB)
			if tc.expectedDistance != actualDistance {
				t.Fatalf("Expected %.20f, Actual %.20f", tc.expectedDistance, actualDistance)
			}
		})
	}
}

func TestFindClosestLoads(t *testing.T) {
	loadOne := model.Load{Number: 1, Pickup: model.Point{X: 5.0, Y: -4.0}, Dropoff: model.Point{X: 8.0, Y: -4.0}}
	loadTwo := model.Load{Number: 2, Pickup: model.Point{X: 5.0, Y: 4.0}, Dropoff: model.Point{X: 7.0, Y: 4.0}}
	loadThree := model.Load{Number: 3, Pickup: model.Point{X: -2.0, Y: -8.0}, Dropoff: model.Point{X: -6.0, Y: -8.0}}
	loadFour := model.Load{Number: 4, Pickup: model.Point{X: 5.0, Y: -2.0}, Dropoff: model.Point{X: 8.0, Y: -2.0}}
	loads := map[int]model.Load{
		1: loadOne,
		2: loadTwo,
		3: loadThree,
		4: loadFour,
	}

	tests := map[string]struct {
		currentPoint  model.Point
		expectedOrder []int
	}{
		"2,3": {
			currentPoint:  model.Point{X: 2.0, Y: 3.0},
			expectedOrder: []int{2, 4, 1, 3},
		},
		"-5,-10": {
			currentPoint:  model.Point{X: -5.0, Y: -10.0},
			expectedOrder: []int{3, 1, 4, 2},
		},
		"4,-2": {
			currentPoint:  model.Point{X: 4.0, Y: -2.0},
			expectedOrder: []int{4, 1, 2, 3},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actualOrder := tc.currentPoint.FindClosestLoads(loads)
			if reflect.DeepEqual(tc.expectedOrder, actualOrder) != true {
				t.Errorf("Expected %v Actual %v", tc.expectedOrder, actualOrder)
			}
		})
	}
}
