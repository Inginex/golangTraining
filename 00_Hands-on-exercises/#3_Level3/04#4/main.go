//Hands-on exercise #4
//Create a for loop using this syntax for {} Have it print out the years you have been alive.

package main

import (
	"fmt"
	"time"
)

func main(){
	var t = time.Now().Year()
	var year = 1997

	for {
		if year > int(t){
			break;
		}
		fmt.Println(year)
		year++
	}
}
