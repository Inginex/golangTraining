package main

import (
	"sort"
	"log"
	"os"
)

var country = []string{"Brasil", "Argentina", "Colombia", "Uruguai", "Guatemala"}

func main() {
	cfile, err := os.Create("paises.txt")
	if err != nil {
		log.Fatal("err creating file")
	}
	defer cfile.Close()

	writeFile()
}

func writeFile() {
	file, err := os.OpenFile("paises.txt", os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sort.Sort(sort.Reverse(sort.StringSlice(country)))

	for _, c := range country {
		file.WriteString(c + "\r\n")
	}

	file.Sync()

}
