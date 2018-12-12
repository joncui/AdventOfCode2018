package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parseData() []string {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	return strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
}

func main() {
	lines := parseData()
	currState := strings.Repeat(".", 100) + string(lines[0][15:]) + strings.Repeat(".", 110)
	stateLength := len(currState)

	combinations := parseCombinations(lines[2:])

	for i := 1; i <= 100; i++ {
		var b strings.Builder
		b.WriteString("..")
		for j := 2; j < stateLength-2; j++ {
			if combinations[getStringSlice(&currState, j)] {
				b.WriteString("#")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("..")

		currState = b.String()

		if i == 20 {
			fmt.Printf("After 20 generations, the sum of the numbers of all pots which contain a plant is %d.\n", sumState(&currState))
		}
	}

	fmt.Printf("After 50 billion generations, the sum of the numbers of all pots which contain a plant is %d.\n", (50000000000-100)*80+sumState(&currState))
}

func parseCombinations(lines []string) (combinations map[string]bool) {
	combinations = make(map[string]bool)
	for _, line := range lines {
		splitLine := strings.Split(line, " => ")
		combinations[splitLine[0]] = splitLine[1] == "#"
	}

	return
}

func getStringSlice(state *string, index int) string {
	return (*state)[index-2 : index+3]
}

func sumState(state *string) (sum int) {
	stateArr := strings.Split(strings.TrimRight(*state, "."), "")
	for i := 0; i < len(stateArr); i++ {
		if stateArr[i] == "#" {
			sum += i - 100
		}
	}

	return
}
