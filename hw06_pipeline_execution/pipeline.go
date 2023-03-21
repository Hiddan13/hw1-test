package main

import (
	"fmt"
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

const (
	sleepPerStage1 = time.Millisecond * 100
	fault1         = sleepPerStage1 / 2
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	interfaceChan := make(chan interface{})

	fmt.Println("len stages", len(stages))
	go func() {

		for _, i := range stages {
			select {
			case <-done:
				return
			case interfaceChan <- i(in):
				i(interfaceChan)

			}
		}
	}()

	//defer close(interfaceChan)
	return interfaceChan

	// Place your code here.

}
func main() {
	g := func(_ string, f func(v interface{}) interface{}) Stage {
		return func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for v := range in {
					time.Sleep(sleepPerStage1)
					out <- f(v)
				}
			}()
			return out
		}
	}

	stages := []Stage{
		g("Dummy", func(v interface{}) interface{} { return v }),
		g("Multiplier (* 2)", func(v interface{}) interface{} { return v.(int) * 2 }),
		g("Adder (+ 100)", func(v interface{}) interface{} { return v.(int) + 100 }),
		//g("Stringifier", func(v interface{}) interface{} { return v.(int) }),
	}

	in := make(Bi)
	data := []int{1, 2, 3, 4, 5}

	go func() {
		for _, v := range data {
			in <- v
		}
		close(in)
	}()

	result := make([]string, 0, 10)
	start := time.Now()
	for s := range ExecutePipeline(in, nil, stages...) {
		fmt.Println(s)
		fmt.Printf("%T\n", s) //Судя по тесту я должен записать сбда строку , но у меня тут <- chan interface{}
		//result = append(result, s.(string))//из за этого я не могу преобразовать в стринг

	}
	fmt.Println("result", result)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	time.Sleep(10000 * time.Millisecond)

}
