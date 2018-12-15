package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parseData() (res [][]string) {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	rows := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
	res = make([][]string, len(rows))

	for i, row := range rows {
		res[i] = strings.Split(row, "")
	}

	return
}

func main() {
	tracks := parseData()
	carts := getMineCarts(&tracks)

	firstCrash := part1(tracks, carts)
	fmt.Printf("The location of the first crash is at %d,%d.\n", firstCrash.x, firstCrash.y)
	lastCart := part2(tracks, carts)
	fmt.Printf("The location of the last cart is at %d,%d.\n", lastCart.x, lastCart.y)
}

func part1(tracks [][]string, carts []Cart) Coord {
	cartMap := make(map[Coord]*Cart)
	cartCopy := append([]Cart{}, carts...)
	for {
		for i := 0; i < len(cartCopy); i++ {
			if crashed := cartCopy[i].UpdateCoordinates(tracks, &cartMap); crashed {
				return cartCopy[i].coordinates
			}
		}

		sort.Sort(ByCoord(cartCopy))
	}
}

func part2(tracks [][]string, carts []Cart) Coord {
	cartMap := make(map[Coord]*Cart)
	cartCopy := append([]Cart{}, carts...)
	for len(cartCopy) > 1 {
		for i := 0; i < len(cartCopy); i++ {
			if crashed := cartCopy[i].UpdateCoordinates(tracks, &cartMap); crashed {
				cartCopy[i].SetCrashed()
				cartMap[cartCopy[i].coordinates].SetCrashed()
				delete(cartMap, cartCopy[i].coordinates)
			}
		}

		aliveCarts := []Cart{}
		for _, cart := range cartCopy {
			if cart.alive {
				aliveCarts = append(aliveCarts, cart)
			}
		}

		cartCopy = aliveCarts
		sort.Sort(ByCoord(cartCopy))
	}

	return cartCopy[0].coordinates
}

func getMineCarts(tracks *[][]string) (carts []Cart) {
	for y := 0; y < len(*tracks); y++ {
		row := (*tracks)[y]
		for x := 0; x < len(row); x++ {
			cartDirection := row[x]
			if cartDirection == "<" || cartDirection == ">" {
				(*tracks)[y][x] = "-"
				carts = append(carts, InitCart(x, y, cartDirection))
			} else if cartDirection == "^" || cartDirection == "v" {
				(*tracks)[y][x] = "|"
				carts = append(carts, InitCart(x, y, cartDirection))
			}
		}
	}

	return
}

func print2DArr(arr [][]string) {
	for _, row := range arr {
		fmt.Println(row)
	}
}
