package main

import (
	"os"
	"encoding/json"
)

type Person struct {
	First string
	Second string
	Age int
}

func main(){
	p1 := Person{"James", "Bond", 20}
	json.NewEncoder(os.Stdout).Encode(p1)
}