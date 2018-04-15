package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetchAll(url, ch)
		go fetchAll(url, ch)
		go fetchAll(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}
	end := time.Since(start).Seconds()
	fmt.Printf("Task finished in %.2fs.", end)

}

func fetchAll(url string, ch chan string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while reading: %s, %v\n", url, err)
		return
	}
	defer res.Body.Close()

	nbytes, err := io.Copy(os.Stdout, res.Body)
	if err != nil {
		ch <- fmt.Sprintf("while copying: %v, %v\n", url, err)
		return
	}
	end := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Fetched url: %s in %.2fs - size: %d", url, end, nbytes)
}
