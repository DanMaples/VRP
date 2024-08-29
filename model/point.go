package model

import (
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

// FindClosestLoad will return an ordered slice of the closest loads.
func (p *Point) FindClosestLoads(loads map[int]Load) []int {
	numberLoads := len(loads)

	type loadDistance struct {
		loadNumber int
		distance   float64
	}

	loadDistances := make([]loadDistance, numberLoads)
	index := 0
	for _, load := range loads {
		loadDistances[index] = loadDistance{loadNumber: load.Number, distance: p.Distance(load.Pickup)}
		index++
	}

	sort.Slice(loadDistances, func(i, j int) bool {
		return loadDistances[i].distance < loadDistances[j].distance
	})

	orderedLoads := make([]int, numberLoads)
	for index, loadDistance := range loadDistances {
		orderedLoads[index] = loadDistance.loadNumber
	}

	return orderedLoads
}
