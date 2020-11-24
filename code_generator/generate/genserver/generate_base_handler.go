package genserver

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

const (
	handlersPackageName = "handlers"
	protoImportKey      = "pb"
)

func GenerateBaseHandler(allConfig map[string]*config.StorageConfig) {
	conf := getBaseHandlerConfig(allConfig)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.ServerTemplatePath, paths.BaseHandlerTemplateFile))
	outFile, err := getOutputBaseHandlerFile()
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getBaseHandlerConfig(allConfig map[string]*config.StorageConfig) *HandlersConfig {
	handlersConfig := &HandlersConfig{}
	handlersConfig.PackageName = handlersPackageName
	for _, conf := range allConfig {
		tableName := generateUtils.GetTableName(conf.Table, conf.Common.IsPlural)
		handlersConfig.Tables = append(handlersConfig.Tables, &TableConfig{TableName: tableName})
	}
	handlersConfig.Imports = getBaseHandlerImports(allConfig)
	return handlersConfig
}

func getBaseHandlerImports(allConfig map[string]*config.StorageConfig) []*ImportConfig {
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

func getOutputBaseHandlerFile() (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.BaseHanderFilePath, paths.GeneratedFilePath)
	return os.Create(outputFilePath)
}
