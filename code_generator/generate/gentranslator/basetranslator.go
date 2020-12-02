package gentranslator

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

const (
	modelsImportKey = "%s_models"
	protoImportKey  = "pb"
)

var defaultFields = []string{config.FIELD_CREATED_AT, config.FIELD_UPDATED_AT, config.FIELD_DELETED_AT}

func GenerateBaseTranslator(config *config.StorageConfig) {
	conf := getBaseTranslatorConfig(config)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.TranslatorTemplatePath, paths.BaseTranslatorTemplateFile))
	tableName := generateUtils.GetTableName(config.Table, config.Common.IsPlural)
	outFile, err := getOutputBaseTranslatorFile(tableName)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getBaseTranslatorConfig(config *config.StorageConfig) *BaseTranslatorConfig {
	tableName := generateUtils.GetTableName(config.Table, config.Common.IsPlural)
	return &BaseTranslatorConfig{
		PackageName:   strings.ToLower(tableName),
		ViewName:      tableName,
		ViewNameLower: strings.ToLower(tableName),
		ViewFields:    getBaseTranslatorFields(config),
		Imports:       getBaseTranslatorImports(config),
		PrimaryKeys:   getPrimaryKeys(config),
	}
}

func getPrimaryKeys(storageConfig *config.StorageConfig) []*ViewFieldConfig {
	_, primaryKeys := generateUtils.GetFieldConfig(storageConfig)
	var res []*ViewFieldConfig
	for _, field := range primaryKeys {
		res = append(res, &ViewFieldConfig{ProtoFieldName: field.FieldName, ViewFieldName: field.FieldName})
	}
	return res
}

func getBaseTranslatorFields(config *config.StorageConfig) []*ViewFieldConfig {
	var viewFieldConfig []*ViewFieldConfig
	for _, field := range config.Fields {
		if utils.ContainsStr(defaultFields, field.FieldName) != -1 {
			continue
		}
		viewFieldConfig = append(viewFieldConfig, &ViewFieldConfig{ProtoFieldName: field.FieldName, ViewFieldName: field.FieldName})
	}
	return viewFieldConfig
}

func getBaseTranslatorImports(conf *config.StorageConfig) []*ImportConfig {
	tableNameLower := strings.ToLower(generateUtils.GetTableName(conf.Table, conf.Common.IsPlural))
	imports := []*ImportConfig{
		{ImportPath: paths.PTypesImportPath},
		{ImportPath: paths.TimestampImportPath},
		{ImportPath: paths.BaseImportPath + paths.DataUtilsRelativePath},
		{ImportKey: protoImportKey, ImportPath: fmt.Sprintf(paths.ProtoImportPath, config.GetBaseImportPath())},
		{ImportKey: fmt.Sprintf(modelsImportKey, tableNameLower), ImportPath: fmt.Sprintf(paths.ModelsImportPath, config.GetBaseImportPath(), tableNameLower)},
	}
	return imports
}

func getOutputBaseTranslatorFile(tableName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.BaseTranslatorFilePath, *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath, strings.ToLower(tableName))
	return os.Create(outputFilePath)
}
