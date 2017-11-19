package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(string(bs))
	fmt.Printf("%s", bs)
}