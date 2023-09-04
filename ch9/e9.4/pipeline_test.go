package pipeline

import (
	"flag"
	"fmt"
	"strings"
	"testing"
	"time"
)

var runGetMaxPipes = flag.Bool("runGetMaxPipes", false, "TestGetMaxPipes: run out of memory to get the max pipes can be created")
var rounds = flag.Int("rounds", 20, "TestExecutionTime: transit the entire pipeline rounds times")
var numOfValues = flag.Int("numOfValues", 100, "TestMutipleValueExecutionTime: send numOfValues values to the pipeline at a time")
var testIn chan<- interface{}
var testOut <-chan interface{}

func init() {
	testing.Init()
	flag.Parse()
	if !testing.Verbose() {
		return
	}
	// flag.Usage()
	if *runGetMaxPipes {
		return
	}
	fmt.Println("Start to create a pipeline with", readableInteger(MAX_PIPES), "pipes.")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var counter int
	testIn, testOut = Pipeline(MAX_PIPES, &counter)
	for range ticker.C {
		fmt.Println("\t", readableInteger(counter), "pipes created...")
		if counter == MAX_PIPES {
			fmt.Println("Done!")
			break
		}
	}
	fmt.Println("Start to warm up the pipeline, transit it the first time.")
	start := time.Now()
	testIn <- interface{}(0)
	<-testOut
	fmt.Printf("\tIt took %v.\n", time.Since(start))
	fmt.Println("Done!")
}

// Run out of memory to get the max pipes which can be created
// on your machine.
//
// go test -v -run=^TestGetMaxPipes$ -runGetMaxPipes=true > maxpipes.txt
func TestGetMaxPipes(t *testing.T) {
	if !testing.Verbose() {
		t.Fatalf("Please run it under verbose mode.\n")
	}
	if !*runGetMaxPipes {
		fmt.Println(`Run out of memory to get the max pipes which can be created on your machine.
go test -v -run=^TestGetMaxPipes$ -runGetMaxPipes=true > maxpipes.txt`)
		return
	}
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	var counter int
	Pipeline(MAX_PIPES*10, &counter)
	for range ticker.C {
		fmt.Println(counter)
	}
}

// Transit the entire pipeline which has 2,000,000 pipes 20 times to get
// the average pipeline/pipe execution time.
func TestExecutionTime(t *testing.T) {
	if !testing.Verbose() {
		t.Fatalf("Please run it under verbose mode.\n")
	}
	if *runGetMaxPipes {
		return
	}
	fmt.Printf("We are going to transit the entire pipeline %d times.\n", *rounds)
	var total int64
	for i := 0; i < *rounds; i++ {
		start := time.Now()
		testIn <- interface{}(i)
		v := (<-testOut).(int)
		diff := time.Since(start)
		fmt.Printf("\t%d: %v\n", v+1, diff)
		total += diff.Nanoseconds()
	}

	fmt.Printf("Average pipeline/pipe execution time = %s/%s nanoseconds.\n",
		readableInteger(total/int64(*rounds)), readableInteger(total/int64((*rounds)*MAX_PIPES)))
}

// Send 100 values to the pipeline at a time to get the whole execution time.
func TestMutipleValueExecutionTime(t *testing.T) {
	if !testing.Verbose() {
		t.Fatalf("Please run it under verbose mode.\n")
	}
	if *runGetMaxPipes {
		return
	}
	fmt.Printf("We are going to send %d values to the pipeline at a time.\n", *numOfValues)
	start := time.Now()
	starts := make([]time.Time, *numOfValues)
	go func() {
		for i := 0; i < *numOfValues; i++ {
			starts[i] = time.Now()
			testIn <- interface{}(i)
		}
		// This takes almost no time, starts[] should be the same as start.
		fmt.Printf("Finish sending %d values: %v\n", *numOfValues, time.Since(start))
	}()
	for i := 0; i < *numOfValues; i++ {
		v := (<-testOut).(int)
		fmt.Printf("\t%d: %v\n", v+1, time.Since(starts[v]))
	}
	diff := time.Since(start)
	fmt.Printf("Whole execution time of %d values at a time = %s nanoseconds\n",
		*numOfValues, readableInteger(diff.Nanoseconds()))
}

func readableInteger[T int | int64 | uint | uint64](i T) string {
	s := fmt.Sprintf("%d", i)
	sb := &strings.Builder{}
	if s[0] == '-' {
		s = s[1:]
		sb.WriteString("-")
	}
	for j, k := len(s)%3, 0; j <= len(s); k, j = j, j+3 {
		if k != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(s[k:j])
	}
	return sb.String()
}
