package model

type Load struct {
	Number  int
	Pickup  Point
	Dropoff Point
}

func NewLoad(number int, pickup Point, dropoff Point) Load {
	return Load{
		Number:  number,
		Pickup:  pickup,
		Dropoff: dropoff,
	}
}

// DistanceToComplete reprents the distance required to complete the load
// from the given point.
func (l *Load) DistanceToComplete(p Point) float64 {
	return p.Distance(l.Pickup) + l.Pickup.Distance(l.Dropoff)
}
