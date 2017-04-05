package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmaesaka/stapler/config"
)

func usage() {
	fmt.Printf("stapler server\n")
	fmt.Printf("  --map  <path>  %s\n", flag.Lookup("map").Usage)
	fmt.Printf("  --port <num>   %s\n", flag.Lookup("port").Usage)

	os.Exit(0)
}

func main() {
	config := config.NewConfig()

	flag.StringVar(&config.MapPath, "map", "", "Path to the map file (required)")
	flag.IntVar(&config.Port, "port", 8084, "TCP port number to listen on (default: 8084)")
	flag.Parse()

	if len(config.MapPath) == 0 {
		usage()
	}

	fmt.Printf("listening on port: %d\n", config.Port)
}
