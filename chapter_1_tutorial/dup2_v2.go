// Dup2 prints the count and text of lines that appear more than once in the input.
// It reads from stdin or from a list of named files; and then pritn the names of all files in which each duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineFileMap := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, nil)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, lineFileMap)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, lineFileMap[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineFileMap map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// enter "end" in the terminal will end the input cycle and break the for loop
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++

		lineFileMap[input.Text()] = append(lineFileMap[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
