package main

import (
	"fmt"
)

func main(){
	golang := struct {
		name string
		year int
		creators []string
	}{
		name: "Go",
		year: 2009,
		creators: []string{
			"Rob Pike", 
			"Ken Thompson",
		    "Robert Griesemer",
		},
	}

	fmt.Printf("Language: %v, Birthday: %v\n",golang.name, golang.year)
	for i, v := range golang.creators {
		fmt.Printf("%v - %v \t", i, v)
	}
}