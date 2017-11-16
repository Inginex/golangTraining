package main

import (
	"fmt"
)

func main(){
	for i := 1; i < 5; i++ {
		for j := 10; j > 1; j-- {
			if i < 3 {
				fmt.Println("i: ",i + 2)
			} else {
				fmt.Println("i: ",i - 1)
			}
			if j > 5 {
				fmt.Println("j: ",j - 2)
			} else {
				fmt.Println("j: ",j + 1)
			}
		}
	}
}