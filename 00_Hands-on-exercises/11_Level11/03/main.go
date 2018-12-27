package main

import "fmt"

type customErr struct {
	Info string
	cerr error
}

func (err customErr) Error() string {
	return fmt.Sprintf("Info error: %s", err.Info)
}

func main() {
	cErr := customErr{"info string", fmt.Errorf("This is a custom error")}
	foo(cErr)
}

func foo(err error) {
	fmt.Println(err.(customErr).cerr)
	fmt.Println(err)
}
