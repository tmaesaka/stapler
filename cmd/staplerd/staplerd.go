package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("staplerd | the stapler daemon\n")
	os.Exit(0)
}

func main() {
	usage()
}
