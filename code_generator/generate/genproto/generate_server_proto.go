package genproto

import (
	"fmt"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

const viewNameFormat = "%s_%s"

type ServerProtoConfig struct {
	Views []string
}

func GenerateServerProto(storageConfigs map[string]*config.StorageConfig) {

	serverProtoConfig := getServerProtoConfig(storageConfigs)
	template := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.ProtoTemplatePath, paths.ProtoServerTemplateFile))
	outputFile, err := getOutputServerProtoFile()
	if err != nil {
		panic(err)
	}
	err = template.Execute(outputFile, serverProtoConfig)
	if err != nil {
		panic(err)
	}
}

func getServerProtoConfig(storageConfigs map[string]*config.StorageConfig) *ServerProtoConfig {
	serverProtoConfig := &ServerProtoConfig{}
	for _, storageConfig := range storageConfigs {
		for _, view := range storageConfig.Views {
			tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
			viewName := fmt.Sprintf(viewNameFormat, tableName, view.Name)
			serverProtoConfig.Views = append(serverProtoConfig.Views, viewName)
		}
	}
	return serverProtoConfig
}

func getOutputServerProtoFile() (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ServerProtoPath, paths.GeneratedFilePath)
	return os.Create(outputFilePath)
}
