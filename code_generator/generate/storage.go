package generate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genmodels"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
)

func Storage() {
	files, err := ioutil.ReadDir(paths.StorageConfigPath)
	if err != nil {
		log.Panic(err)
	}
	// sqlConf := config.GetSQLConfig()
	for _, f := range files {
		if f.Name() != paths.CommonConfigFileName {
			storageConf := config.GetStorageConfig(f.Name())
			dirPath := fmt.Sprintf("%s/%s/%s/%s", paths.GeneratedFilePath, paths.StorageOutputPath, strings.ToLower(storageConf.Table), paths.ModelsOutputPath)
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				log.Panic(err)
			}
			genmodels.GenerateSchema(storageConf)
			genmodels.GenerateQueryModels(storageConf)
			genstorage.GenerateViews(storageConf)
		}
	}
}
