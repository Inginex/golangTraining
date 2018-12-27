package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	files, err := filepath.Glob("file*.*")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f)
	}
}
