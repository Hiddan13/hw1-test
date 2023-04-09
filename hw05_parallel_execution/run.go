package main

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func main() {
}

func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup

	chTask := make(chan Task)
	chError := make(chan error)
	chDone := make(chan struct{})

	for i := 0; i < n; i++ {
		wg.Add(1)
		go doSmthWork(&wg, chTask, chError, chDone)
	}

	go func() {
		defer close(chTask)

		for _, task := range tasks {
			select {
			case <-chDone:
				return
			case chTask <- task:
			}
		}
	}()
	var realError bool
	go func() {
		ErrNum := 0
		for range chError {
			ErrNum++
			if ErrNum == m {
				realError = true
				close(chDone)
			}
		}
	}()
	wg.Wait()
	close(chError)
	if realError {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func doSmthWork(wg *sync.WaitGroup, readChTask <-chan Task, readchError chan<- error, readchDone <-chan struct{}) {
	defer wg.Done()

	for {
		select {
		case <-readchDone:
			return
		case task, ok := <-readChTask:
			if !ok {
				return
			}
			err := task()
			if err != nil {
				readchError <- err
			}
		}
	}
}
