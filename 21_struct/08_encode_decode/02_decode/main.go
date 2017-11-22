package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

type Person struct {
	Name string
	Last string
	Age int
}

func main(){
	var p1 Person
	rdr := strings.NewReader(`{"Name":"James","Last":"Bond","Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1)
}