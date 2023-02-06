package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var Err int

// type Task func() error // что тут нужно реализовывать?
type Task struct {
	Value int
	Error bool
}

var Slice = []Task{
	Task{300, false},
	Task{500, true},
	Task{200, false},
	Task{600, true},
	Task{1200, false},
	Task{6460, true},
	Task{2400, false},
	Task{2200, false},
	Task{6600, true},
	Task{7500, false},
	Task{6000, false},
}
var wg = sync.WaitGroup{}

func main() {
	start := time.Now()
	Run(Slice, 10, 6)

	wg.Wait()
	duration := time.Since(start)
	fmt.Println("Время работы функции:", duration)
}
func Run(tasks []Task, n, m int) error {
	ch := make(chan Task)
	go func() {
		for _, task := range tasks {
			if Err < m {
				ch <- task
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
				t, ok := <-ch
				if !ok {
					break
				}
				if t.Error {
					Err++
				}

				time.Sleep(time.Duration(t.Value) * time.Millisecond)
				fmt.Println(t.Value)

				fmt.Println("Number of runnable goroutines: ", runtime.NumGoroutine())
			}
			wg.Done()

		}()

	}
	if Err >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
