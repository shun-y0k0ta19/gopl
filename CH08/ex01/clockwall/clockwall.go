// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

//var times map[string]string

func main() {
	//times := make(map[string]string)
	for _, arg := range os.Args[1:] {
		timezoneAndURL := strings.Split(arg, "=")
		timezone := timezoneAndURL[0]
		url := timezoneAndURL[1]
		fmt.Printf("arg: %s\n", arg)
		fmt.Printf("timeZone: %s\n", timezone)
		fmt.Printf("url: %s\n", url)
		conn, err := net.Dial("tcp", url)
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()
		go printTime(conn, timezone)
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if sc.Text() == "q" {
			return
		}
	}

}

func printTime(conn net.Conn, timezone string) {
	connsc := bufio.NewScanner(conn)
	for connsc.Scan() {
		fmt.Printf("%s=%v\n", timezone, connsc.Text())
	}
}
