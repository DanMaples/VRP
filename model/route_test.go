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
	tests := map[string]struct {
		loads        []model.Load
		expectedList string
	}{
		"oneLoad": {
			loads: []model.Load{
				model.NewLoad(1, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0}),
			},
			expectedList: "[1]",
		},
		"twoLoads": {
			loads: []model.Load{
				model.NewLoad(2, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0}),
				model.NewLoad(1, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0}),
			},
			expectedList: "[2,1]",
		},
		"threeLoads": {
			loads: []model.Load{
				model.NewLoad(2, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0}),
				model.NewLoad(3, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0}),
				model.NewLoad(1, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0}),
			},
			expectedList: "[2,3,1]",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			route := model.NewRoute()
			for _, load := range tc.loads {
				route.AppendLoad(load)
			}
			actual := route.LoadList()
			if strings.Compare(tc.expectedList, actual) != 0 {
				t.Fatalf("Expected %s, Actual %s", tc.expectedList, actual)
			}
		})
	}
}
