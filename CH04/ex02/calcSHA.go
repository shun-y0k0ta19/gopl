package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	opt := "256"
	if len(os.Args) > 1 {
		opt = os.Args[1]
	}
	for scanner.Scan() {
		fmt.Printf("SHA%s %x\n", opt, calcSHA(scanner.Text(), opt))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func calcSHA(s, opt string) []byte {
	switch opt {
	case "256":
		res := sha256.Sum256([]byte(s))
		return res[:]
	case "384":
		res := sha512.Sum384([]byte(s))
		return res[:]
	case "512":
		res := sha512.Sum512([]byte(s))
		return res[:]
	default:
		fmt.Println("undefined option.")
		return []byte{}
	}
}
