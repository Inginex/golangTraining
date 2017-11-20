package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	first string
	second string
	age int
}

func main(){
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
}