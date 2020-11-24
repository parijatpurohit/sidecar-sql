package genserver

import (
	"fmt"
	"log"

	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

func Generate(config map[string]*config.StorageConfig) {
	log.Printf("generating base handlers for all views")
	GenerateBaseHandler(config)
	for _, conf := range config {
		log.Printf(fmt.Sprintf("generating table handlers for table: %s", conf.Table))
		GenerateTableHandler(conf)
	}
	GenerateServer(config)
}
