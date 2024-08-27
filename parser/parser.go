package parser

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/DanMaples/VRP/model"
)

// TODO: return errs instead of panicking
// TODO: unit test, probably using afero,
//
//	or break out file read.
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

	loads := make(map[int]model.Load, len(data)-1)
	for row := 1; row < len(data); row++ {
		loadNumber, err := strconv.Atoi(data[row][0])
		if err != nil {
			panic(err)
		}
		loads[loadNumber] = model.NewLoad(loadNumber, model.NewPoint(data[row][1]), model.NewPoint(data[row][2]))
	}
	return loads
}
