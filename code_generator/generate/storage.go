package generate

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genmodels"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
)

func Storage() {
	for _, conf := range config.GetAllStorage() {
		createPath(conf.Table)
		genmodels.GenerateSchema(conf)
		genmodels.GenerateQueryModels(conf)
		genstorage.GenerateViews(conf)
	}
}

func createPath(tableName string) {
	dirPath := fmt.Sprintf("%s/%s/%s/%s", paths.GeneratedFilePath, paths.StorageOutputPath, strings.ToLower(tableName), paths.ModelsOutputPath)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		log.Panic(err)
	}
}
