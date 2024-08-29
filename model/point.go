package model

import (
	"fmt"
	"math"
	"sort"
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

// Distance returns the distance from this point to the supplied point.
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

// FindClosetLoad will return an ordered slice of the closest loads.
func (p *Point) FindClosestLoads(loads map[int]Load) ([]int, error) {
	if len(loads) == 0 {
		return []int{}, fmt.Errorf("input has no loads")
	}

	type loadDistance struct {
		loadNumber int
		distance   float64
	}

	distances := make([]loadDistance, len(loads))
	index := 0
	for _, load := range loads {
		distances[index] = loadDistance{loadNumber: load.Number, distance: p.Distance(load.Pickup)}
		index++
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	orderedLoads := make([]int, len(loads))
	for index, ld := range distances {
		orderedLoads[index] = ld.loadNumber
	}

	return orderedLoads, nil
}
