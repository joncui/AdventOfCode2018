package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	frequencyChanges := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
	var intFrequencyChanges []int
	frequency := 0

	for _, frequencyChange := range frequencyChanges {
		intFreqChange, err := strconv.Atoi(frequencyChange)
		check(err)

		intFrequencyChanges = append(intFrequencyChanges, intFreqChange)

		frequency += intFreqChange
	}

	fmt.Printf("The resulting frequency is %d.\n", frequency)
	fmt.Printf("The first frequency the device reaches twice is: %d.\n", getPart2Solution(intFrequencyChanges))
}

func getPart2Solution(intFrequencyChanges []int) int {

	frequency := 0
	frequencyMap := make(map[int]bool)
	frequencyMap[frequency] = false

	for i := 0; true; i++ {
		frequency += intFrequencyChanges[i]

		value, exists := frequencyMap[frequency]
		if !exists {
			frequencyMap[frequency] = false
		} else if exists && !value {
			frequencyMap[frequency] = true
			break
		}

		if len(intFrequencyChanges) == (i + 1) {
			i = -1
		}
	}

	return frequency
}
