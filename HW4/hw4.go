// Homework 4: Concurrency
// Due February 21, 2017 at 11:59pm
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	// Feel free to use the main function for testing your functions
	in, _ := os.Open("file_sum.txt")
	out, _ := os.Create("output.txt")
	IOSum(in, out)
	in.Close()
	out.Close()
	fmt.Println("Run")
}

// Problem 1a: File processing
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// FileSum sums the integers in input and writes them to an output file.
// The two parameters, input and output, are the filenames of those files.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func FileSum(input, output string) {
	in, errIn := os.Open(input)
	out, errOut := os.Create(output)
	if errIn != nil {
		log.Fatal(errIn)
	}
	if errOut != nil {
		log.Fatal(errOut)
	}
	defer in.Close()
	defer out.Close()

	total := 0
	r := bufio.NewScanner(in)
	for r.Scan() {
		line := r.Text()
		num, _ := strconv.Atoi(line)
		total += num
	}
	if err := r.Err(); err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(out)
	w.WriteString(strconv.Itoa(total) + "\n")
	w.Flush()
}

// Problem 1b: IO processing with interfaces
// You must do the exact same task as above, but instead of being passed 2
// filenames, you are passed 2 interfaces: io.Reader and io.Writer.
// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.

// IOSum sums the integers in input and writes them to output
// The two parameters, input and output, are interfaces for io.Reader and
// io.Writer. The type signatures for these interfaces is in the Go
// documentation.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func IOSum(input io.Reader, output io.Writer) {
	total := 0
	r := bufio.NewScanner(input)
	for r.Scan() {
		line := r.Text()
		num, _ := strconv.Atoi(line)
		total += num
	}
	if err := r.Err(); err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(output)
	w.WriteString(strconv.Itoa(total) + "\n")
	w.Flush()
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// For this assignment, you will be building a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// PennDirectory is a mapping from PennID number to PennKey (12345678 -> adelq).
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex, see lecture 2 for a review on embedding
type PennDirectory struct {
	sync.RWMutex
	directory map[int]string
}

// Add inserts a new student to the Penn Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *PennDirectory) Add(id int, name string) {
	if _, ok := d.directory[id]; ok {
		log.Panicln("Warning: id already exists in directory")
	} else {
		d.Lock()
		d.directory[id] = name
		d.Unlock()
	}
}

// Get fetches a student from the Penn Directory by their PennID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *PennDirectory) Get(id int) string {
	if _, ok := d.directory[id]; !ok {
		log.Panicln("Warning: id does not exist in directory")
		return ""
	} else {
		d.RLock()
		result := d.directory[id]
		d.RUnlock()
		return result
	}
}

// Remove deletes a student to the Penn Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *PennDirectory) Remove(id int) {
	if _, ok := d.directory[id]; !ok {
		log.Panicln("Warning: id does not exist in directory")
	} else {
		d.Lock()
		delete(d.directory, id)
		d.Unlock()
	}
}
