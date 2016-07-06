// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

const (
	a1 = 127
	a2 = 0
	a3 = 0
	a4 = 1
	p1 = 4
	p2 = 255
)

var user string

func pwd() string {
	curDir, err := os.Getwd()
	//fmt.Println(curDir)
	if err != nil {
		log.Fatalln(err)
	}
	return curDir
}

func passAuth(pass string, conn net.Conn) bool {
	fmt.Println("in pathAuth")
	fmt.Println(pass)
	fmt.Fprintf(conn, "%d User %s logged in.\n", 230, user)
	return true
}

func userAuth(commandSc *bufio.Scanner, conn net.Conn) bool {
	if commandSc.Scan() {
		user = commandSc.Text()
		fmt.Println(user)
		if user != "root" {
			fmt.Fprintln(conn, 530)
			return false
		}
		fmt.Fprintf(conn, "%d Please specify the password.\n", 331)
		return true
	}
	return false
}

func cd(commandSc *bufio.Scanner, conn net.Conn) {
	if commandSc.Scan() {
		path := commandSc.Text()
		fmt.Println(path)
		err := os.Chdir(path)
		fmt.Println(err)
		if err != nil {
			fmt.Fprintf(conn, "%d Failed to change directory.\n", 550)
			return
		}
		fmt.Fprintf(conn, "%d Directory successfully changed.\n", 250)
	}

}

func handleDconn(dconn net.Conn, msg chan interface{}) {
	fmt.Println("start dconn")
	for {
		o := <-msg
		switch data := o.(type) {
		case os.FileInfo:
			fmt.Println("entry print")
			fmt.Fprintf(dconn, "%s\n", data.Name())
		}
	}
}

//!+handleConn
func handleConn(conn net.Conn) {
	msg := make(chan interface{})

	//fmt.Println("start scan")
	fmt.Fprintf(conn, "%d Golang FTP server Ready.\n", 220)
	input := bufio.NewScanner(conn)
ftpCon:
	for input.Scan() {
		s := input.Text()
		fmt.Println(s)
		fmt.Println([]byte(s))
		commandSc := bufio.NewScanner(strings.NewReader(s))
		commandSc.Split(bufio.ScanWords)
		commandSc.Scan()
		command := commandSc.Text()

		fmt.Println(command)
		switch command {
		case "USER":
			if !userAuth(commandSc, conn) {
				break ftpCon
			}
		case "PASS":
			if !passAuth(s[4:], conn) {
				break ftpCon
			}
		case "SYST":
			fmt.Fprintf(conn, "%d UNIX\n", 215)
		case "FEAT":
			fmt.Fprintf(conn, "%d no features\n", 211)
		case "PWD":
			//fmt.Println("PWDda")
			fmt.Fprintf(conn, "%d \"%s\" \n", 257, pwd())
			//fmt.Fprintf(conn, "%d\n", 250)
		case "CWD":
			cd(commandSc, conn)
		case "EPSV":
			fmt.Fprintf(conn, "%d EPSV is not supported.\n", 500)
			//ls(commandSc, conn)
		case "PASV":
			ch := make(chan struct{})
			go func() {
				data, err := net.Listen("tcp", "localhost:1279")
				if err != nil {
					log.Fatal(err)
				}
				ch <- struct{}{}
				for {
					dconn, err := data.Accept()
					if err != nil {
						log.Print(err)
						continue
					}
					go handleDconn(dconn, msg)
				}
			}()
			<-ch
			fmt.Fprintf(conn, "%d Entering Passive Mode %d,%d,%d,%d,%d,%d\n", 227, a1, a2, a3, a4, p1, p2)
		case "LIST":
			entries, err := ioutil.ReadDir(pwd())
			if err != nil {
				fmt.Fprintln(conn, 450)
			}
			fmt.Fprintln(conn, 150)

			for _, entry := range entries {
				msg <- entry
			}

		case "PORT":

		case "QUIT":
			fmt.Fprintf(conn, "%d See you again!\n", 221)
			break ftpCon
		default:
			fmt.Println("default")
			fmt.Fprintln(conn, 500)
		}
	}
	//io.Copy(os.Stdout, conn) // NOTE: ignoring errors
	time.Sleep(2 * time.Second)
	//done <- struct{}{} // signal the main goroutine
	log.Println("done")

	// NOTE: ignoring potential errors from input.Err()

	conn.Close()
}

//!-handleConn

//!+main
func main() {
	control, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}
	/*
		data, err := net.Listen("tcp", "localhost:8010")
		if err != nil {
			log.Fatal(err)
		}
	*/
	//go broadcaster()
	for {
		conn, err := control.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
		/*
			dconn, err := data.Accept()
			if err != nil {
				log.Print(err)
				continue
			}
			go handleConn(dconn)
		*/
	}
}

//!-main
