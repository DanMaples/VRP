package model

import (
	"strconv"
	"strings"
)

type Route struct {
	loads []Load
	cost  float64
}

func NewRoute() Route {
	return Route{
		loads: []Load{},
		cost:  0,
	}
}

// AppendLoad will append a load to the route.
func (r *Route) AppendLoad(load Load) {
	if len(r.loads) == 0 {
		origin := Point{X: 0.0, Y: 0.0}
		r.cost = load.DistanceToComplete(origin)
	} else {
		lastLoad := r.loads[len(r.loads)-1]
		r.cost += load.DistanceToComplete(lastLoad.Dropoff)
	}
	r.loads = append(r.loads, load)
}

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

// Distance returns the total distance from the origin,
// through all stops, and back to the origin.
func (r *Route) Distance() float64 {
	if len(r.loads) == 0 {
		return 0
	}

	origin := Point{X: 0.0, Y: 0.0}
	lastLoad := r.loads[len(r.loads)-1]
	return r.cost + origin.Distance(lastLoad.Dropoff)
}
