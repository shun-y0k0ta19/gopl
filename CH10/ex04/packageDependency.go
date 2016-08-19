// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

import "os/exec"

//JSONFormat is struct to decode json.
type JSONFormat struct {
	//Dir        string
	ImportPath string
	Deps       []string
}

func main() {
	flag.Parse()
	for _, pack := range flag.Args() {
		pp, err := getPackagePath(pack)
		if err != nil {
			log.Fatal(err)
		}
		for _, ip := range pp {
			fmt.Printf("PackageName: %s\n", ip)
			dps := findDependedPackage(ip)
			fmt.Print("Depended Packages: \n")
			for dp := range dps {
				fmt.Printf("%s\n", dp)
			}
			fmt.Println()
		}
	}
}

func findDependedPackage(ip string) map[string]bool {
	args := []string{"list", "-e", "-json", "..."}
	res, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(bytes.NewReader(res))
	var jf JSONFormat
	deps := make(map[string]bool)
	for decoder.More() {
		err = decoder.Decode(&jf)
		if err != nil {
			log.Fatal(err)
		}
		for _, dep := range jf.Deps {
			if ip == dep {
				deps[jf.ImportPath] = true
			}
		}
	}
	return deps
}

func getPackagePath(pack string) ([]string, error) {
	var packagePath []string
	args := []string{"list", "-e", "-json", pack}
	res, err := exec.Command("go", args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(res))
	decoder := json.NewDecoder(bytes.NewReader(res))
	var jf JSONFormat
	for decoder.More() {
		err = decoder.Decode(&jf)
		if err != nil {
			return nil, err
		}
		packagePath = append(packagePath, jf.ImportPath)
	}
	return packagePath, nil
}
