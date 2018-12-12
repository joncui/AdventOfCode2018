package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var input = 1309

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func strToInt(str string) int {
	value, err := strconv.Atoi(str)
	check(err)

	return value
}

func main() {
	if len(os.Args) > 1 {
		input = strToInt(os.Args[1])
	}

	grid := initGrid()
	calculatePowerLevel(&grid)
	x, y, _ := maxSquarePower(&grid, 3)

	fmt.Printf("The %d,%d coordinate has the largest power.\n", x, y)

	maxPower := math.MinInt32
	maxX := 0
	maxY := 0
	maxSize := 0
	for size := 1; size <= 300; size++ {
		x1, y1, power := maxSquarePower(&grid, size)

		if power > maxPower {
			maxPower = power
			maxSize = size
			maxX = x1
			maxY = y1
		}
	}

	fmt.Printf("The identifier of the square with the largest total power is: %d,%d,%d.\n", maxX, maxY, maxSize)
}

func initGrid() [][]int {
	grid := make([][]int, 300)
	for i := 0; i < 300; i++ {
		grid[i] = make([]int, 300)
	}

	return grid
}

func printGrid(grid [][]int) {
	fmt.Println(strings.Repeat("-", 1500))
	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			fmt.Printf("|%4d", grid[y][x])
		}

		fmt.Print("|\n")
		fmt.Println(strings.Repeat("-", 1500))
	}
}

func calculatePowerLevel(grid *[][]int) {
	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			(*grid)[y][x] = getPowerLevel(x+1, y+1)
		}
	}
}

func maxSquarePower(grid *[][]int, size int) (maxX, maxY, maxPowerLevel int) {
	maxPowerLevel = math.MinInt32
	for y := 0; y < 300-size; y++ {
		for x := 0; x < 300-size; x++ {
			power := 0
			for g := y; g < y+size; g++ {
				power += sumArr((*grid)[g][x : x+size])
			}

			if maxPowerLevel < power {
				maxPowerLevel = power
				maxX = x
				maxY = y
			}
		}
	}

	maxX++
	maxY++

	return
}

func getPowerLevel(x, y int) (powerLevel int) {
	rackId := x + 10
	powerLevel = (((((rackId * y) + input) * rackId) / 100) % 10) - 5

	return
}

func sumArr(arr []int) (sum int) {
	for _, e := range arr {
		sum += e
	}

	return
}
