package main

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	valin := in
	for _, stage := range stages {
		valin = resiver(valin, done)
		valin = stage(valin)
	}
	return valin
}

func resiver(in, done In) Out {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case valIn, ok := <-in:
				if !ok {
					return
				}
				out <- valIn
			case <-done:
				return
			}
		}
	}()
	return out
}
