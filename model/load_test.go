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
	l := model.NewLoad(start, stop)
	expectedCost := 5.0

	if l.Cost() != expectedCost {
		t.Errorf("Expected %.20f, Actual %.20f", expectedCost, l.Cost())
	}
}
