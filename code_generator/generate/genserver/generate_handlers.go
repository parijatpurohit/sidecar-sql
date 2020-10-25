package genserver

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

const (
	handlersPackageName = "handlers"
	protoImportKey      = "pb"
)

func GenerateHandlers(allConfig map[string]*config.StorageConfig) {
	conf := getHandlerConfig(allConfig)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.ServerTemplatePath, paths.HandlersTemplateFile))
	outFile, err := getOutputHandlerSchema()
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getHandlerConfig(allConfig map[string]*config.StorageConfig) *HandlersConfig {
	handlersConfig := &HandlersConfig{}
	handlersConfig.PackageName = handlersPackageName
	for _, conf := range allConfig {
		tableName := generateUtils.GetTableName(conf.Table, conf.Common.IsPlural)
		handlersConfig.Views = append(handlersConfig.Views, &ViewConfig{ViewName: tableName})
	}
	handlersConfig.Imports = getHandlerImports(allConfig)
	return handlersConfig
}

func getHandlerImports(allConfig map[string]*config.StorageConfig) []*ImportConfig {
	var imports []*ImportConfig
	protoImport := &ImportConfig{ImportKey: protoImportKey, ImportPath: paths.ProtoImportPath}
	imports = append(imports, protoImport)
	for _, conf := range allConfig {
		tableName := generateUtils.GetTableName(conf.Table, conf.Common.IsPlural)
		viewFilePath := fmt.Sprintf(paths.ViewsImportPath, strings.ToLower(tableName))
		importKey := fmt.Sprintf("%s_Views", tableName)
		imports = append(imports, &ImportConfig{ImportKey: importKey, ImportPath: viewFilePath})
	}

	return imports
}

func getOutputHandlerSchema() (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.HandersFilePath, paths.GeneratedFilePath)
	return os.Create(outputFilePath)
}

func GenerateViewHandler(config *config.StorageConfig) {

}
