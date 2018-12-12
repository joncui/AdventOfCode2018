package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Star struct {
	posX, posY, vX, vY int
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parseData() (stars []Star) {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	lines := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")

	for _, line := range lines {
		star := Star{
			strToInt(line[10:16]),
			strToInt(line[18:24]),
			strToInt(line[36:38]),
			strToInt(line[40:42]),
		}

		stars = append(stars, star)
	}

	return
}

func strToInt(str string) int {
	num, err := strconv.Atoi(strings.TrimSpace(str))
	check(err)

	return num
}

func main() {
	part1(parseData())
}

func part1(stars []Star) {
	smallestArea := math.MaxUint32
	smallestI := 0
	for i := 0; i < 100000; i++ {
		minX, maxX, minY, maxY := 0, 0, 0, 0

		for _, star := range stars {
			x := star.posX + (star.vX * i)
			y := star.posY + (star.vY * i)

			if maxX < x {
				maxX = x
			} else if minX > x {
				minX = x
			}

			if maxY < y {
				maxY = y
			} else if minY > y {
				minY = y
			}
		}

		lenX := maxX - minX + 1
		lenY := maxY - minY + 1
		area := lenX + lenY

		if smallestArea > area {
			smallestArea = area
			smallestI = i
		}
	}

	minX, maxX, minY, maxY := 0, 0, 0, 0
	for j := 0; j < len(stars); j++ {
		x := stars[j].posX + (stars[j].vX * smallestI)
		y := stars[j].posY + (stars[j].vY * smallestI)

		stars[j].posX = x
		stars[j].posY = y

		if maxX < x {
			maxX = x
		} else if minX > x {
			minX = x
		}

		if maxY < y {
			maxY = y
		} else if minY > y {
			minY = y
		}
	}

	mapper := initMapper(minX, maxX, minY, maxY)
	for _, star := range stars {
		mapper[star.posY][star.posX] = true
	}

	printMap(mapper)
}

func initMapper(minX, maxX, minY, maxY int) [][]bool {
	mapper := make([][]bool, maxY-minY+1)
	for i := 0; i < len(mapper); i++ {
		mapper[i] = make([]bool, maxX-minX+1)
	}

	return mapper
}

func printMap(mapper [][]bool) {
	for y := 0; y < len(mapper); y++ {
		for x := 0; x < len(mapper[0]); x++ {
			if mapper[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
	fmt.Println()
}
