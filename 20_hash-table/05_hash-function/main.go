package main

import (
	"fmt"
)

func main(){
	s := HashBucket("Go", 12)
	fmt.Println(s)
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}