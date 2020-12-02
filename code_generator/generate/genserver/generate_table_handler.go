package genserver

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genviews"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

const (
	contextImport = "context"
	logImport     = "log"
)

func GenerateTableHandler(config *config.StorageConfig) {
	conf := getTableHandlerConfig(config)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.ServerTemplatePath, paths.TableHandlerTemplateFile))
	outFile, err := getOutputTableHandlerFile(conf.TableNameLower)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getTableHandlerConfig(conf *config.StorageConfig) *TableHandlerConfig {
	tableName := generateUtils.GetTableName(conf.Table, conf.Common.IsPlural)
	return &TableHandlerConfig{
		PackageName:    handlersPackageName,
		Imports:        getTableHandlerImports(conf),
		TableName:      tableName,
		TableNameLower: strings.ToLower(tableName),
		Views:          getTableViews(conf),
	}
}

func getTableHandlerImports(conf *config.StorageConfig) []*ImportConfig {
	var imports []*ImportConfig

	protoImport := &ImportConfig{ImportKey: protoImportKey, ImportPath: fmt.Sprintf(paths.ProtoImportPath, config.GetBaseImportPath())}
	contextImport := &ImportConfig{ImportPath: contextImport}
	logImport := &ImportConfig{ImportPath: logImport}

	tableName := generateUtils.GetTableName(conf.Table, conf.Common.IsPlural)
	viewFilePath := fmt.Sprintf(paths.TranslatorImportPath, config.GetBaseImportPath(), paths.GeneratedFilePath, strings.ToLower(tableName))
	importKey := fmt.Sprintf("%s", strings.ToLower(tableName))

	imports = append(imports, &ImportConfig{ImportKey: importKey, ImportPath: viewFilePath})
	imports = append(imports, protoImport, contextImport, logImport)

	return imports
}

func getTableViews(conf *config.StorageConfig) []*ViewConfig {
	var res []*ViewConfig
	for _, view := range conf.Views {
		res = append(res, &ViewConfig{ViewName: genviews.GetViewName(view)})
	}
	return res
}

func getOutputTableHandlerFile(tableName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.TableHanderFilePath, *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath, tableName)
	return os.Create(outputFilePath)
}
