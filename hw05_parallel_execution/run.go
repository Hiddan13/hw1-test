package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// type Task struct {
// 	V int
// 	R bool
// }

var (
	Err                    int
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
)

// var m int
type Task func() error

//	var Slice = []Task{
//		Task{300, false},
//		Task{500, true},
//		Task{200, false},
//		Task{600, true},
//		Task{1200, false},
//		Task{6460, true},
//		Task{2400, false},
//		Task{2200, false},
//		Task{6600, true},
//		Task{7500, false},
//		Task{6000, false},
//	}
var wg = sync.WaitGroup{}

func main() {
	tasksCount := 50
	tasks := make([]Task, 0, tasksCount)

	var runTasksCount int32
	var sumTime time.Duration

	for i := 0; i < tasksCount; i++ {

		taskSleep := time.Millisecond * time.Duration(rand.Intn(100))

		sumTime += taskSleep

		tasks = append(tasks, func() error { // Почему он не добавляет в слайс ? Вернее он добавляет, но что то не понятное

			time.Sleep(taskSleep)

			atomic.AddInt32(&runTasksCount, 1) // из за этого не заходит сюда и не возвращает false в error

			return nil
		})
	}

	workersCount := 5
	maxErrorsCount := 1

	// start := time.Now()
	err := Run(tasks, workersCount, maxErrorsCount)
	// err := Run(tasks, workersCount, maxErrorsCount)
	if err != nil {
		fmt.Println("error running tasks", Err, err)
	}
	defer fmt.Println(runTasksCount)

	wg.Wait()
}

func Run(tasks []Task, n, m int) error {
	ch := make(chan Task)

	go func() {
		for _, w := range tasks {
			fmt.Println(&w)
			if Err < m {
				ch <- w
				fmt.Println("write in ch")
			} else {
				fmt.Println("много ошибок")
				break
			}
		}
		close(ch)
	}()
	for i := 0; i < n; i++ { // n
		wg.Add(1)

		go func() {
			for {
				_, ok := <-ch
				if !ok {
					fmt.Println("err = ", ok)
					break
				} else {
					fmt.Println("err =", ok)
					Err++
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
