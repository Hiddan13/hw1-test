package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var wg = sync.WaitGroup{}

func main() {
	tasksCount := 50
	tasks := make([]Task, 0, tasksCount)

	var runTasksCount int32

	for i := 0; i < tasksCount; i++ {
		err := fmt.Errorf("error from task %d", i)
		tasks = append(tasks, func() error {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			atomic.AddInt32(&runTasksCount, 1)
			return err
		})
	}

	workersCount := 10
	maxErrorsCount := 23
	err := Run(tasks, workersCount, maxErrorsCount)
	if err != nil {
		fmt.Println("error running tasks", err)
	}

	wg.Wait()

	fmt.Println(runTasksCount, "<=", workersCount+maxErrorsCount)
}

// func main() {

// 	tasksCount := 50
// 	tasks := make([]Task, 0, tasksCount)

// 	var runTasksCount int32
// 	var sumTime time.Duration

// 	for i := 0; i < tasksCount; i++ {
// 		taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
// 		sumTime += taskSleep

// 		tasks = append(tasks, func() error {
// 			time.Sleep(taskSleep)
// 			atomic.AddInt32(&runTasksCount, 1)
// 			return nil
// 		})
// 	}

// 	workersCount := 5
// 	maxErrorsCount := 1

// 	err := Run(tasks, workersCount, maxErrorsCount)

// 	if err != nil {
// 		fmt.Println("error running tasks", err)
// 	}

// 	wg.Wait()

// 	fmt.Println("runTasksCount:", runTasksCount, "tasksCount", tasksCount)
// }

func Run(tasks []Task, n, m int) error {
	ch := make(chan error)
	Err := 0

	go func() {
		for _, w := range tasks {
			if Err <= m {
				ch <- w()

			}
		}
		close(ch)
	}()
	for i := 0; i < len(tasks); i++ { // n
		wg.Add(1)

		go func() {
			for {
				t := <-ch
				if t == nil {
					break
				} else if t != nil {

					Err++

					break

				}

			}
			wg.Done()
		}()

	}
	if Err >= m {
		return ErrErrorsLimitExceeded
	}
	return nil

}
