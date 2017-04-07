package app

import (
	"fmt"

	"github.com/tmaesaka/stapler/config"
)

func Run(cfg *config.Config) error {
	fmt.Printf("listening on port: %d\n", cfg.Port)
	return nil
}
