package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bs, err := ioutil.ReadFile("../br.dic.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Create slice to hold counts
	buckets := make(map[string]int, 200)

	for _, byte := range bs {
		// Filters the letters to A to Z
		if byte > 64 && byte < 123 {
			n := hashBucket(string(byte))
			buckets[string(byte)] = n + 1
		}
	}

	fmt.Println(buckets)
}

func hashBucket(word string) int {
	return int(word[0])
}
