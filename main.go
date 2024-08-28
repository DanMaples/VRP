package main

import (
	"fmt"

	"github.com/DanMaples/VRP/model"
	"github.com/DanMaples/VRP/parser"
)

func main() {
	loads := parser.Parse("problem1.txt")
	route := model.NewRoute()

	currentLocation := model.Point{X: 0.0, Y: 0.0}

	for len(loads) > 0 {
		closestLoadNumber, err := currentLocation.FindClosestLoad(loads)
		if err != nil {
			panic(err)
		}
		route.AppendLoad(loads[closestLoadNumber])
		currentLocation = loads[closestLoadNumber].Dropoff
		delete(loads, closestLoadNumber)
	}

	fmt.Println(route.LoadList())
	fmt.Printf("\nDistance:%f\n", route.Distance())
}
