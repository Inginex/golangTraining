package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open file
	file, err := os.Open("../br.dic.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer file.Close()

	var wordlist []byte

	// Define buffer size
	bf := make([]byte, 32*1024)

	// Make a loop to read through the file
	for {
		// Get the file length
		n, err := file.Read(bf)

		// Get all the words found
		if n > 0 {
			wordlist = append(wordlist, bf[:n]...)
		}
		// Break loop if reachs the end of the length
		if err == io.EOF {
			break
		}
		// Break the loop if any error occurs
		if err != nil {
			log.Printf("Read %d at bytes: %v", n, err)
			break
		}
	}

	// Create a file
	newFile, err := os.Create("wordlist.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer newFile.Close()

	n, err := newFile.Write(wordlist)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)
}
