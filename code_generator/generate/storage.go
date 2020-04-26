package generate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
)

func Storage() {
	files, err := ioutil.ReadDir(paths.StorageConfigPath)
	if err != nil {
		log.Panic(err)
	}
	err = os.MkdirAll(fmt.Sprintf("%s/%s", paths.GeneratedFilePath, paths.StorageOutputPath), 0755)
	if err != nil {
		log.Panic(err)
	}
	for _, f := range files {
		if f.Name() != paths.CommonConfigFileName {
			conf := config.GetStorageConfig(f.Name())
			genstorage.GenerateStorage(conf)
		}
	}
}
