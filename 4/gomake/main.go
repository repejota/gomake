package main

import (
	"fmt"
)

// Variables to identiy the build
var (
	Version string
	Build   string
)

func main() {
	fmt.Println("Version: ", Version)
	fmt.Println("Git commit hash: ", Build)
}
