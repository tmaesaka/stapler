package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("staplerd | the stapler daemon\n")
	fmt.Printf("  --port <num>  %s\n", flag.Lookup("port").Usage)
	os.Exit(0)
}

func main() {
	var port int

	flag.IntVar(&port, "port", 8084, "TCP port number to listen on (default: 8084)")

	fmt.Printf("listening on port: %d\n", port)
}
