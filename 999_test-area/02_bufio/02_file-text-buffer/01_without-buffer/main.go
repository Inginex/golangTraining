package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("../br.dic.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}

	words := string(file)
	fmt.Println(words)

}
