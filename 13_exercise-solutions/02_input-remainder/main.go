package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&num1)
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&num2)
	fmt.Println("The remainder between the number is: ",  num2 / num1)

}
