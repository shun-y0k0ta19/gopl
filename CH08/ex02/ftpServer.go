// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
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
)

var user string

type close struct{}

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
	switch user {
	case "root":
		if pass != "root" {
			fmt.Fprintln(conn, 530)
			return false
		}
	default:
	}
	fmt.Fprintf(conn, "%d User %s logged in.\r\n", 230, user)
	return true
}

func userAuth(commandSc *bufio.Scanner, conn net.Conn) bool {
	if commandSc.Scan() {
		user = commandSc.Text()
		fmt.Println(user)
		switch user {
		case "anonymous", "root":
			fmt.Fprintf(conn, "%d Please specify the password.\r\n", 331)
			return true
		default:
			fmt.Fprintln(conn, 530)
			return false
		}
	}
	return false
}

func cdPath(path string, conn net.Conn) {
	err := os.Chdir(path)
	fmt.Println(err)
	if err != nil {
		fmt.Fprintf(conn, "%d Failed to change directory.\r\n", 550)
		return
	}
	fmt.Fprintf(conn, "%d Directory successfully changed.\r\n", 250)
}

func cd(commandSc *bufio.Scanner, conn net.Conn) {
	if commandSc.Scan() {
		path := commandSc.Text()
		fmt.Println(path)
		cdPath(path, conn)
	}
}

func mkdir(commandSc *bufio.Scanner, conn net.Conn) {
	for commandSc.Scan() {
		path := commandSc.Text()
		fmt.Println(path)
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Fprintf(conn, "%d Failed to make directory \"%s\".\r\n", 550, path)
			continue
		}
		//fmt.Fprintf(conn, "%d \"%s\" \r\n", 257, pwd())

		fmt.Fprintf(conn, "%d \"%s\" successfully made.\r\n", 257, path)
	}
}

func get(commandSc *bufio.Scanner, conn net.Conn, msg chan interface{}) {
	if commandSc.Scan() {
		path := commandSc.Text()
		fmt.Println(path)
		fp, err := os.Open(path)
		fmt.Println(err)
		if err != nil {
			fmt.Fprintf(conn, "%d Failed to open file(%s).\r\n", 550, err.Error())
			return
		}
		fmt.Fprintf(conn, "%d Get %s.----------->\r\n", 150, path)
		sc := bufio.NewScanner(fp)
		for sc.Scan() {
			msg <- fmt.Sprintf("%s\r\n", sc.Text())
		}
	}
	msg <- close{}
	fmt.Fprintf(conn, "%d Closing data connection.\r\n", 226)

}

func handleDconn(dconn net.Conn, msg chan interface{}) {
	defer dconn.Close()
	fmt.Println("start dconn")
	for {
		o := <-msg
		switch data := o.(type) {
		case string:
			fmt.Fprintf(dconn, "%s", data)
		case close:
			return
		}
	}
}

func listenDataConn(port int) (net.Listener, error) {
	dataConnURL := fmt.Sprintf("localhost:%d", port)
	return net.Listen("tcp", dataConnURL)
}

func wordsScanner(s string) *bufio.Scanner {
	commandSc := bufio.NewScanner(strings.NewReader(s))
	commandSc.Split(bufio.ScanWords)
	return commandSc
}

func listAll(entry os.FileInfo) interface{} {
	modT := entry.ModTime()
	dt := fmt.Sprintf("%02d %02d %02d:%02d", modT.Month(), modT.Day(), modT.Hour(), modT.Minute())
	//fmt.Printf("%s %10d %s %s\n", data.Mode(), data.Size(), dt, data.Name())
	return fmt.Sprintf("%s %10d %s %s\r\n", entry.Mode(), entry.Size(), dt, entry.Name())
}

func listName(entry os.FileInfo) interface{} {
	return fmt.Sprintf("%s\r\n", entry.Name())
}

func list(csc *bufio.Scanner, conn net.Conn, msg chan interface{}, elist func(os.FileInfo) interface{}) {
	path := pwd()
	if csc.Scan() {
		t := csc.Text()
		fmt.Println(t)
		path = fmt.Sprintf("%s/%s", path, t)
	}
	fmt.Println(path)
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(conn, "%d Requested file action not taken.\r\n", 450)
		msg <- close{}
		return
	}
	fmt.Fprintf(conn, "%d Here comes the directory listing.\r\n", 150)

	for _, entry := range entries {
		msg <- elist(entry)
	}
	msg <- close{}
	fmt.Fprintf(conn, "%d Closing data connection.\r\n", 226)
}

//!+handleConn
func handleConn(conn net.Conn) {
	msg := make(chan interface{})

	random := rand.New(rand.NewSource(int64(time.Now().Second())))
	port := random.Intn(0xFFFF-1024) + 1024
	p1 := port >> 8
	p2 := port & 0xFF
	data, err := listenDataConn(port)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("start scan")
	fmt.Fprintf(conn, "%d Golang FTP server Ready.\r\n", 220)
	input := bufio.NewScanner(conn)

ftpCon:
	for input.Scan() {
		s := input.Text()
		fmt.Println(s)
		fmt.Println([]byte(s))
		commandSc := wordsScanner(s)
		commandSc.Scan()
		command := commandSc.Text()

		fmt.Println(command)
		switch command {
		case "USER":
			if !userAuth(commandSc, conn) {
				break ftpCon
			}
		case "PASS":
			if !passAuth(s[5:], conn) {
				break ftpCon
			}
		case "SYST":
			fmt.Fprintf(conn, "%d UNIX\r\n", 215)
		case "FEAT":
			fmt.Fprintf(conn, "%d no features\r\n", 211)
		case "PWD":
			//fmt.Println("PWDda")
			fmt.Fprintf(conn, "%d \"%s\" \r\n", 257, pwd())
			//fmt.Fprintf(conn, "%d\n", 250)
		case "CWD":
			cd(commandSc, conn)
		case "CDUP":
			cdPath("..", conn)
		case "MKD":
			if user != "root" {
				fmt.Fprintf(conn, "%d Parmission denied.\r\n", 550)
			} else {
				mkdir(commandSc, conn)
			}
		case "EPSV":
			fmt.Fprintf(conn, "%d EPSV is not supported.\r\n", 500)
			//ls(commandSc, conn)
		case "PASV":
			go func() {
				dconn, err := data.Accept()
				if err != nil {
					log.Print(err)
				}
				handleDconn(dconn, msg)
			}()
			fmt.Fprintf(conn, "%d Entering Passive Mode %d,%d,%d,%d,%d,%d\r\n", 227, a1, a2, a3, a4, p1, p2)
		case "LIST":
			list(commandSc, conn, msg, listAll)
		case "NLST":
			list(commandSc, conn, msg, listName)
		case "RETR":
			get(commandSc, conn, msg)
		case "QUIT":
			fmt.Fprintf(conn, "%d See you again!\r\n", 221)
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

func main() {
	control, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := control.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
