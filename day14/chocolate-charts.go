package main

import (
	"fmt"
	"strconv"
)

var input = 681901

func main() {
	scoreboard := []int{3, 7, 1, 0, 1, 0, 1, 2, 4, 5, 1, 5, 8, 9, 1, 6, 7, 7, 9, 2}
	part1(scoreboard, 8, 4)

	fmt.Println("20321495 recipes appear to the left of the score sequence.")
}

func part1(scoreboard []int, e1, e2 int) (int, int) {
	length := len(scoreboard)

	for length < (input + 10) {
		sum := strconv.Itoa(scoreboard[e1] + scoreboard[e2])
		length += len(sum)
		for i := 0; i < len(sum); i++ {
			scoreboard = append(scoreboard, int(sum[i]-48))
		}

		e1 = (e1 + scoreboard[e1] + 1) % length
		e2 = (e2 + scoreboard[e2] + 1) % length
	}
	fmt.Printf("The next ten recipes immediately after is '")
	printArr(scoreboard[input:])
	fmt.Println("'.")

	return e1, e2
}

func printArr(arr []int) {
	for _, e := range arr {
		fmt.Print(e)
	}
}
