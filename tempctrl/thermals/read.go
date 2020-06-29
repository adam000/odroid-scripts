package thermals

import (
	"io/ioutil"
	"log"
	"strconv"
)

func Read(temperatureFiles []string) []int {
	temps := make([]int, len(temperatureFiles))

	for i, temperatureFile := range temperatureFiles {
		// Read file
		bytes, err := ioutil.ReadFile(temperatureFile)
		if err != nil {
			// TODO
			panic(err)
		}
		// Parse int
		valueStr := string(bytes)
		if valueStr[len(valueStr)-1] != '\n' {
			// TODO ?!
			continue
		}
		// Chop off trailing newline
		valueStr = valueStr[:len(valueStr)-1]
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			log.Fatalf("Value could not be converted to int: %s", valueStr)
		}

		// Store int in temps
		temps[i] = value
	}

	return temps
}
