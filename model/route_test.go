package model_test

import (
	"strings"
	"testing"

	"github.com/DanMaples/VRP/model"
)

func TestRouteDistanceWithLoad(t *testing.T) {
	pickup := model.Point{X: -3.0, Y: 4.0}
	dropoff := model.Point{X: -6.0, Y: 0.0}
	loadA := model.NewLoad(1, pickup, dropoff)

	pickup = model.Point{X: -6.0, Y: 4.0}
	dropoff = model.Point{X: -9.0, Y: 0.0}
	loadB := model.NewLoad(2, pickup, dropoff)

	tests := map[string]struct {
		routeLoads       []model.Load
		newLoad          model.Load
		expectedDistance float64
	}{
		"noLoads": {
			routeLoads:       []model.Load{},
			newLoad:          loadA,
			expectedDistance: 16.0,
		},
		"hasLoads": {
			routeLoads:       []model.Load{loadA},
			newLoad:          loadB,
			expectedDistance: 28.0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			route := model.NewRoute()
			for _, load := range tc.routeLoads {
				route.AppendLoad(load)
			}
			actualDistance := route.DistanceWithLoad(tc.newLoad)
			if tc.expectedDistance != actualDistance {
				t.Errorf("Expected %.20f, Actual %.20f", tc.expectedDistance, actualDistance)
			}
		})
	}
}

func TestLoadList(t *testing.T) {
	loadOne := model.NewLoad(1, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0})
	loadTwo := model.NewLoad(2, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0})
	loadThree := model.NewLoad(3, model.Point{X: -3.0, Y: 4.0}, model.Point{X: -6.0, Y: 0.0})

	tests := map[string]struct {
		loads        []model.Load
		expectedList string
	}{
		"oneLoad": {
			loads:        []model.Load{loadOne},
			expectedList: "[1]",
		},
		"twoLoads": {
			loads:        []model.Load{loadTwo, loadOne},
			expectedList: "[2,1]",
		},
		"threeLoads": {
			loads:        []model.Load{loadTwo, loadThree, loadOne},
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
