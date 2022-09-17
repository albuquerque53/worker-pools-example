package theorical

import (
	"fmt"
	"time"
)

func MoveBooks() {
	const nTasks = 10
	const nWorkers = 3

	// In worker pools we have tasks that must be processed
	tasks := make(chan int, nTasks) // tasks is a list of tasks

	// In wp we also have results that are the result of task processing
	results := make(chan int, nTasks) // results is a list of results

	// Using the worker pools!!
	for worker := 1; worker < nWorkers; worker++ {
		go work(worker, tasks, results)
	}

	// Adding values to tasks
	for value := 0; value < nTasks; value++ {
		tasks <- value
	}
	close(tasks)

	// Catching the results
	for i := 0; i < nTasks; i++ {
		<-results
	}
}

// this is our worker function, inside of it are the bussiness logic that
// workers will follow
func work(worker int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		// simple logging
		fmt.Printf("[WORKER_%d] Working with TASK_%d\n", worker, task)

		// in this case, the process result is the product of task per 100
		processRes := task * 100
		time.Sleep(time.Second * 1)
		results <- processRes

		// simple logging
		fmt.Printf("[WORKER_%d] Finished TASK_%d with success!\n", worker, task)
	}
}
