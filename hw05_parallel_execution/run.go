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

func doSmthWork(wg *sync.WaitGroup, read_from_chTask <-chan Task, read_from_chError chan<- error, read_from_chDone <-chan struct{}) {
	defer wg.Done()

	for {
		select {
		case <-read_from_chDone:
			return
		case task, ok := <-read_from_chTask:
			if !ok {
				return
			}
			err := task()
			if err != nil {
				read_from_chError <- err
			}
		}
	}
}
