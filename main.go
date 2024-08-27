package main

import (
	"fmt"

	"github.com/DanMaples/VRP/model"
	"github.com/DanMaples/VRP/parser"
)

func main() {
	orign := model.Point{X: 0.0, Y: 0.0}
	pickup := model.Point{X: 3.0, Y: 4.0}
	dropoff := model.Point{X: 6.0, Y: 8.0}
	l := model.NewLoad(pickup, dropoff)

	fmt.Printf("l.Cost(origin)=%f\n", l.Cost(orign))

	loads := parser.Parse("problem1.txt")

	for loadNumber, l := range loads {
		fmt.Printf("Num:%d, load:%+v\n", loadNumber, l)
	}
}
