package parser_test

import (
	"reflect"
	"testing"

	"github.com/DanMaples/VRP/model"
	"github.com/DanMaples/VRP/parser"
)

func TestParseCSV(t *testing.T) {
	inputData := [][]string{
		{"loadNumber", "pickup", "dropoff"},
		{"1", "(3.0,4.0)", "-5.0,-6.0"},
		{"2", "-7.0,-8.0", "(9.0,10.0)"},
	}

	actual := parser.ParseData(inputData)
	expected := map[int]model.Load{
		1: {
			Number:  1,
			Pickup:  model.Point{X: 3.0, Y: 4.0},
			Dropoff: model.Point{X: -5.0, Y: -6.0},
		},
		2: {
			Number:  2,
			Pickup:  model.Point{X: -7.0, Y: -8.0},
			Dropoff: model.Point{X: 9.0, Y: 10.0},
		},
	}

	if reflect.DeepEqual(expected, actual) != true {
		t.Errorf("Expected %+v Actual %+v", expected, actual)
	}
}
