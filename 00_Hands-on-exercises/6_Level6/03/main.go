package main

import "fmt"

func main() {
	defer func() {
		defer func() {
			// This will be the last func executed
			fmt.Println(" bar func DEFER !!!!!")
		}()
		// This will be the second
		fmt.Println(" foo func DEFER !!!!!")
	}()
	// This executs first
	fmt.Println(" RUN FUNC FOOOOOO BAAAAAAAAR")
}
