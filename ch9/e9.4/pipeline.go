package pipeline

// On my machine, it is possible to create 3700000 goroutines,
// but performance of this pipeline is very poor.
// goos: windows
// goarch: amd64
// cpu: AMD Ryzen 5 4600U with Radeon Graphics
const MAX_PIPES = 2000000

// Creates a pipeline with numOfPipes pipes, and returns an input
// channel and an output channel.
//
// The counter represents the number of pipes that have been created
// during the creation process.
func Pipeline(numOfPipes int, counter *int) (chan<- interface{}, <-chan interface{}) {
	if numOfPipes <= 0 {
		return nil, nil
	}
	in := make(chan interface{})
	out := make(chan interface{})
	if counter == nil {
		counter = new(int)
	}
	*counter = 0
	go func() { pipe(in, out, counter, numOfPipes) }()
	return in, out
}

func pipe(in <-chan interface{}, out chan<- interface{}, counter *int, numOfPipes int) {
	*counter++
	var next chan<- interface{}
	if *counter == numOfPipes {
		next = out
	} else {
		nnext := make(chan interface{})
		next = nnext
		go pipe(nnext, out, counter, numOfPipes)
	}
	for v := range in {
		next <- v
	}
}
