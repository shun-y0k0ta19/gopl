// Copyright © 2016 "Shun Yokota" All rights reserved

package main

import "os"
import "os/exec"

func main() {
	pack := os.Args[1]
	exec.Command("go list", pack).Run()
}
