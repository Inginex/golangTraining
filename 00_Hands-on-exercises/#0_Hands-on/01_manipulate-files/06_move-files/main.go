package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Rename("./file02.txt", "../05_delete-files/file02.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File moved with success!")
}
