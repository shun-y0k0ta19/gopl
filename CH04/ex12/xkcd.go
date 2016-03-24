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
	"strings"
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
	//latestNum := getLatestNum()
	//fetchJSONs(1, latestNum)

	args := os.Args[1:]
	for _, index := range args {
		if strings.Contains(index, "-") {
			nums := strings.Split(index, "-")
			begin, err := strconv.Atoi(nums[0])
			if err != nil {
				fmt.Println("begin Atoi :" + err.Error())
			}
			end, err := strconv.Atoi(nums[1])
			if err != nil {
				fmt.Println("end Atoi :" + err.Error())
			}
			if begin > end {
				begin, end = end, begin
			}
			fetchJSONs(begin, end)
			for i := begin; i <= end; i++ {
				showURLandTranscript(i)
			}
		} else {
			num, err := strconv.Atoi(index)
			if err != nil {
				fmt.Println("Atoi :" + err.Error())
			}
			fetchJSONs(num, num)
			showURLandTranscript(num)
		}
	}

}

func showURLandTranscript(num int) {
	index := strconv.Itoa(num)
	jsonData, err := ioutil.ReadFile("json/" + index + ".json")
	if err != nil {
		fmt.Println("JSON ReadFile :" + err.Error())
	}
	var result xkcdJSON
	if err := json.Unmarshal(jsonData, &result); err != nil {
		fmt.Println("Unmarshal :" + err.Error())
	}
	fmt.Printf("\n----------------------------------\nURL :https://xkcd.com/%d/info.0.json\ntranscript :%s\n", result.Num, result.Transcript)

}

func fetchJSONs(begin, end int) {
	fmt.Printf("fetchall :%d\n", end)
	pool := make(chan int, 50)
	notify := make(chan int)
	var cachedNum int
	for i := begin; i <= end; i++ {
		if jsonIsExisting(i) {
			cachedNum++
			continue
		}
		go parallelFetch(i, pool, notify)
	}
	for cachedNum <= end-begin {
		cachedNum += <-notify
	}
}

func parallelFetch(fetchNum int, pool chan int, notify chan<- int) {
	pool <- 1
	if fetchNum != 0 {
		url := "https://xkcd.com/" + strconv.Itoa(fetchNum) + "/info.0.json"
		fmt.Println("fetch :" + url)
		cachexkcdJSON(fetchxkcdJSON(url))
	}
	notify <- 1
	<-pool
}

func getLatestNum() int {
	req := "https://xkcd.com/info.0.json"
	return fetchxkcdJSON(req).Num
}

func jsonIsExisting(latest int) bool {
	path := "json/" + strconv.Itoa(latest) + ".json"
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func fetchxkcdJSON(req string) xkcdJSON {
	resp, err := http.Get(req)
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
}
