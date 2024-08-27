package model_test

import (
	"testing"

	"github.com/DanMaples/VRP/model"
)

// @Todo: do a proper float comparison, may not be needed if cost is not float
// @Todo: convert to table driven test
func TestNewLoad(t *testing.T) {
	start := model.Point{X: 0.0, Y: 0.0}
	stop := model.Point{X: 3.0, Y: 4.0}
	l := model.NewLoad(1, start, stop)

	if l.Pickup.X != start.X || l.Pickup.Y != start.Y {
		t.Errorf("Pickup incorrect: Expected %+v, Actual %+v", start, l.Pickup)
	}

	if l.Dropoff.X != stop.X || l.Dropoff.Y != stop.Y {
		t.Errorf("Dropoff incorrect: Expected %+v, Actual %+v", start, l.Pickup)
	}
}

func TestLoadCost(t *testing.T) {
	start := model.Point{X: 3.0, Y: 4.0}
	stop := model.Point{X: 6.0, Y: 8.0}
	a := model.NewLoad(1, start, stop)

	expectedCost := 10.0

	actualCost := a.Cost(model.NewPoint("0.0,0.0"))

	if expectedCost != actualCost {
		t.Errorf("Expected %.20f, Actual %.20f", expectedCost, actualCost)
	}
}
