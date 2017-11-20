package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Second string
	Age int
	notExported int
}

func main(){
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
}