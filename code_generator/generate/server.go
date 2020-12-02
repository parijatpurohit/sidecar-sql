package generate

import (
	"fmt"
	"log"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genserver"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

func Server() {
	createServerPath()
	genserver.Generate(config.GetAllStorage())
}

func createServerPath() {
	path := fmt.Sprintf("%s/%s/%s/%s", *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath, paths.ServerOutputPath, paths.HandlersOutputPath)
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Panic(err)
	}
}
