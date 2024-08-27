package main

import (
	"fmt"

	"github.com/DanMaples/VRP/model"
	"github.com/DanMaples/VRP/parser"
)

func main() {
	start := model.Point{X: 2.0, Y: 2.0}
	stop := model.Point{X: 5.0, Y: 6.0}
	l := model.NewLoad(start, stop)

	fmt.Printf("l.Cost()=%f\n", l.Cost())

	loads := parser.Parse("problem1.txt")

	for loadNumber, l := range loads {
		fmt.Printf("Num:%d, load:%+v\n", loadNumber, l)
	}
}
