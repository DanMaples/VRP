package model_test

import (
	"testing"

	"github.com/DanMaples/VRP/model"
)

func TestRouteCost(t *testing.T) {
	pickup := model.Point{X: 3.0, Y: 4.0}
	dropoff := model.Point{X: 9.0, Y: 12.0}
	a := model.NewLoad(1, pickup, dropoff)

	pickup = model.Point{X: 9.0, Y: -12.0}
	dropoff = model.Point{X: 3.0, Y: -4.0}
	b := model.NewLoad(2, pickup, dropoff)

	route := model.NewRoute()
	route.AppendLoad(a)
	route.AppendLoad(b)

	actualDistance := route.Distance()
	expectedDistance := 54.0

	if expectedDistance != actualDistance {
		t.Errorf("Expected %.20f, Actual %.20f", expectedDistance, actualDistance)
	}
}
