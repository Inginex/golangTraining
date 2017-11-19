package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 122; i++ {
		// number	-	letter	-	bucket
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}
