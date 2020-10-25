package generate

import (
	"fmt"
	"log"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genserver"
)

func Server() {
	createServerPath()
	genserver.Generate(config.GetAllStorage())
}

func createServerPath() {
	path := fmt.Sprintf("%s/%s/%s", paths.GeneratedFilePath, paths.ServerOutputPath, paths.HandlersOutputPath)
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Panic(err)
	}
}
