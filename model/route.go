package model

import (
	"strconv"
	"strings"
)

type Route struct {
	loads    []Load
	distance float64
}

func NewRoute() Route {
	return Route{
		loads:    []Load{},
		distance: 0,
	}
}

// AppendLoad will append a load to the route.
func (r *Route) AppendLoad(load Load) {
	if len(r.loads) == 0 {
		origin := Point{X: 0.0, Y: 0.0}
		r.distance = load.DistanceToComplete(origin)
	} else {
		lastLoad := r.loads[len(r.loads)-1]
		r.distance += load.DistanceToComplete(lastLoad.Dropoff)
	}
	r.loads = append(r.loads, load)
}

// LoadList returns a properly formatted string of the load numbers
// in the form of "[x,y,z]"
func (r *Route) LoadList() string {
	var loadList strings.Builder
	loadList.WriteString("[")
	seperator := ""
	for _, load := range r.loads {
		loadList.WriteString(seperator + strconv.Itoa(load.Number))
		seperator = ","
	}
	loadList.WriteString("]")

	return loadList.String()
}

// DistanceWithLoad will return the total distance of the route
// if the supplied load were to be added to the end of it.
// It includes driving from the origin, doing all the loads in the route,
// then the additional load, then driving back to the origin.
func (r *Route) DistanceWithLoad(load Load) float64 {
	origin := Point{X: 0.0, Y: 0.0}

	if len(r.loads) == 0 {
		return load.DistanceToComplete(origin) + load.Dropoff.Distance(origin)
	}

	lastLoad := r.loads[len(r.loads)-1]
	return r.distance + load.DistanceToComplete(lastLoad.Dropoff) + origin.Distance(load.Dropoff)
}
