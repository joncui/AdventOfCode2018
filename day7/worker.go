package main

type Worker struct {
	task string
	time int
	done bool
}

func InitWorker() Worker {
	return Worker{"", 0, false}
}

func (w *Worker) SetTask(task string) {
	w.task = task
	w.time = int(task[0] - 4)
	w.done = false
}

func (w *Worker) IsBusy() bool {
	return w.time > 0
}

func (w *Worker) NextTimeStep(stepSize int) {
	if w.IsBusy() {
		w.time -= stepSize
		if w.time == 0 {
			w.done = true
		}
	}
}

func InitWorkers() (workers [5]Worker) {
	for i := 0; i < 5; i++ {
		workers[i] = InitWorker()
	}

	return
}

func UpdateAllWorkers(workers *[5]Worker, amount int) {
	for i := 0; i < 5; i++ {
		workers[i].NextTimeStep(amount)
	}
}

func GetMinWorkerTime(workers *[5]Worker) (minTime int) {
	minTime = 100
	for _, worker := range *workers {
		if worker.time > 0 && worker.time < minTime {
			minTime = worker.time
		}
	}

	return
}

func GetAvailableWorkersIndex(workers *[5]Worker) (availableWorkersIndex []int) {
	for i, worker := range *workers {
		if !worker.IsBusy() {
			availableWorkersIndex = append(availableWorkersIndex, i)
		}
	}

	return
}
