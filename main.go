package main

import (
	"fmt"
	"os"

	"github.com/DanMaples/VRP/model"
	"github.com/DanMaples/VRP/parser"
)

const (
	costPerDriver     = 500
	maxDriverDistance = 720.0
)

func main() {
	filePath := os.Args[1]

	loads := parser.Parse(filePath)
	routes := []model.Route{model.NewRoute()}
	currentDriverNumber := 0

	currentLocation := model.Point{X: 0.0, Y: 0.0}

	for len(loads) > 0 {
		closestLoadNumber, err := currentLocation.FindClosestLoad(loads)
		if err != nil {
			panic(err)
		}
		if routes[currentDriverNumber].DistanceWithLoad(loads[closestLoadNumber]) > maxDriverDistance {
			routes = append(routes, model.NewRoute())
			currentDriverNumber++
			currentLocation = model.Point{X: 0.0, Y: 0.0}
			continue
		}
		routes[currentDriverNumber].AppendLoad(loads[closestLoadNumber])
		currentLocation = loads[closestLoadNumber].Dropoff
		delete(loads, closestLoadNumber)
	}

	totalDistance := 0.0
	totalDrivers := 0
	for _, route := range routes {
		fmt.Println(route.LoadList())
		//fmt.Println(route.LoadList(), "Distance ", route.Distance())
		totalDistance += route.Distance()
		totalDrivers += 1
	}

	//fmt.Println("Total Drivers ", totalDrivers)
	//totalCost := float64(costPerDriver*totalDrivers) + totalDistance
	//fmt.Println("Total Cost: ", totalCost)

}
