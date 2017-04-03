package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmaesaka/stapler/config"
)

func usage() {
	fmt.Printf("staplerd | the stapler daemon\n")
	fmt.Printf("  --port <num>  %s\n", flag.Lookup("port").Usage)
	os.Exit(0)
}

func main() {
	config := config.NewConfig()

	flag.IntVar(&config.Port, "port", 8084, "TCP port number to listen on (default: 8084)")

	fmt.Printf("listening on port: %d\n", config.Port)
}
