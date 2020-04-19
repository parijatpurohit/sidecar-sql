package genproto

import (
	"fmt"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

const viewNameFormat = "%s_%s"

type ServerProtoConfig struct {
	Views []string
}

func GenerateServerProto(storageConfigs map[string]*config.StorageConfig) {
	serverProtoConfig := ServerProtoConfig{}
	for _, storageConfig := range storageConfigs {
		for _, view := range storageConfig.Views {
			tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
			viewName := fmt.Sprintf(viewNameFormat, tableName, view.Name)
			serverProtoConfig.Views = append(serverProtoConfig.Views, viewName)
		}
	}

	template := getTemplate(paths.ServerTemplateFile)
	outputFile, err := getOutputServerProtoFile()
	if err != nil {
		panic(err)
	}
	err = template.Execute(outputFile, serverProtoConfig)
	if err != nil {
		panic(err)
	}
}

func getOutputServerProtoFile() (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ServerProtoPath, paths.GeneratedFilePath)
	return os.Create(outputFilePath)
}
