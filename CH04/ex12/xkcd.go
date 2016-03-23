// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type xkcdJSON struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {
	if _, err := os.Stat("json"); err != nil {
		if err := os.Mkdir("json", os.ModePerm); err != nil {
			fmt.Println("mkdir :" + err.Error())
		}
		fmt.Println("mkdir json!!!")
	}
	latestNum := getLatestNum()
	fetchAllNewJSON(latestNum)

}

func fetchAllNewJSON(latest int) {
	fmt.Printf("fetchall :%d\n", latest)
	//ch := make(chan int, 5)
	ctl := make(chan int, 50)
	notify := make(chan int)
	var cachedNum int
	for i := 1; i <= latest; i++ {
		//go jsonCrawler(i, ch)
		if jsonIsExisting(i) {
			cachedNum++
			continue
		}
		go parallelFetch(i, ctl, notify)
	}
	for cachedNum < latest {
		cachedNum += <-notify
	}
}

func jsonCrawler(jsonNum int, ch chan<- int) {
	if jsonIsExisting(jsonNum) {
		ch <- 0
	} else {
		ch <- jsonNum
	}
}

func parallelFetch(fetchNum int, ctl chan int, notify chan<- int) {
	ctl <- 1
	if fetchNum != 0 {
		url := "https://xkcd.com/" + strconv.Itoa(fetchNum) + "/info.0.json"
		fmt.Println("fetch :" + url)
		cachexkcdJSON(fetchxkcdJSON(url))
	}
	notify <- 1
	<-ctl
}

func getLatestNum() int {
	url := "https://xkcd.com/info.0.json"
	return fetchxkcdJSON(url).Num

}

func jsonIsExisting(latest int) bool {
	path := "json/" + strconv.Itoa(latest) + ".json"
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func fetchxkcdJSON(url string) xkcdJSON {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get :" + err.Error())
	}
	var result xkcdJSON
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Println("json decode :" + err.Error())
	}
	resp.Body.Close()

	return result

}

func cachexkcdJSON(result xkcdJSON) {
	num := result.Num
	if num == 0 {
		return
	}
	path := "json/" + strconv.Itoa(num) + ".json"
	jsonData, err := json.Marshal(result)

	if err = ioutil.WriteFile(path, jsonData, 0644); err != nil {
		fmt.Println("WriteFIle :" + err.Error())
	}
	fmt.Println("cached :" + path)
	//	jsonData, err = json.MarshalIndent(result, "", " ")
	//	fmt.Printf("%s\n", jsonData)
}
