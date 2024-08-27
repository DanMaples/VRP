package model_test

import (
	"testing"

	"github.com/DanMaples/VRP/model"
)

func TestRouteCost(t *testing.T) {
	start := model.Point{X: 3.0, Y: 4.0}
	stop := model.Point{X: 9.0, Y: 12.0}
	a := model.NewLoad(1, start, stop)

	start = model.Point{X: 9.0, Y: -12.0}
	stop = model.Point{X: 3.0, Y: -4.0}
	b := model.NewLoad(2, start, stop)

	route := model.NewRoute()
	route.AppendLoad(a)
	route.AppendLoad(b)

	actualCost := route.Cost()
	expectedCost := 54.0

	if expectedCost != actualCost {
		t.Errorf("Expected %.20f, Actual %.20f", expectedCost, actualCost)
	}
}
