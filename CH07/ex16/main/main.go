// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"gopl.io/ch7/eval"
	//"../" //GOPATH下にないときのimportパス

	"log"
)

func calc(w http.ResponseWriter, req *http.Request) {
	input := req.URL.Query().Get("formula")
	fmt.Println(input)
	if len(input) < 1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "formula is not found.")
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
		for v := range vars {
			s := req.URL.Query().Get(string(v))
			env[v], err = strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(fmt.Errorf("this value is invalid.\n%s", err))
			}
		}
	}
	fmt.Fprintf(w, "Answer: %f", e.Eval(env))
}

func main() {
	http.HandleFunc("/", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
