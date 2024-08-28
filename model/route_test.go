package model_test

import (
	"strings"
	"testing"

	"github.com/DanMaples/VRP/model"
)

func TestRouteDistanceWithLoad(t *testing.T) {
	pickup := model.Point{X: -3.0, Y: 4.0}
	dropoff := model.Point{X: -6.0, Y: 0.0}
	a := model.NewLoad(1, pickup, dropoff)

	pickup = model.Point{X: -6.0, Y: 4.0}
	dropoff = model.Point{X: -9.0, Y: 0.0}
	b := model.NewLoad(2, pickup, dropoff)

	route := model.NewRoute()

	actualDistance := route.DistanceWithLoad(a)
	expectedDistance := 16.0

	if expectedDistance != actualDistance {
		t.Errorf("Expected %.20f, Actual %.20f", expectedDistance, actualDistance)
	}

	route.AppendLoad(a)
	actualDistance = route.DistanceWithLoad(b)
	expectedDistance = 28.0

	if expectedDistance != actualDistance {
		t.Errorf("Expected %.20f, Actual %.20f", expectedDistance, actualDistance)
	}
}

func TestLoadList(t *testing.T) {
	pickup := model.Point{X: -3.0, Y: 4.0}
	dropoff := model.Point{X: -6.0, Y: 0.0}
	a := model.NewLoad(1, pickup, dropoff)

	pickup = model.Point{X: -6.0, Y: 4.0}
	dropoff = model.Point{X: -9.0, Y: 0.0}
	b := model.NewLoad(2, pickup, dropoff)

	route := model.NewRoute()
	route.AppendLoad(a)
	route.AppendLoad(b)

	expected := "[1,2]"
	actual := route.LoadList()

	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected %s, Actual %s", expected, actual)
	}

}
