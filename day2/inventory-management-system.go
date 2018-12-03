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

func main() {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	boxIds := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")

	fmt.Printf("The checksum is %d\n", checkSum(boxIds))
	fmt.Printf("The letters common between the two correct box IDs are %v\n", findCorrectBoxId(boxIds))
}

func checkSum(boxIds []string) int {
	doubleCount := 0
	tripleCount := 0

	for _, boxId := range boxIds {
		double, triple := hasDoubleOrTriple(boxId)
		if double {
			doubleCount++
		}

		if triple {
			tripleCount++
		}
	}

	return doubleCount * tripleCount
}

func hasDoubleOrTriple(boxId string) (double, triple bool) {
	boxIdRunes := []rune(boxId)
	runeMap := make(map[rune]int)

	for _, rune := range boxIdRunes {
		runeMap[rune]++
	}

	for _, value := range runeMap {
		if value == 2 {
			double = true
		} else if value == 3 {
			triple = true
		}

		if double && triple {
			break
		}
	}

	return
}

func findCorrectBoxId(boxIds []string) (correctBoxId string) {

	for i := 0; i < len(boxIds)-1; i++ {
		for j := i + 1; j < len(boxIds); j++ {
			if numDiff, diffLocation := difference(boxIds[i], boxIds[j]); numDiff == 1 {
				correctBoxId = boxIds[i][:diffLocation] + boxIds[i][diffLocation+1:]
			}
		}
	}

	return
}

func difference(id1, id2 string) (numDiff int, diffLocation int) {
	id1Rune := []rune(id1)
	id2Rune := []rune(id2)

	for i := 0; i < len(id1); i++ {
		if id1Rune[i] != id2Rune[i] {
			numDiff++
			diffLocation = i
		}
	}

	return
}
