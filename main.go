package main

import (
	"fmt"

	"github.com/DanMaples/VRP/model"
	"github.com/DanMaples/VRP/parser"
)

func main() {
	loads := parser.Parse("problem1.txt")
	route := model.NewRoute()

	for _, l := range loads {
		fmt.Printf("load:%+v\n", l)
	}

	for i := 10; i > 0; i-- {
		load, exists := loads[i]
		if !exists {
			panic("not in map")
		}
		route.AppendLoad(load)
	}

	fmt.Println(route.Loads())
	fmt.Printf("\ncost:%f\n", route.Distance())
}
