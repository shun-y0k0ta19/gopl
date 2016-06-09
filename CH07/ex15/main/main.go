// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"golang_training/CH07/ex15"
	"strconv"
	//"../" //GOPATH下にないときのimportパス

	"bufio"
	"log"
	"os"
)

func main() {
	fScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("input formula: ")
	for fScanner.Scan() {
		input := fScanner.Text()
		if len(input) < 1 {
			continue
		}
		//fmt.Println(input)
		e, err := eval.Parse(input)
		if err != nil {
			log.Fatal(fmt.Errorf("parse error\n%s", err))
		}
		//		vars := eval.ExtractVars(e)
		vars := make(map[eval.Var]bool)
		if err := e.Check(vars); err != nil {
			log.Fatal(err)
		}
		env := make(eval.Env)
		if len(vars) > 0 {
			fmt.Println("\ninput value:")
			for v := range vars {
				fmt.Printf("%s: ", v)
				if fScanner.Scan() {
					s := fScanner.Text()
					env[v], err = strconv.ParseFloat(s, 64)
					if err != nil {
						log.Fatal(fmt.Errorf("this value is invalid.\n%s", err))
					}
				}
			}
		}
		fmt.Println(e.Eval(env))
		return
	}
}
