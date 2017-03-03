// Homework 5: Goroutines
// Due February 28, 2017 at 11:59pm
package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Feel free to use the main function for testing your functions
	tasks := []func() (string, error){
		func() (string, error) {
			time.Sleep(time.Second)
			fmt.Println(1)
			return "hello", nil
		},
		func() (string, error) {
			time.Sleep(time.Second)
			fmt.Println(2)
			return "world", nil
		},
	}
	ch := ConcurrentRetry(tasks, 4, 2)
	fmt.Println(<-ch)
}

// Filter copies values from the input channel into an output channel that match the filter function p
// The function p determines whether an int from the input channel c is sent on the output channel
func Filter(c <-chan int, p func(int) bool) <-chan int {
	result := make(chan int, len(c))
	defer close(result)
	for i := range c {
		if p(i) {
			result <- i
		}
	}
	return result
}

// Result is a type representing a single result with its index from a slice
type Result struct {
	index  int
	result string
}

// ConcurrentRetry runs all the tasks concurrently and sends the output in a Result channel
//
// concurrent is the limit on the number of tasks running in parallel. Your
// solution must not run more than `concurrent` number of tasks in parallel.
//
// retry is the number of times that the task should be attempted. If a task
// returns an error, the function should be retried immediately up to `retry`
// times. Only send the results of a task into the output channel if it does not error.
//
// Multiple instances of ConcurrentRetry should be able to run simultaneously
// without interfering with one another, so global variables should not be used.
// The function must return the channel without waiting for the tasks to
// execute, and all results should be sent on the output channel. Once all tasks
// have been completed, close the channel.
func ConcurrentRetry(tasks []func() (string, error), concurrent int, retry int) <-chan Result {
	result := make(chan Result, concurrent)
	var wg sync.WaitGroup
	for i, t := range tasks {
		go func() {
			wg.Add(1)
			for j := 0; j < retry; j++ {
				str, err := t()
				if err == nil {
					add := Result{index: i, result: str}
					result <- add
					break
				}
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	fmt.Println("Done") // This print statement is left in the code because, for an undetermined reason, the function does not work without it
	return result
}

// Task is an interface for types that process integers
type Task interface {
	Execute(int) (int, error)
}

// Fastest returns the result of the fastest running task
// Fastest accepts any number of Task structs. If no tasks are submitted to
// Fastest(), it should return an error.
// You should return the result of a Task even if it errors.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
type placer struct {
	i int
	e error
}

func Fastest(input int, tasks ...Task) (int, error) {
	if len(tasks) == 0 {
		return 0, errors.New("No Tasks submitted")
	}
	result := make(chan placer, len(tasks))
	for _, t := range tasks {
		go func() {
			in, er := t.Execute(input)
			result <- placer{i: in, e: er}
		}()
	}
	fastest := <-result
	return fastest.i, fastest.e
}

// MapReduce takes any number of tasks, and feeds their results through reduce
// If no tasks are supplied, return an error.
// If any of the tasks error during their execution, return an error immediately.
// Once all tasks have completed successfully, return the value of reduce on
// their results in any order.
// Do not leave any pending goroutines. Make sure all goroutines are cleaned up
// properly and any synchronizing mechanisms closed.
func MapReduce(input int, reduce func(results []int) int, tasks ...Task) (int, error) {
	if len(tasks) == 0 {
		return 0, errors.New("No Tasks submitted")
	}
	resultChan := make(chan placer, len(tasks))
	for _, t := range tasks {
		go func() {
			in, er := t.Execute(input)
			add := placer{i: in, e: er}
			resultChan <- add
		}()
	}
	resultSlice := make([]int, len(tasks))
	for result := range resultChan {
		if result.e != nil {
			return 0, errors.New("An error occurred")
		} else {
			resultSlice = append(resultSlice, result.i)
		}
	}
	close(resultChan)
	return reduce(resultSlice), nil
}
