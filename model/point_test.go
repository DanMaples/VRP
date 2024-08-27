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
