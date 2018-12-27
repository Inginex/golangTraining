package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	// GENERATE A RANDOM INT
	rand.Seed(time.Now().Unix())
	num := rand.Intn(6 - 1) + 1

	fmt.Printf("%d \t %b \t %#X \n", num, num, num)

	sb := num << 1
	fmt.Printf("%d \t %b \t %#X", sb, sb, sb)
}