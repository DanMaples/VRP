package main

import (
	"fmt"
	"os"

	"github.com/DanMaples/VRP/model"
	"github.com/DanMaples/VRP/parser"
)

const (
	maxDriverDistance = 720.0
)

func main() {
	filePath := os.Args[1]

	loads := parser.Parse(filePath)

	routes := enhancedNextClosestAlgorithm(loads)

	for _, route := range routes {
		fmt.Println(route.LoadList())
	}
}

// nextClosestAlgorithm sends out 1 driver to the closet pickup location and assigns that load to them.
// The algorithm then finds the next closest pickup location from the dropoff point and checks to
// see if the driver is capable of handling that load. If so, it is assigned to that driver.
// This repeats until a driver can't take the next closest load. At that point, the driver is sent back
// to the depot and a new driver is dispatched.
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

func enhancedNextClosestAlgorithm(loads map[int]model.Load) []model.Route {
	routes := []model.Route{model.NewRoute()}
	currentDriverNumber := 0

	currentLocation := model.Point{X: 0.0, Y: 0.0}

	for len(loads) > 0 {
		closestLoadNumbers, err := currentLocation.FindClosestLoads(loads)
		if err != nil {
			panic(err)
		}
		closestViableLoadNumber := 0
		for _, loadNum := range closestLoadNumbers {
			if routes[currentDriverNumber].DistanceWithLoad(loads[loadNum]) > maxDriverDistance {
				continue
			} else {
				closestViableLoadNumber = loadNum
				break
			}
		}
		if closestViableLoadNumber == 0 {
			routes = append(routes, model.NewRoute())
			currentDriverNumber++
			currentLocation = model.Point{X: 0.0, Y: 0.0}
		} else {
			routes[currentDriverNumber].AppendLoad(loads[closestViableLoadNumber])
			currentLocation = loads[closestViableLoadNumber].Dropoff
			delete(loads, closestViableLoadNumber)
		}
	}
	return routes
}
