package main

import (
	"fmt"
)

// Variables to identify the build
var (
	Version = "1.0.0"
	Build   = "2015-11-25T00:23:32+0100"
)

func main() {
	fmt.Println("Version: ", Version)
	fmt.Println("Build Time: ", Build)
}
