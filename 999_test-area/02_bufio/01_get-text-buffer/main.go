package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://github.com/thoughtworks/dadoware/blob/master/7776palavras-numeradas-2e.txt")
	if err != nil {
		log.Fatalf("Errrrrrrrrrr: %v\n")
		return
	}

	words := make(map[string]string)
	scan := bufio.NewScanner(res.Body)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()] = ""
	}

	if err := scan.Err(); err != nil {
		log.Println(os.Stderr, "reading output: ", err)
	}

	i := 0
	for text, _ := range words {
		fmt.Println(text)
		if i == 200 {
			break
		}
		i++
	}
}
