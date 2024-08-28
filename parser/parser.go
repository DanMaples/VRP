package parser

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/DanMaples/VRP/model"
)

// Parse will parse an input file and return a map of loads
// where the key is the loadNumber of the load.
func Parse(f string) map[int]model.Load {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ' '

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return parseData(data)
}

// ParseData will parse the raw data read from the file into a map
// of loads where the key is the loadNumber of the load.
func parseData(data [][]string) map[int]model.Load {
	loads := make(map[int]model.Load, len(data)-1)
	for row := 1; row < len(data); row++ {
		loadNumber, err := strconv.Atoi(data[row][0])
		if err != nil {
			panic(err)
		}
		pickup := model.NewPoint(data[row][1])
		dropoff := model.NewPoint(data[row][2])
		loads[loadNumber] = model.NewLoad(loadNumber, pickup, dropoff)
	}
	return loads
}
