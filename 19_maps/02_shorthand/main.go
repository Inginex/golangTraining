package main

import (
	"fmt"
)

func main(){
	words := map[string]string{
		"word1": 	"Hello",
		"word2": 	"World",
	}

	words["word3"] = "Hello again"
	delete(words, "word3")

	fmt.Println(words["word3"])
}