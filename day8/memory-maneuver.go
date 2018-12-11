package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var headerSize = 2

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parseData() []string {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	return strings.Split(strings.TrimSuffix(string(rawData), "\n"), " ")
}

func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	check(err)

	return num
}

func transformDataToIntArr(stringData []string) (intData []int) {
	for _, str := range stringData {
		intData = append(intData, strToInt(str))
	}

	return
}

func sumArr(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}

	return
}

func main() {
	solve(transformDataToIntArr(parseData()))
}

func solve(nums []int) {
	total, values, _ := postOrderTraversal(nums)

	fmt.Printf("The sum of all metadata entries is %d\n", total)
	fmt.Printf("The value of the root node is %d\n", values)
}

func postOrderTraversal(data []int) (int, int, []int) {
	childNodeNum := data[0]
	metaDataNum := data[1]
	data = data[2:]
	values := []int{}
	totals := 0

	for i := 0; i < childNodeNum; i++ {
		var total int
		var value int
		total, value, data = postOrderTraversal(data)
		totals += total
		values = append(values, value)
	}

	totals += sumArr(data[:metaDataNum])

	if childNodeNum == 0 {
		return totals, sumArr(data[:metaDataNum]), data[metaDataNum:]
	} else {
		sumValues := 0
		for _, metadata := range data[:metaDataNum] {
			if metadata > 0 && metadata <= len(values) {
				sumValues += values[metadata-1]
			}
		}

		return totals, sumValues, data[metaDataNum:]
	}
}
