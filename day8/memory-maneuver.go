package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	return strings.Split(strings.TrimSuffix(string(rawData), " "), "\n")
}

func main() {
	nums := parseData()
	// fmt.Println(nums)
	fmt.Println("vim go")

	licenseFileCheck := 0
	stack := []int{}
	for i := 0; i < len(nums); {
		childNodeNum := strToNum(nums[i])
		metaDataNum := strToNum(nums[i+1])

		if childNodeNum == 0 {
			for _, metaData := range nums[i+2 : i+2+metaDataNum] {
				licenseFileCheck += strToNum(metaData)
			}

			i += (2 + metaDataNum)
		} else {
			stack = append(stack, metaDataNum)
		}
	}
}

func strToNum(str string) int {
	num, err := strconv.Atoi(str)
	check(err)

	return num
}
