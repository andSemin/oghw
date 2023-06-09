package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		ch := make(Bi)
		go func(in In) {
			defer close(ch)
			for {
				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					ch <- v
				case <-done:
					return
				}
			}
		}(in)
		in = stage(ch)
	}

	return in
}
