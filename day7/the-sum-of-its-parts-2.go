package main

import (
	"container/list"
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

func parseData() []string {
	file := "data"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	rawData, err := ioutil.ReadFile(file)
	check(err)

	return strings.Split(strings.TrimSuffix(string(rawData), "\n"), "\n")
}

func main() {
	instructions := parseData()

	parentMap := make(map[string][]string)
	childrenMap := make(map[string][]string)
	rootMap := make(map[string]bool)

	for _, instruction := range instructions {
		startingStep := string(instruction[5])
		nextStep := string(instruction[36])

		insertToMap(&parentMap, startingStep, nextStep)
		insertToMap(&childrenMap, nextStep, startingStep)

		if _, ok := childrenMap[startingStep]; !ok {
			rootMap[startingStep] = true
		}

		if value := rootMap[nextStep]; value {
			delete(rootMap, nextStep)
		}
	}

	queue := list.New()
	initQueue(queue, rootMap)

	fmt.Printf("The order the instructions should be completed is \"%v\".\n", part1(queue, parentMap, childrenMap))

	initQueue(queue, rootMap)
	fmt.Println(part2(queue, parentMap, childrenMap))
}

func part1(queue *list.List, parentMap, childrenMap map[string][]string) string {
	var instructionOrder []string
	for first := queue.Front(); first != nil; first = queue.Front() {
		orderLetter := first.Value.(string)
		instructionOrder = append(instructionOrder, orderLetter)
		nextSteps := parentMap[orderLetter]

		for _, child := range nextSteps {
			parents := childrenMap[child]
			for i, parent := range parents {
				if parent == orderLetter {
					parents = append(parents[0:i], parents[i+1:]...)
				}
			}

			if len(parents) == 0 {
				insertNode(queue, child)
			}

			childrenMap[child] = parents
		}

		queue.Remove(first)
	}

	return strings.Join(instructionOrder, "")
}

func part2(queue *list.List, parentMap, childrenMap map[string][]string) (time int) {
	var workers []int
	for time = 0; ; {
		updateAllWorkers(&workers, -1)

		if queue.Len() == 0 || len(workers) == 5 {
			continue
		}

		if queue.Len() == 0 && len(workers) == 0 {
			break
		}

		for letter := queue.Front(); len(workers) < 5 && queue.Len() > 0; letter = letter.Next() {
			workers = append(workers, strToTime(letter.Value.(string)))
		}

		minTime := getMinWorkerTime(&workers)
		time += minTime
		updateAllWorkers(&workers, minTime*-1)

	}
	// for first := queue.Front(); first != nil; first = queue.Front() {
	//	orderLetter := first.Value.(string)
	//	nextSteps := parentMap[orderLetter]

	//	for _, child := range nextSteps {
	//		parents := childrenMap[child]
	//		for i, parent := range parents {
	//			if parent == orderLetter {
	//				parents = append(parents[0:i], parents[i+1:]...)
	//			}
	//		}

	//		if len(parents) == 0 {
	//			insertNode(queue, child)
	//		}

	//		childrenMap[child] = parents
	//	}

	//	queue.Remove(first)
	// }

	return
}

func insertToMap(theMap *(map[string][]string), key, value string) {
	node, ok := (*theMap)[key]
	if ok {
		node = append(node, value)
	} else {
		node = []string{value}
	}

	(*theMap)[key] = node
}

func initQueue(queue *list.List, theMap map[string]bool) {
	for key := range theMap {
		insertNode(queue, key)
	}

	return
}

func insertNode(queue *list.List, step string) {
	if queue.Len() == 0 {
		queue.PushBack(step)
		return
	}

	if queue.Front().Value.(string) > step {
		queue.PushFront(step)
		return
	}

	for start := queue.Front(); start != nil; start = start.Next() {
		value := start.Value.(string)
		if value == step {
			return
		} else if value > step {
			queue.InsertBefore(step, start)
			return
		}
	}

	queue.PushBack(step)
}

func removeNode(queue *list.List, step string) {
	if queue.Len() == 0 {
		return
	}

	for start := queue.Front(); start != nil; start = start.Next() {
		if start.Value.(string) == step {
			queue.Remove(start)
			return
		}
	}
}

func printQueue(queue *list.List) {
	fmt.Printf("Print Queue => ")
	for node := queue.Front(); node != nil; node = node.Next() {
		fmt.Printf("%s ", node.Value)
	}

	fmt.Println()
}

func strToTime(str string) int {
	return int(str[0] - 64)
}

func updateAllWorkers(workers *[]int, amount int) {
	for i, time := range *workers {
		if time += amount; time == 0 {
			*workers = append((*workers)[:i], (*workers)[i+1:]...)
		}
	}
}

func getMinWorkerTime(workers *[]int) (minTime int) {
	minTime = 100
	for _, time := range *workers {
		if time < 100 {
			minTime = time
		}
	}

	return
}

func allWorkersBusy(workers *[]int) bool {
	for _, time := range *workers {
		if time == 0 {
			return false
		}
	}

	return true
}
