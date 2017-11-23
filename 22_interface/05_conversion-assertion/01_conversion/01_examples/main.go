package main

import (
	"fmt"
)

func main(){
	// conversion: int to float64
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x))

	// conversion: float64 to int
	var x1 = 12
	var y1 = 12.1230123
	fmt.Println(int(y1) + x1)

	// conversion: runes to string
	var x2 rune = 'a' // rune ia an alis for int32
	var y2 int32 = 'b'
	fmt.Println(string(x2))
	fmt.Println(string(y2))

	// conversion: []bytes to string
	fmt.Println(string([]byte{'h','e','l','l','o'}))

	// conversion: string to []bytes
	fmt.Println([]byte("hello"))
}