package generate

import (
	"io/ioutil"
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
)

func Storage() {
	files, err := ioutil.ReadDir(paths.StorageConfigPath)
	if err != nil {
		log.Panic(err)
	}
	for _, f := range files {
		if f.Name() != paths.CommonFileName {
			conf := config.GetStorageConfig(f.Name())
			genstorage.GenerateStorage(conf)
		}
	}
}
