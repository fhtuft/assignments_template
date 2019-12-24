package cos418_hw1_1

import (
    "fmt"
	"bufio"
	"io"
	"strconv"
    "os"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	// TODO: implement me
	// HINT: use for loop over `nums`
    sum := 0
    for num := range nums {
        //fmt.Printf("num : %d\n", num)
        sum += num
    }
    out <- sum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// TODO: implement me
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers
    fmt.Printf("num : %d fileName : %s\n", num, fileName)
    file , err := os.Open(fileName)
    if err != nil {
        panic("could not open file") // Right to panic her ? 
    }
    elems, err := readInts(file)
    if err != nil {
        panic("failed reading int's")
    }
    // Create the need channels
    nums := make(chan int)
    out  := make(chan int)
    // Create gorutines
    for i := 0 ; i < num ; i++ {
        go sumWorker(nums, out)
    }
    // Send the data, one element at a time..
    for _, e := range elems {
        nums <- e
    }
    close(nums) // signals end of data
    // Sum the remote computed sums
    sum := 0
    for i := 0; i < num; i++ {
        sum  += <-out
    }

	return sum
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
        //fmt.Printf("val : %d\n", val)
		elems = append(elems, val)
	}
	return elems, nil
}
