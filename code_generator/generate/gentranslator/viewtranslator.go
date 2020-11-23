package gentranslator

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

func GenerateViewTranslator(config *config.StorageConfig) {
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.TranslatorTemplatePath, paths.ViewTranslatorTemplateFile))
	tableName := generateUtils.GetTableName(config.Table, config.Common.IsPlural)
	for _, view := range config.Views {
		conf := getViewTranslatorConfig(config, view)
		outFile, err := getOutputViewTranslatorFile(tableName, view.Name)
		if err != nil {
			panic(err)
		}
		err = tpl.Execute(outFile, conf)
		if err != nil {
			panic(err)
		}
	}

}

func getViewTranslatorConfig(conf *config.StorageConfig, viewConfig *config.View) *ViewTranslatorConfig {
	tableName := generateUtils.GetTableName(conf.Table, conf.Common.IsPlural)
	viewTranslatorConfig := &ViewTranslatorConfig{
		PackageName:    strings.ToLower(tableName),
		TableName:      tableName,
		TableNameLower: strings.ToLower(tableName),
		ViewName:       viewConfig.Name,
		IsRead:         viewConfig.ViewType == config.VIEW_TYPE_READ,
		ReturnType:     viewConfig.Config.ReturnType,
		MultiInput:     viewConfig.Config.MultiInput,
		MultiReturn:    viewConfig.Config.MultiReturn,
		QueryFields: 	getQueryFields(viewConfig),
	}
	viewTranslatorConfig.Imports = getViewTranslatorImports(conf)
	return viewTranslatorConfig
}

func getQueryFields(viewConfig *config.View) []*ViewFieldConfig {
	var res []*ViewFieldConfig
	if viewConfig.Query == nil {
		return res
	}
	for _,field := range viewConfig.Query.Fields {
		res = append(res, &ViewFieldConfig{
			ProtoFieldName: field.FieldName,
			ViewFieldName:  field.FieldName,
		})
	}
	return res
}
func getViewTranslatorImports(config *config.StorageConfig) []*ImportConfig {
	tableNameLower := strings.ToLower(generateUtils.GetTableName(config.Table, config.Common.IsPlural))
	imports := []*ImportConfig{
		{ImportKey: protoImportKey, ImportPath: paths.ProtoImportPath},
		{ImportKey: fmt.Sprintf(modelsImportKey, tableNameLower), ImportPath: fmt.Sprintf(paths.ModelsImportPath, tableNameLower)},
	}
	return imports
}

func getOutputViewTranslatorFile(tableName string, viewName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ViewTranslatorFilePath, paths.GeneratedFilePath, strings.ToLower(tableName), strings.ToLower(viewName))
	return os.Create(outputFilePath)
}
