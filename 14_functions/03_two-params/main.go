package main

import (
	"fmt"
)

func main() {
	greet("Sam", "Palm")
}

func greet(fname, lname string) {
	fmt.Println(fname + " - " + lname)
}
