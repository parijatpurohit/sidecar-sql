package genviews

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

const (
	ViewPackageName = "%s_views"
	ImportSQL       = "database/sql"
	ImportContext   = "context"
	ImportTime      = "time"
)

type NonReadViewConfig struct {
	PackageName string
	Imports     []string
	FuncDef     *ViewFuncConfig
}

var fileNameForView = map[config.ViewType]string{
	config.VIEW_TYPE_CREATE: paths.CreateFileName,
	config.VIEW_TYPE_UPDATE: paths.UpdateFileName,
	config.VIEW_TYPE_DELETE: paths.DeleteFileName,
}

var templateForView = map[config.ViewType]string{
	config.VIEW_TYPE_CREATE: paths.CreateViewTemplateFile,
	config.VIEW_TYPE_UPDATE: paths.UpdateViewTemplateFile,
	config.VIEW_TYPE_DELETE: paths.DeleteViewTemplateFile,
}

func GenerateNonReadView(storageConfig *config.StorageConfig, viewConfig *config.View) {
	conf := getNonReadConfig(storageConfig, viewConfig)
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, templateForView[viewConfig.ViewType]))
	outFile, err := getOutputNonReadViewFile(tableName, viewConfig.ViewType)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getNonReadConfig(storageConfig *config.StorageConfig, viewConfig *config.View) *NonReadViewConfig {
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	return &NonReadViewConfig{
		PackageName: fmt.Sprintf(ViewPackageName, tableName),
		Imports:     getNonReadImports(tableName, viewConfig.ViewType),
		FuncDef:     getFuncDef(tableName, viewConfig),
	}
}

func getNonReadImports(tableName string, viewType config.ViewType) []string {
	var imports = []string{
		ImportContext,
		ImportSQL,
		fmt.Sprintf(paths.ModelsImportPath, tableName),
	}
	if viewType == config.VIEW_TYPE_DELETE || viewType == config.VIEW_TYPE_UPDATE {
		imports = append(imports, ImportTime, fmt.Sprintf("%sutils", paths.BaseImportPath))
	}
	return imports
}

func getOutputNonReadViewFile(tableName string, viewType config.ViewType) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ViewsFilePath, paths.GeneratedFilePath, strings.ToLower(tableName), fileNameForView[viewType])
	return os.Create(outputFilePath)
}
