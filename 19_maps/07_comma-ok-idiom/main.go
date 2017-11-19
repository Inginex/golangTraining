package main

import (
	"fmt"
)

func main(){

	wordList := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Bom dia",
		3: "Bongiorno",
	}

	fmt.Println(wordList)

	// delete(wordList, 2)

	if val, exists := wordList[2]; exists {
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	}
}