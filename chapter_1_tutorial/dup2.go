// Dup2 prints the count and text of lines that appear more than once in the input.
// It reads from stdin or from alist of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// enter "end" in the terminal will end the input cycle and break the for loop
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

// Run "go run dup2.go dup2textfile.txt" will print out these results in the terminal
// 2	a
// 2	b

// Run "go run dup2.go dup2textfile.txt dup2textfile2.txt" will print out these results in the terminal
// 3	a
// 2	b

// Learning
// A map is a reference to the data structure created by make.
// When a map is passed to a function, the function receives a copy of the reference,
// so any changes the called function makes to the underlying data structure will be
// visible through the caller’s map reference too. In our example, the values inserted
// into the counts map by countLines are seen by main.
