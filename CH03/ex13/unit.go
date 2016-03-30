// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved
package main

import "fmt"

//Unit of data
const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("1KB = %v\n", KB)
	fmt.Printf("1MB = %v\n", MB)
	fmt.Printf("1GB = %v\n", GB)
	fmt.Printf("1TB = %v\n", TB)
	fmt.Printf("1PB = %v\n", PB)
	fmt.Printf("1EB = %v\n", EB)
	//fmt.Printf("1ZB = %v\n", ZB)
	//fmt.Printf("1YB = %v\n", YB)
}
