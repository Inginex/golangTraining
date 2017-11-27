package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"https://www.youtube.com/",
		"https://www.amazon.com/",
		"https://www.kabum.com.br/",
		"https://www.facebook.com/",
		"https://www.google.com.br/",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			bs, err := ioutil.ReadAll(res.Body)
			defer res.Body.Close()

			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}

		}(url)
	}

	var bigger HomePageSize

	for range urls {
		result := <- results
		if result.Size > bigger.Size {
			bigger = result
		}			
		fmt.Println("URL: ",result.URL,"\t Size: ", result.Size)
	}

	fmt.Println("The biggest site: ", bigger.Size)

}
