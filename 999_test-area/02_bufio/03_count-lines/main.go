package main

import (
	"bufio"
	"fmt"
	"os"
)

var counts = map[string]int{}

func main() {
	files := os.Args[1:]
	// Check if any argument was passed
	if len(files) == 0 {
		// If not use Stdin to get the lines
		countLine(os.Stdin)
	} else {
		// If so range over the files and open all of them
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "File err: %v", err)
				continue
			}
			defer f.Close()
			countLine(f)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("Text: %v - Counting: %v\n", line, n)
		}
	}
}

func countLine(f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
