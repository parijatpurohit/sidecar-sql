package genserver

import (
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
)

func Generate(config map[string]*config.StorageConfig) {
	log.Printf("generating handlers for all views")
	GenerateHandlers(config)
}
