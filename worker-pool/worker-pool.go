package workerpool

import (
	"fmt"
	"sync"
	"time"
)

// Task Definition
type Task struct {
	ID int
}

// way to process tasks
func (t *Task) Process() {
	fmt.Printf("Processing Task %d\n", t.ID)

	time.Sleep(2 * time.Second)
}

// worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

//function to execute the worker pool

func (wp *WorkerPool) worker(){
	for task := range wp.taskChan{
		task.Process()
		wp.wg.Done()
	}
}

