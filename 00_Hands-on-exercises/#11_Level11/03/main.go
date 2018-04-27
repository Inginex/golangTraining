package main

import "fmt"

type customErr struct {
	error
}

func main() {
	cErr := customErr{fmt.Errorf("This is a custom error")}
	foo(cErr)
}

func foo(err error) {
	fmt.Println(err)
}
