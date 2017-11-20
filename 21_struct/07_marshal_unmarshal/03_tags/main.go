package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Second string 		`json:"-"`
	Age int				`json:"wisdom score"`
}

func main(){
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs))
}