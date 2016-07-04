// Copyright © 2016 "Shun Yokota" All rights reserved

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func request(hostname string, cancel chan struct{}) *http.Response {
	req, err := http.NewRequest("GET", hostname, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: NewRequest: %v\n", err)
		return nil
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: DefaultClient.Do: %v\n", err)
		return nil
	}
	/*
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", b)
			fmt.Fprintf(os.Stderr, "fetch: request: ReadAll: %v\n", err)
			//return nil
		}
	*/
	return resp
}

func main() {
	response := make(chan *http.Response, len(os.Args[1:]))
	cancel := make(chan struct{})
	for _, url := range os.Args[1:] {
		//resp, err := http.Get(url)
		go func(url string) {
			response <- request(url, cancel)
		}(url)
		//time.Sleep(1 * time.Second)
	}
	//resp, err := http.Get(url)
	resp := <-response
	go func(resp http.Response) {
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s", b)
	}(*resp)
	close(cancel)
	fmt.Println("cancel")
	b, err := ioutil.ReadAll(resp.Body)
	url := fmt.Sprintf("%s%s", resp.Request.URL.Host, resp.Request.URL.Path)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}
