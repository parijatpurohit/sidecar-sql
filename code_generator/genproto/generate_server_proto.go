package genproto

import (
	"fmt"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
)

const (
	serverProtoPath    = "%s/pb/storage/server.proto"
	serverTemplateFile = "server.tgo"
	viewNameFormat     = "%s_%s"
)

type ServerProtoConfig struct {
	Views []string
}

func GenerateServerProto(storageConfigs map[string]*config.StorageConfig) {
	serverProtoConfig := ServerProtoConfig{}
	for _, storageConfig := range storageConfigs {
		for _, view := range storageConfig.Views {
			tableName := GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
			viewName := fmt.Sprintf(viewNameFormat, tableName, view.Name)
			serverProtoConfig.Views = append(serverProtoConfig.Views, viewName)
		}
	}

	template := getTemplate(serverTemplateFile)
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
	outputFilePath := fmt.Sprintf(serverProtoPath, config.GeneratedFilePath)
	return os.Create(outputFilePath)
}
