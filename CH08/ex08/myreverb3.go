// Copyright © 2016 "Shun Yokota" All rights reserved

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	countdown := make(chan int, 1)
	countdown <- 0
	tick := time.Tick(1 * time.Second)

	go func() {
		for input.Scan() {
			<-countdown
			countdown <- 0
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(c, input.Text(), 1*time.Second)
			}()
		}
	}()
	for {
		select {
		case <-tick:
			cd := <-countdown
			if cd >= 10 {
				wg.Wait()
				fmt.Fprintln(c, "timeout!! disconnected.")
				c.Close()
				return
			}
			countdown <- cd + 1
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
