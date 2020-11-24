package genviews

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/alias"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"

	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

const (
	AND_SEPARATOR = " AND "
)

type ReadViewConfig struct {
	PackageName    string
	Imports        []string
	ViewName       string
	TableName      string
	TableNameLower string
	Query          string
	QueryFields    []string
}

func GenerateReadView(storageConfig *config.StorageConfig, viewConfig *config.View) {
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	conf := getReadConfig(storageConfig, viewConfig, tableName)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, alias.GetTemplateForView[viewConfig.ViewType]))
	outFile, err := getOutputReadViewFile(tableName, viewConfig)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getReadConfig(storageConfig *config.StorageConfig, viewConfig *config.View, tableName string) *ReadViewConfig {
	return &ReadViewConfig{
		PackageName:    strings.ToLower(fmt.Sprintf(ViewPackageName, tableName)),
		Imports:        getReadImports(tableName, viewConfig.ViewType),
		ViewName:       viewConfig.Name,
		TableName:      tableName,
		TableNameLower: strings.ToLower(tableName),
		Query:          getQuery(viewConfig),
		QueryFields:    getQueryFields(viewConfig),
	}
}

func queryParamsToString(fieldName string, operator config.SearchType) string {
	return fmt.Sprintf("%s %s (?)", strings.ToLower(fieldName), alias.OperatorMap[operator])
}

func getQuery(viewConfig *config.View) string {
	var finalQuery []string
	for _, field := range viewConfig.Query.Fields {
		finalQuery = append(finalQuery, queryParamsToString(field.FieldName, field.SearchType))
	}
	return strings.Join(finalQuery, AND_SEPARATOR)
}

func getQueryFields(viewConfig *config.View) []string {
	var res []string
	for _, field := range viewConfig.Query.Fields {
		res = append(res, field.FieldName)
	}
	return res
}

func getReadImports(tableName string, viewType config.ViewType) []string {
	var imports = []string{
		ImportContext,
		fmt.Sprintf(paths.ModelsImportPath, strings.ToLower(tableName)),
	}
	return imports
}

func getOutputReadViewFile(tableName string, viewConfig *config.View) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ViewsFilePath, paths.GeneratedFilePath, strings.ToLower(tableName), strings.ToLower(viewConfig.Name))
	return os.Create(outputFilePath)
}
