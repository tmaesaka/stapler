package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmaesaka/stapler/app"
	"github.com/tmaesaka/stapler/config"
)

func usage() {
	fmt.Printf("stapler server\n")
	fmt.Printf("  --map  <path>  %s\n", flag.Lookup("map").Usage)
	fmt.Printf("  --port <num>   %s\n", flag.Lookup("port").Usage)

	os.Exit(0)
}

func init() {
	flag.Usage = usage
}

func main() {
	config := config.NewConfig()

	flag.StringVar(&config.MapPath, "map", "", "Path to the map file (required)")
	flag.IntVar(&config.Port, "port", 8084, "TCP port number to listen on (default: 8084)")
	flag.Parse()

	if err := config.Validate(); err != nil {
		fmt.Printf("%s\n\n", err.Error())
		usage()
	}

	if err := app.Run(config); err != nil {
		fmt.Printf("Failed to start stapler")
		os.Exit(1)
	}
}
