package main

import (
	"fmt"
	"time"
)

func main() {
	var t = time.Now().Year()
	var year = 1997

	for year <= int(t) {
		fmt.Println(year)
		year++
	}

}
