package main 

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First string
	Last string
	Age int
}

func main(){
	var p1 Person
	
	sb := []byte(`{"First":"James", "Last":"Bond", "Age": 30}`)
	json.Unmarshal(sb, &p1)
	
	fmt.Println(sb)
}