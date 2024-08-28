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

	routes := nextClosestAlgorithm(loads)

	for _, route := range routes {
		fmt.Println(route.LoadList())
	}
}

// nextClosestAlgorithm sends out 1 driver to the closet pickup location and assigns that load to them.
// The algorithm then finds the next closest pickup location to from the dropoff point and checks to
// see if the driver is capable of handling that load. If so, it is assigned to that driver.
// If the driver can't take that load, the driver is sent home and a new driver is dispatched,
// starting the process over.
func nextClosestAlgorithm(loads map[int]model.Load) []model.Route {
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
		} else {
			routes[currentDriverNumber].AppendLoad(loads[closestLoadNumber])
			currentLocation = loads[closestLoadNumber].Dropoff
			delete(loads, closestLoadNumber)
		}
	}
	return routes
}

// func dualDriverAlgorithm(loads map[int]model.Load) []model.Route {
// 	routes := []model.Route{model.NewRoute()}

// 	return routes
// }
