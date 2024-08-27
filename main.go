package main

import (
	"fmt"

	"github.com/DanMaples/VRP/model"
)

func main() {
	start := model.Point{X: 2.0, Y: 2.0}
	stop := model.Point{X: 5.0, Y: 6.0}
	l := model.NewLoad(start, stop)

	fmt.Printf("l.cost=%f\n", l.Cost())
}
