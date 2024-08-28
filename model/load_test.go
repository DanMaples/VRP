package model_test

import (
	"testing"

	"github.com/DanMaples/VRP/model"
)

// @Todo: do a proper float comparison, may not be needed if cost is not float
// @Todo: convert to table driven test
func TestNewLoad(t *testing.T) {
	pickup := model.Point{X: 1.0, Y: 2.0}
	dropoff := model.Point{X: 3.0, Y: 4.0}
	actual := model.NewLoad(1, pickup, dropoff)
	expected := model.Load{
		Number:  1,
		Pickup:  model.Point{X: 1.0, Y: 2.0},
		Dropoff: model.Point{X: 3.0, Y: 4.0},
	}

	if expected.Number != actual.Number ||
		expected.Pickup.X != actual.Pickup.X ||
		expected.Pickup.Y != actual.Pickup.Y ||
		expected.Dropoff.X != actual.Dropoff.X ||
		expected.Dropoff.Y != actual.Dropoff.Y {
		t.Errorf("Expected %+v, Actual %+v", expected, actual)
	}
}

func TestDistanceToComplete(t *testing.T) {
	start := model.Point{X: 3.0, Y: 4.0}
	stop := model.Point{X: 6.0, Y: 8.0}
	a := model.NewLoad(1, start, stop)

	expectedDistance := 10.0

	actualDistance := a.DistanceToComplete(model.NewPoint("0.0,0.0"))

	if expectedDistance != actualDistance {
		t.Errorf("Expected %.20f, Actual %.20f", expectedDistance, actualDistance)
	}
}
