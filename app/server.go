package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/tmaesaka/stapler/config"
)

// Temporarily unmarshal the map config into this map.
type StaplerMap struct {
	Routes []map[string]interface{}
}

func Run(cfg *config.Config) error {
	var (
		mapContent []byte
		err        error
		staplerMap StaplerMap
	)

	if mapContent, err = ioutil.ReadFile(cfg.MapPath); err != nil {
		log.Fatalf("Failed to open: %s", cfg.MapPath)
	}

	if err = json.Unmarshal(mapContent, &staplerMap); err != nil {
		log.Fatalf("Failed to load: %s %v", cfg.MapPath, err)
	}

	for _, route := range staplerMap.Routes {
		fmt.Printf("%s => %s\n", route["path"], route["file"])
	}

	fmt.Printf("listening on port: %d\n", cfg.Port)
	return nil
}
