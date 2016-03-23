// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	var lessthanOneYear []string
	var lessthanOneMonth []string
	var morethanOneYear []string
	now := time.Now()
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		log := fmt.Sprintf("%v\n#%-5d %9.9s %.55s\n",
			item.CreatedAt.Format("2006-01-02"), item.Number, item.User.Login, item.Title)

		addy := item.CreatedAt.AddDate(1, 0, 0)
		addm := item.CreatedAt.AddDate(0, 1, 0)
		if now.After(addy) {
			morethanOneYear = append(morethanOneYear, log)
		} else if now.Before(addm) {
			lessthanOneMonth = append(lessthanOneMonth, log)
		} else {
			lessthanOneYear = append(lessthanOneYear, log)
		}
	}

	fmt.Println("----------------more than a year ago----------------")
	for _, log := range morethanOneYear {
		fmt.Print(log)
	}
	fmt.Println("\n----------------less than a year----------------")
	for _, log := range lessthanOneYear {
		fmt.Print(log)
	}
	fmt.Println("\n----------------less than a month----------------")
	for _, log := range lessthanOneMonth {
		fmt.Print(log)
	}
}
