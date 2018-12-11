package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Queue struct {
	queue *list.List
}

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

	queue := Queue{list.New()}

	// queue.initQueue(rootMap)
	// fmt.Printf("The order the instructions should be completed is \"%v\".\n", part1(queue, parentMap, childrenMap))

	queue.initQueue(rootMap)
	fmt.Println(part2(queue, parentMap, childrenMap))
}

func part1(q Queue, parentMap, childrenMap map[string][]string) string {
	var instructionOrder []string
	for first := q.queue.Front(); first != nil; first = q.queue.Front() {
		orderLetter := first.Value.(string)
		instructionOrder = append(instructionOrder, orderLetter)

		for _, child := range parentMap[orderLetter] {
			allParentsVisited := true
			for _, parent := range childrenMap[child] {
				if !includes(instructionOrder, parent) {
					allParentsVisited = false
					break
				}
			}

			if allParentsVisited {
				q.insertNode(child)
			}
		}

		q.queue.Remove(first)
	}

	return strings.Join(instructionOrder, "")
}

func part2(q Queue, parentMap, childrenMap map[string][]string) (time int) {
	workers := InitWorkers()
	var instructionOrder []string
	for len(instructionOrder) != 26 {
		availableWorkersIndex := GetAvailableWorkersIndex(&workers)

		for _, i := range availableWorkersIndex {
			if q.len() != 0 {
				nextTask := q.pop()
				workers[i].SetTask(nextTask)
			}
		}

		minWorkerTime := GetMinWorkerTime(&workers)
		UpdateAllWorkers(&workers, minWorkerTime)

		for i := 0; i < 5; i++ {
			if workers[i].done {
				instructionOrder = append(instructionOrder, workers[i].task)
				for _, child := range parentMap[workers[i].task] {
					allParentsVisited := true
					for _, parent := range childrenMap[child] {
						if !includes(instructionOrder, parent) {
							allParentsVisited = false
							break
						}
					}

					if allParentsVisited {
						q.insertNode(child)
					}
				}

				workers[i].done = false
			}
		}

		time += minWorkerTime
	}

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

func (q *Queue) initQueue(theMap map[string]bool) {
	for key := range theMap {
		q.insertNode(key)
	}

	return
}

func (q *Queue) insertNode(step string) {
	if q.len() == 0 {
		q.queue.PushBack(step)
		return
	}

	if q.queue.Front().Value.(string) > step {
		q.queue.PushFront(step)
		return
	}

	for start := q.queue.Front(); start != nil; start = start.Next() {
		value := start.Value.(string)
		if value == step {
			return
		} else if value > step {
			q.queue.InsertBefore(step, start)
			return
		}
	}

	q.queue.PushBack(step)
}

func (q *Queue) removeNode(step string) {
	if q.len() == 0 {
		return
	}

	for start := q.queue.Front(); start != nil; start = start.Next() {
		if start.Value.(string) == step {
			q.queue.Remove(start)
			return
		}
	}
}

func (q *Queue) pop() string {
	if q.len() == 0 {
		return ""
	}

	first := q.queue.Front()
	q.queue.Remove(first)

	return first.Value.(string)
}

func (q *Queue) len() int {
	return q.queue.Len()
}

func includes(arr []string, value string) bool {
	if value == "" || len(arr) == 0 {
		return false
	}

	for _, e := range arr {
		if e == value {
			return true
		}
	}

	return false
}

func printQueue(queue *list.List) {
	fmt.Printf("Print Queue => ")
	for node := queue.Front(); node != nil; node = node.Next() {
		fmt.Printf("%s ", node.Value)
	}

	fmt.Println()
}
