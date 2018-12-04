package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type parsedClaimType struct {
	claimId, leftEdge, topEdge, width, height int
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

	claims := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
	overlapMap := make(map[coord]bool)
	coverMap := make(map[coord]bool)
	var parsedClaims []parsedClaimType

	for _, claim := range claims {
		parsedClaim := parseClaim(claim)
		parsedClaims = append(parsedClaims, parsedClaim)

		for y := parsedClaim.topEdge; y < parsedClaim.topEdge+parsedClaim.height; y++ {
			for x := parsedClaim.leftEdge; x < parsedClaim.leftEdge+parsedClaim.width; x++ {
				coordinates := coord{x, y}
				if overlapMap[coordinates] {
					continue
				} else if coverMap[coordinates] {
					overlapMap[coordinates] = true
					delete(coverMap, coordinates)
				} else {
					coverMap[coordinates] = true
				}
			}
		}
	}

	fmt.Printf("There are %d square inches of fabric with two or more claims.\n", len(overlapMap))

	for _, parsedClaim := range parsedClaims {
		noOverlap := checkClaimNoOverlap(parsedClaim, &coverMap)

		if noOverlap {
			fmt.Printf("Claim ID %d is the only claim that doesn't overlap.\n", parsedClaim.claimId)
			break
		}
	}
}

func parseClaim(claim string) (res parsedClaimType) {
	re := regexp.MustCompile("\\#(\\d+)\\s\\@\\s(\\d{0,3}),(\\d{0,3}):\\s(\\d+)x(\\d+)")
	match := re.FindStringSubmatch(claim)

	claimId, err := strconv.Atoi(match[1])
	check(err)

	leftEdge, err := strconv.Atoi(match[2])
	check(err)

	topEdge, err := strconv.Atoi(match[3])
	check(err)

	width, err := strconv.Atoi(match[4])
	check(err)

	height, err := strconv.Atoi(match[5])
	check(err)

	return parsedClaimType{claimId, leftEdge, topEdge, width, height}
}

func checkClaimNoOverlap(parsedClaim parsedClaimType, coverMap *map[coord]bool) (noOverlap bool) {
	for y := parsedClaim.topEdge; y < parsedClaim.topEdge+parsedClaim.height; y++ {
		for x := parsedClaim.leftEdge; x < parsedClaim.leftEdge+parsedClaim.width; x++ {
			coordinates := coord{x, y}
			if !(*coverMap)[coordinates] {
				return false
			}
		}
	}

	return true
}
