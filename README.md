# MrWordCounter
A goofy code-golf with coworkers.  The premise is simple: write a function that takes in a string as an input, and outputs a data structure expressing the number of times each word appears in that string.

## Examples:  

* In -> "hello world"  
Out -> hello 1, world 1

* In -> "hello hello world"  
Out -> hello 2, world 1

## Increasing Complexity:

For extra fun, the inputs to the function can have the following complications:  
* Character case must be ignored.
* Non-alphanumeric characters must be stripped.
* The above should apply for all UTF-8 characters.

Additionally, the function can be extended from a single input to multiples, with the option of spreading the workload concurrently among many goroutines.

# Running

`go run main.go helpers.go [flags]`

flags:
* `-basic`  
A synchronous version using the most basic inputs.  The default mode.
* `-complicated`  
A synchronous version using inputs with mixed case and non-alphanumeric characters.
* `-channeled`  
A concurrent version that uses channels to achieve concurrency across multiple inputs.
* `-sliced`  
A concurrent version that hands chunks of the input slice to each goroutine.
* `-lorem`  
An input set variation for use when running channeled or sliced versions.
* `-routines N`  
Count of go routines to create when running concurrent versions.  Default: 2.

# Benchmarking

Benchmark tests can be run with:  
`go test -bench .`

The benchmark tests compose a slice of 1024 strings, where each string is a random selection of words from Lorem Ipsum, between 64 and 256 words long.  Benchmarks increase in slice length from 2 to 1024.  The tests are repeated once with a single thread, and once with 256 threads, for both of the concurrency funcs.