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

func parseFile() []string {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	return strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
}

func main() {
	gameMap := buildMap(parseFile())
	fmt.Println(gameMap)
	// elves, goblins := findCreatures(&gameMap)
}

func buildMap(lines []string) (gameMap [][]string) {
	gameMap = [][]string{}
	for _, line := range lines {
		gameMap = append(gameMap, strings.Split(line, ""))
	}

	return
}

func findCreatures(gameMap *[][]string) (elves []Creature, goblins []Creature) {
	for y := 0; y < len(*gameMap); y++ {
		for x := 0; x < len((*gameMap)[y]); x++ {
			if (*gameMap)[y][x] == "E" {
				elves = append(elves, InitElf(x, y))
				(*gameMap)[y][x] = "."
			} else if (*gameMap)[y][x] == "G" {
				goblins = append(goblins, InitGoblin(x, y))
				(*gameMap)[y][x] = "."
			}
		}
	}
	return
}
