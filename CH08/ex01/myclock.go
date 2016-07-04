// Copyright © 2016 "Shun Yokota" All rights reserved

// Clock is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type PortNum string

func (pn *PortNum) Set(s string) error {
	if _, err := strconv.Atoi(s); err != nil {
		return fmt.Errorf("invalid port number: %s", s)
	}
	*pn = PortNum(s)
	return nil
}

func (pn *PortNum) String() string {
	return string(*pn)
}

// PortFlag defines a portNum flag with the specified name,
// default value, and usage
func PortFlag(name string, value string, usage string) *PortNum {
	f := PortNum(value)
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	portNum := PortFlag("port", "8000", "port number")
	flag.Parse()
	url := fmt.Sprintf("localhost:%s", portNum)
	fmt.Printf("url: %s\n", url)
	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
