package main

import (
	"fmt"
)

func main(){
	fName := "Mar"

	switch {
	case len(fName) == 3:
		fmt.Println("Wassup my firend witch name of length 3")
	case fName == "Tim":
		fmt.Println("Wassup Tim")
	case fName == "Jenny", fName == "Marcus":
		fmt.Println("Your name is either Marcus or Jenny")
	default:
		fmt.Println("Nothing matched; this is the default")
	}
}