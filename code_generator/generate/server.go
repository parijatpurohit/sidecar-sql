package generate

import (
	"io/ioutil"
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genserver"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
)

func Server() {
	files, err := ioutil.ReadDir(paths.StorageConfigPath)
	if err != nil {
		log.Panic(err)
	}
	for _, f := range files {
		if f.Name() != paths.CommonConfigFileName {
			conf := config.GetStorageConfig(f.Name())
			genserver.Generate(conf)
		}
	}
}
