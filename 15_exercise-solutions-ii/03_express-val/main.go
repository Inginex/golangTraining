package main

import (
	"fmt"
)

func main() {
	ex := (true && false) || (false && true) || !(false && false)
	fmt.Println(ex)
}
