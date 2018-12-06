package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var size = 350
var lessThan = 10000

type coordinateType struct {
	x, y int
}

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

	coordinates := convertToCoordinates(rawData)

	fmt.Printf("The size of the largest area is %d.\n", part1(coordinates))
	fmt.Printf("The size of the region containing all locations which have a total distance to all given coordinates of less than %d is %d.\n", lessThan, part2(coordinates))
}

func part1(coordinates []coordinateType) (maxArea int) {
	theMap, hasInfiniteMap := initializeMaps(coordinates)

	for y := 0; y <= size; y++ {
		for x := 0; x <= size; x++ {
			currentCoord := coordinateType{x, y}
			closestGivenCoordinate := findClosestGivenCoordinate(coordinates, currentCoord)

			if closestGivenCoordinate.x < 0 {
				continue
			}

			if x == 0 || y == 0 || x == size || y == size {
				hasInfiniteMap[closestGivenCoordinate] = true
			}

			theMap[closestGivenCoordinate] = append(theMap[closestGivenCoordinate], currentCoord)
		}
	}

	for _, coordinate := range coordinates {
		if hasInfiniteMap[coordinate] {
			continue
		}

		if area := theMap[coordinate]; len(area) > maxArea {
			maxArea = len(area)
		}
	}

	return
}

func part2(coordinates []coordinateType) int {
	var locations []coordinateType
	for y := 0; y <= size; y++ {
		for x := 0; x <= size; x++ {
			currentCoord := coordinateType{x, y}
			totalDistance := getTotalDistanceFromCoordinates(coordinates, currentCoord)

			if totalDistance < lessThan {
				locations = append(locations, currentCoord)
			}
		}
	}

	return len(locations)
}

func findClosestGivenCoordinate(coordinates []coordinateType, currentCoord coordinateType) coordinateType {
	minManhattanDistance := 1000
	var distanceArr []int
	var minManhattanDistanceCoord coordinateType
	for _, coordinate := range coordinates {
		distance := calculateManhattanDistance(currentCoord, coordinate)
		distanceArr = append(distanceArr, distance)
		if distance < minManhattanDistance {
			minManhattanDistance = distance
			minManhattanDistanceCoord = coordinate
		}
	}

	sort.Ints(distanceArr)
	if distanceArr[0] == distanceArr[1] {
		return coordinateType{-1, -1}
	}

	return minManhattanDistanceCoord
}

func getTotalDistanceFromCoordinates(coordinates []coordinateType, currentCoord coordinateType) (totalDistance int) {
	for _, coordinate := range coordinates {
		distance := calculateManhattanDistance(currentCoord, coordinate)
		totalDistance += distance
	}

	return
}

func convertToCoordinates(rawData []byte) (coordinates []coordinateType) {
	coordinateStrings := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
	for _, coord := range coordinateStrings {
		coordinate := getXYCoordinate(coord)
		coordinates = append(coordinates, coordinate)
	}

	return
}

func initializeMaps(coordinates []coordinateType) (theMap map[coordinateType][]coordinateType, hasInfiniteMap map[coordinateType]bool) {
	theMap = make(map[coordinateType][]coordinateType)
	hasInfiniteMap = make(map[coordinateType]bool)
	for _, coordinate := range coordinates {
		theMap[coordinate] = []coordinateType{}
		hasInfiniteMap[coordinate] = false
	}

	return
}

func getXYCoordinate(coordinateString string) coordinateType {
	coordinates := strings.Split(coordinateString, ", ")

	return coordinateType{strToInt(coordinates[0]), strToInt(coordinates[1])}
}

func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	check(err)

	return num
}

func calculateManhattanDistance(p1, p2 coordinateType) int {
	// fmt.Printf("%v => %v == %d\n", p1, p2, int(math.Abs(float64(p1.x-p2.x))+math.Abs(float64(p1.y-p2.y))))
	return int(math.Abs(float64(p1.x-p2.x)) + math.Abs(float64(p1.y-p2.y)))
}

func printMap(theMap map[coordinateType][]coordinateType) {
	for coordinate, coordinates := range theMap {
		fmt.Printf("%v ==> %v\n", coordinate, coordinates)
	}
}
