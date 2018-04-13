package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		fetch(url, os.Stdout)
	}
}

func fetch(url string, out *os.File) {
	if ok := strings.HasPrefix(url, "https://"); !ok {
		url = fmt.Sprint("https://" + url)
	}

	start := time.Now()
	rs, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurs on: %s, %v\n", url, err)
		return
	}
	defer rs.Body.Close()
	// _, err = ioutil.ReadAll(rs.Body)
	// if err != nil {
	// 	fmt.Printf("Error occurs on: %s, %v\n", url, err)
	// 	return
	// }

	_, err = io.Copy(out, rs.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurs on: %s, %v\n", url, err)
		return
	}
	status := rs.Status
	finish := time.Since(start).Seconds()
	fmt.Fprintf(os.Stdout, "\n ::::: Fetched in %.2fs, url: %s, StatusCode: %s :::::\n", finish, url, status)
}
