package main

import "fmt"

func main() {
	foo := bar()
	foo() // 1
	foo() // 2
	foo() // 3
	foo() // 4

	foobar := bar()
	foobar() // 1
	foobar() // 2
	foobar() // 3
}

func bar() func() {
	num := 0
	return func() {
		num = num + 1
		fmt.Println(num)
	}
}
