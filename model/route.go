package model

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

func (r *Route) AppendLoad(load Load) {
	if len(r.loads) == 0 {
		origin := Point{X: 0.0, Y: 0.0}
		r.cost = load.Cost(origin)
	} else {
		r.cost += load.Cost(r.loads[len(r.loads)-1].Dropoff)
	}
	r.loads = append(r.loads, load)
}

func (r *Route) GetLoads() []Load {
	return r.loads
}

func (r *Route) Cost() float64 {
	if len(r.loads) == 0 {
		return 0
	}

	origin := Point{X: 0.0, Y: 0.0}
	return r.cost + origin.Distance(r.loads[len(r.loads)-1].Dropoff)
}
