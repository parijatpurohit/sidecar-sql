package generate

import (
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genserver"
)

func Server() {
	for name, conf := range config.GetAllStorage() {
		log.Printf("generating server for: %s", name)
		genserver.Generate(conf)
	}
}
