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

	polymer := strings.TrimSpace(string(rawData))

	fmt.Printf("%d units remain after fully reacting the polymer.\n", part1(polymer))

	fmt.Printf("The length of the shortest polymer is %d.\n", part2(polymer))
}

func part1(polymer string) int {
	return getFullyReactedPolymerSizeWithStack(polymer)
}

func part2(polymer string) int {
	alpha := "abcdefghijklmnopqrstuvwxyz"

	minLength := 50000
	for _, c := range alpha {
		replacer := strings.NewReplacer(string(c), "", strings.ToUpper(string(c)), "")
		newPolymer := replacer.Replace(polymer)
		length := getFullyReactedPolymerSizeWithStack(newPolymer)

		if length < minLength {
			minLength = length
		}
	}

	return minLength
}

// First attempt (brute force)
func getFullyReactedPolymerSize(polymer string) int {
	masterReplacer := strings.NewReplacer("aA", "", "Aa", "", "bB", "", "Bb", "", "cC", "", "Cc", "", "dD", "", "Dd", "", "eE", "", "Ee", "", "fF", "", "Ff", "", "gG", "", "Gg", "", "hH", "", "Hh", "", "iI", "", "Ii", "", "jJ", "", "Jj", "", "kK", "", "Kk", "", "lL", "", "Ll", "", "mM", "", "Mm", "", "nN", "", "Nn", "", "oO", "", "Oo", "", "pP", "", "Pp", "", "qQ", "", "Qq", "", "rR", "", "Rr", "", "sS", "", "Ss", "", "tT", "", "Tt", "", "uU", "", "Uu", "", "vV", "", "Vv", "", "wW", "", "Ww", "", "xX", "", "Xx", "", "yY", "", "Yy", "", "zZ", "", "Zz", "")

	startingLength := len(polymer) + 1
	for startingLength > len(polymer) {
		startingLength = len(polymer)
		polymer = masterReplacer.Replace(polymer)
	}

	return len(polymer)
}

// Second attempt with algo from reddit, MUCH FASTER!
func getFullyReactedPolymerSizeWithStack(polymer string) int {
	var stack []rune

	for _, c := range polymer {
		// 'a' XOR 'A' == 32
		// test if character is the same letter but different cases
		if len(stack) > 0 && (c^stack[len(stack)-1]) == 32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack)
}
