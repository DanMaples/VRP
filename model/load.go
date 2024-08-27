package model

import "math"

type Load struct {
	Pickup  Point
	Dropoff Point
	cost    float64
}

func NewLoad(pickup Point, dropoff Point) Load {
	return Load{
		Pickup:  pickup,
		Dropoff: dropoff,
		cost:    math.Sqrt(math.Pow(dropoff.X-pickup.X, 2) + math.Pow(dropoff.Y-pickup.Y, 2)),
	}
}

func (l *Load) Cost() float64 { return l.cost }
