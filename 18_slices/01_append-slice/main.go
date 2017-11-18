package main

import (
	"fmt"
)

func main(){
	sliceOne := []string{"Monday", "Tuesday"}
	sliceTwo := []string{"Wednesday", "Thursday", "Friday", "Saturday"}

	sliceOne = append(sliceOne, sliceTwo...)
	fmt.Println(sliceOne)

	sliceOne = append(sliceOne[:1], sliceOne[3:]...)
	fmt.Println(sliceOne)
}