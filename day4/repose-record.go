package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
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

	guardRecords := strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
	sort.Strings(guardRecords)

	guardRecordsMap := buildGuardSchedule(guardRecords)
	// printSchedule(&guardRecordsMap)

	maxGuardId := findGuardWithMostMinAsleep(&guardRecordsMap)
	intGuardId, err := strconv.Atoi(maxGuardId)
	check(err)

	maxTimeIndex, _ := findTimeAsleepMost(guardRecordsMap[maxGuardId])

	fmt.Printf("Part 1: The ID of the guard multiplied by minute is %d\n", intGuardId*maxTimeIndex)

	guardWithMostAsleepByMinute, maxTime := findGuardWithMostAsleepByMinute(&guardRecordsMap)
	intGuardId2, err := strconv.Atoi(guardWithMostAsleepByMinute)
	fmt.Printf("Part 2: The ID of the guard multiplied by minute is %d\n", intGuardId2*maxTime)
}

func getTime(record string) int {
	time, err := strconv.Atoi(record[15:17])
	check(err)

	return time
}

func getIsSleeping(record string) bool {
	return record[19] == 'f'
}

func getIsGuardSwitch(record string) bool {
	return record[19] == 'G'
}

func getGuardId(record string) string {
	return strings.Replace(record[26:], " begins shift", "", -1)
}

func buildGuardSchedule(guardRecords []string) map[string][60]int {
	guardRecordsMap := make(map[string][60]int)

	var currentGuardId string
	var sleepingStart int
	for _, record := range guardRecords {
		time := getTime(record)

		if getIsGuardSwitch(record) {
			currentGuardId = getGuardId(record)
		} else {
			if getIsSleeping(record) {
				sleepingStart = time
			} else {
				timeArr := guardRecordsMap[currentGuardId]
				for i := sleepingStart; i < time; i++ {
					timeArr[i]++
				}

				guardRecordsMap[currentGuardId] = timeArr

				sleepingStart = 0
			}
		}
	}

	return guardRecordsMap
}

func findGuardWithMostMinAsleep(guardRecordsMap *map[string][60]int) string {
	maxSum := 0
	var maxGuardId string
	for guardId, schedule := range *guardRecordsMap {
		if sum := sumArray(schedule); sum > maxSum {
			maxSum = sum
			maxGuardId = guardId
		}
	}

	return maxGuardId
}

func findTimeAsleepMost(guardRecords [60]int) (maxTimeIndex, maxTime int) {
	for i, num := range guardRecords {
		if num > maxTime {
			maxTime = num
			maxTimeIndex = i
		}
	}

	return
}

func findGuardWithMostAsleepByMinute(guardRecordsMap *map[string][60]int) (maxGuardId string, maxTime int) {
	maxByMinute := 0
	for guardId, schedule := range *guardRecordsMap {
		if i, max := findTimeAsleepMost(schedule); max > maxByMinute {
			maxGuardId = guardId
			maxTime = i
			maxByMinute = max
		}
	}

	return
}

func sumArray(arr [60]int) (sum int) {
	for _, e := range arr {
		sum += e
	}

	return
}

func findMax(arr [60]int) (max int) {
	for _, e := range arr {
		if e > max {
			max = e
		}
	}

	return
}

func printSchedule(guardRecordsMap *map[string][60]int) {
	for guardId, schedule := range *guardRecordsMap {
		fmt.Printf("%s ==> %v\n", guardId, schedule)
	}
}
