package main

import (
	"flag"
	"strings"
	"sync"
	"unicode"
)

// for choosing which version should run, and the parameters of that run
var (
	basic       = flag.Bool("basic", false, "the hello-world level input set")
	complicated = flag.Bool("complicated", false, "the complicated input set")
	channeled   = flag.Bool("channeled", false, "run the full input set on the channeled async version")
	sliced      = flag.Bool("sliced", false, "run the full input set on the sliced async version")
	lorem       = flag.Bool("lorem", false, "runs the lorem ipsum set for the async test")
	routines    = flag.Int("routines", 2, "count of goroutines to run for async test sets")
)

func main() {
	flag.Parse()
	if *basic || (!*basic && !*complicated && !*channeled && !*sliced) {
		runBasicSet()
	}
	if *complicated {
		runComplicatedSet()
	}
	if *channeled {
		if *lorem {
			runLoremSet(vChanneled)
		} else {
			runAsyncSet(vChanneled)
		}
	}
	if *sliced {
		if *lorem {
			runLoremSet(vSliced)
		} else {
			runAsyncSet(vSliced)
		}
	}
}

// --------------------------------------
// primary functions

type wordCounts map[string]int

// standard func: take in a string, return each word and its counts
func mrWordCount(input string) wordCounts {
	results := make(wordCounts)
	words := strings.Fields(input)
	for _, w := range words {
		results[normalize(w)]++
	}
	return results
}

// normalizes a string to only characters and numbers, in lower case
func normalize(word string) string {
	norm := []rune{}
	runes := []rune(word)
	for _, r := range runes {
		if unicode.IsLetter(r) {
			norm = append(norm, unicode.ToLower(r))
		} else if unicode.IsNumber(r) || unicode.IsDigit(r) {
			norm = append(norm, r)
		}
	}
	return string(norm)
}

// runs a collection of inputs into mrWordCount using
// a channel to communicate the results across goroutines.
func mrChanneledCounter(workers int, inputs []string) wordCounts {
	resultSet := make([]wordCounts, workers)

	// a go-routineable func for wordcounting inputs on a channel
	f := func(idx int, wg *sync.WaitGroup, inputCh chan string) {
		localResults := make(wordCounts)
		defer func() {
			resultSet[idx] = localResults
			wg.Done()
		}()
		for in := range inputCh {
			assign(localResults, mrWordCount(in))
		}
	}

	// spin up all the goroutines and feed in the inputs
	wg := new(sync.WaitGroup)
	wg.Add(workers)
	inputFeed := make(chan string)
	for i := 0; i < workers; i++ {
		go f(i, wg, inputFeed)
	}
	for _, input := range inputs {
		inputFeed <- input
	}

	// wait for all the goroutines to complete
	close(inputFeed)
	wg.Wait()

	// aggregate the results
	results := make(wordCounts)
	for _, r := range resultSet {
		assign(results, r)
	}
	return results
}

// runs a collection of inputs into mrWordCount using
// slices to divvy up the input
func mrSlicedCounter(workers int, inputs []string) wordCounts {
	resultSet := make([]wordCounts, workers)

	// a go-routinable func for wordcounting inputs on a slice
	f := func(idx int, wg *sync.WaitGroup, localInputs []string) {
		localResults := make(wordCounts)
		defer func() {
			resultSet[idx] = localResults
			wg.Done()
		}()
		for _, in := range localInputs {
			assign(localResults, mrWordCount(in))
		}
	}

	// spin up all the goroutines and feed in the inputs
	wg := new(sync.WaitGroup)
	wg.Add(workers)

	// create a slice of inputs for each goroutine to process
	chunk := 1
	inLen := len(inputs)
	if inLen > workers {
		chunk = inLen / workers
	}
	mod := inLen % workers
	for i := 0; i < workers; i++ {
		s := i * chunk
		e := s + chunk

		if s >= inLen { // when we have more workers than inputs
			go f(i, wg, nil)
			continue
		}

		// the last worker takes the last chunk + remaining mod
		if e+mod == inLen {
			e += mod
		}

		// normal chunking, give each worker a portion of inputs
		go f(i, wg, inputs[s:e])
	}

	// wait for all the goroutines to complete
	wg.Wait()

	// aggregate the results
	results := make(wordCounts)
	for _, r := range resultSet {
		assign(results, r)
	}
	return results
}

// merges two wordCounts, assigning all entries in the 'from' onto the 'to'
// precondition: 'to' must be non-nil
func assign(to, from wordCounts) {
	for w, c := range from {
		to[w] += c
	}
}
