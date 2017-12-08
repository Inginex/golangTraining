package main

import (
	"fmt"
)


func main(){
	ps := map[string][]string{
		"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr": []string{`Ice cream`, `Sunsets`},
	}

	for i, v := range ps {
		fmt.Println("Range for index:", i)
		for j, u := range v {
			fmt.Printf("Index: %v \t Value: %v\n", j, u)
		}
	}
}