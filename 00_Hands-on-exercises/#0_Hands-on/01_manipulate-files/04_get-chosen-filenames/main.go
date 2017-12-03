package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	var filename string
	fmt.Print("Enter filename: ")
	fmt.Scan(&filename)

	if filename != "" {
		getFile(filename)
	}
}

func getFile(fname string) {
	files, err := filepath.Glob("*"+fname+"*.*")
	if err != nil {
		log.Fatal(err)
	}

	if files == nil {
		fmt.Println("None files match with these name.")
	}

	for _, f := range files {
		fmt.Println(f)
	}
}
