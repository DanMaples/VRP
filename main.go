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

	routes := DanMaplesVRP(loads)

	for _, route := range routes {
		fmt.Println(route.LoadList())
	}
}

// DanMaplesVRP sends out 1 driver to the closet pickup location and assigns that load to them.
// The algorithm then finds the closest pickup location from the load's dropoff point and checks to
// see if the driver is capable of handling that load.
// If so, it is assigned to that driver.
// If not, the next closest load is checked.
// This repeats until a driver can't take the any more loads.
// At that point, the driver is sent back
// to the depot and a new driver is dispatched.
func DanMaplesVRP(loads map[int]model.Load) []model.Route {
	routes := []model.Route{model.NewRoute()}

	currentLocation := model.Point{X: 0.0, Y: 0.0}

	for len(loads) > 0 {
		closestLoadNumbers := currentLocation.FindClosestLoads(loads)

		closestViableLoadNumber := 0
		for _, loadNum := range closestLoadNumbers {
			if routes[len(routes)-1].DistanceWithLoad(loads[loadNum]) <= maxDriverDistance {
				closestViableLoadNumber = loadNum
				break
			}
		}
		if closestViableLoadNumber == 0 {
			routes = append(routes, model.NewRoute())
			currentLocation = model.Point{X: 0.0, Y: 0.0}
		} else {
			routes[len(routes)-1].AppendLoad(loads[closestViableLoadNumber])
			currentLocation = loads[closestViableLoadNumber].Dropoff
			delete(loads, closestViableLoadNumber)
		}
	}
	return routes
}
