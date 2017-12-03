package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := "./file02.txt"
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted file: ", strings.Split(path, "/")[1])
}
