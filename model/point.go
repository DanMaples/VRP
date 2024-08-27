package model

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X float64
	Y float64
}

// NewPoint takes a string representation of a point, and parses
// it into a point struct. It accepts strings in the form of:
// x,y and (x,y)
func NewPoint(s string) Point {
	s = strings.Trim(s, "()")
	components := strings.Split(s, ",")

	xComponent, err := strconv.ParseFloat(components[0], 64)
	if err != nil {
		panic(err)
	}

	yComponent, err := strconv.ParseFloat(components[1], 64)
	if err != nil {
		panic(err)
	}

	return Point{
		X: xComponent,
		Y: yComponent,
	}
}

func (p *Point) Distance(d Point) float64 {
	return math.Sqrt(math.Pow(d.X-p.X, 2) + math.Pow(d.Y-p.Y, 2))
}
