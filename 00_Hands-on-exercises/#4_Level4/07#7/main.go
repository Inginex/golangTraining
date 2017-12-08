package main

import (
	"fmt"
)

func main(){
	p1 := []string{"James", "Bond", "Shaken, not stirred"}
	p2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	ps := [][]string{p1, p2}

	for i, v := range ps {
		fmt.Printf("Index: %v \t Value: %v\n", i, v)
	}

}