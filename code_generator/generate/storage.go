package generate

import (
	"fmt"
	"log"
	"os"
	"strings"

	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

func Storage() {
	for _, conf := range config.GetAllStorage() {
		tableName := generateUtils.GetTableName(conf.Table, conf.Common.IsPlural)
		createStoragePath(tableName)
		genstorage.GenerateStorage(conf)
		log.Printf("generating view interface for table: %s", tableName)
		genstorage.GenerateViewDef(conf)
		genstorage.GenerateViews(conf)
	}
}

func createStoragePath(tableName string) {
	var storagePaths = []string{
		fmt.Sprintf("%s/%s/%s/%s/%s", *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath, paths.StorageOutputPath, strings.ToLower(tableName), paths.ModelsOutputPath),
		fmt.Sprintf("%s/%s/%s/%s/%s", *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath, paths.StorageOutputPath, strings.ToLower(tableName), paths.ViewsOutputPath),
	}
	for _, path := range storagePaths {
		if err := os.MkdirAll(path, 0755); err != nil {
			log.Panic(err)
		}
	}
}
