package model

import (
	"fmt"
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

// FindClosetLoad will find the load with the closet pickup point to the current point.
func (p *Point) FindClosestLoad(loads map[int]Load) (int, error) {
	if len(loads) == 0 {
		return 0, fmt.Errorf("input has no loads")
	}

	closestLoadNumber := 0
	closestDistance := 0.0
	for _, load := range loads {
		distanceToLoad := p.Distance(load.Pickup)
		if distanceToLoad < closestDistance || closestLoadNumber == 0 {
			closestDistance = distanceToLoad
			closestLoadNumber = load.Number
		}
	}

	return closestLoadNumber, nil
}
