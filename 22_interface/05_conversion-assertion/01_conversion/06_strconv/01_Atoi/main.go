package main 

import (
	"fmt"
	"strconv"
)

func main(){
	// conversion: Atoi(string to int)
	i, _ := strconv.Atoi("-12")
	fmt.Println(i)
	fmt.Printf("%T", i)
}