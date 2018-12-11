package main

import (
	"container/ring"
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

func parseFile() []string {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	return strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
}

func getGameValues(line string) (players, lastMarblePoints int) {
	splitLine := strings.Split(line, " ")

	return strToInt(splitLine[0]), strToInt(splitLine[6])
}

func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	check(err)

	return num
}

func intArrMax(arr []int) (max int) {
	for _, e := range arr {
		if e > max {
			max = e
		}
	}

	return
}

func main() {
	for _, line := range parseFile() {
		players, lastMarblePoints := getGameValues(line)
		fmt.Printf("Players: %d; Last Marble: %d; ", players, lastMarblePoints)
		// fmt.Printf("High score for part 1 is %d\n", part1(players, lastMarblePoints))
		fmt.Printf("High score for part 2 is %d\n", part1(players, lastMarblePoints*100))
	}
}

func part1(players, lastMarblePoints int) int {
	playerPoints := make([]int, players)
	curr := ring.New(1)
	curr.Value = 0

	for marbleNum := 1; marbleNum <= lastMarblePoints; marbleNum++ {
		if marbleNum%23 == 0 {
			playerPoints[marbleNum%players] += marbleNum
			curr = curr.Move(-8)
			removedMarble := curr.Unlink(1).Value.(int)
			playerPoints[marbleNum%players] += removedMarble
			if removedMarble == lastMarblePoints {
				break
			}
			curr = curr.Next()
		} else {
			curr = addRing(curr, marbleNum)
		}
	}

	return intArrMax(playerPoints)
}

/**
 * Adds a new ring after the current ring.
 * Returns the inserted ring.
 */
func addRing(r *ring.Ring, value int) (curr *ring.Ring) {
	newRing := ring.New(1)
	newRing.Value = value

	rs := r.Next().Link(newRing)

	return rs.Prev()
}

func printRing(player int, r *ring.Ring, currValue int) {
	fmt.Printf("[%d]  ", player)
	r.Do(func(p interface{}) {
		if currValue == p.(int) {
			fmt.Printf("(%d) ", currValue)
		} else {
			fmt.Printf("%d ", p.(int))
		}
	})

	fmt.Println()
}
