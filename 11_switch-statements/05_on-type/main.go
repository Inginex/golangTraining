package main

import (
	"fmt"
)

type Contact struct {
	greeting string
	name string
}

func SwitchOnType(x interface{}){
	switch x. (type){
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	case bool:
		fmt.Println("boolean")
	default:
		fmt.Println("unknown")
	}
}

func main(){
	SwitchOnType(7)
}