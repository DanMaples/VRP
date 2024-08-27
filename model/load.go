package model

type Load struct {
	Pickup  Point
	Dropoff Point
}

func NewLoad(pickup Point, dropoff Point) Load {
	return Load{
		Pickup:  pickup,
		Dropoff: dropoff,
	}
}

func (l *Load) Cost(p Point) float64 {
	return p.Distance(l.Pickup) + l.Pickup.Distance(l.Dropoff)
}
