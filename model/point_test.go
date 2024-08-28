package model_test

import (
	"testing"

	"github.com/DanMaples/VRP/model"
)

func TestNewPoint(t *testing.T) {
	inputString := "(-9.100071078494038,-48.89301103772511)"
	expected := model.Point{X: -9.100071078494038, Y: -48.89301103772511}

	actual := model.NewPoint(inputString)
	if actual.X != expected.X || actual.Y != expected.Y {
		t.Errorf("Expected (%.20f,%.20f), Actual (%.20f,%.20f)", expected.X, expected.Y, actual.X, actual.Y)
	}
}

func TestDistance(t *testing.T) {
	a := model.Point{X: 0.0, Y: 0.0}
	b := model.Point{X: 3.0, Y: 4.0}

	expectedDistance := 5.0
	actualDistance := a.Distance(b)

	if expectedDistance != actualDistance {
		t.Errorf("Expected %.20f, Actual %.20f", expectedDistance, actualDistance)
	}
}

func TestFindClosestLoad(t *testing.T) {
	currentPoint := model.Point{X: 9.1, Y: 10.1}
	var loads map[int]model.Load

	// check for nil map
	_, err := currentPoint.FindClosestLoad(loads)
	if err == nil {
		t.Errorf("expected error, but none found")
	}

	// check for empty map
	loads = map[int]model.Load{}
	_, err = currentPoint.FindClosestLoad(loads)
	if err == nil {
		t.Errorf("expected error, but none found")
	}

	loads = map[int]model.Load{
		1: {
			Number:  1,
			Pickup:  model.Point{X: 3.0, Y: 4.0},
			Dropoff: model.Point{X: 5.0, Y: 6.0},
		},
		2: {
			Number:  2,
			Pickup:  model.Point{X: 7.0, Y: 8.0},
			Dropoff: model.Point{X: 9.0, Y: 10.0},
		},
		3: {
			Number:  3,
			Pickup:  model.Point{X: 9.0, Y: 10.0},
			Dropoff: model.Point{X: 9.0, Y: 10.0},
		},
		4: {
			Number:  4,
			Pickup:  model.Point{X: 11.0, Y: 12.0},
			Dropoff: model.Point{X: 9.0, Y: 10.0},
		},
	}

	expected := 3
	actual, err := currentPoint.FindClosestLoad(loads)

	if err != nil {
		t.Errorf("unexpected error returned")
	}

	if expected != actual {
		t.Errorf("Expected %d Actual %d", expected, actual)
	}
}
