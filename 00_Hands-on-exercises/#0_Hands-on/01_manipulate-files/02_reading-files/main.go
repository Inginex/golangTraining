package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// EXAMPLE 01 - TO LARGE FILES
	// file, err := os.OpenFile("../01_writing-files/paises.txt", os.O_RDONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// str := make([]byte, 1024)
	// file.Read(str)

	// EXAMPLE 02 - TO SMALL FILES
	str, err := ioutil.ReadFile("../01_writing-files/paises.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(str))
}
