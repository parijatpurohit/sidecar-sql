package genviews

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/alias"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

const (
	ViewPackageName = "%s_views"
	ImportSQL       = "database/sql"
	ImportContext   = "context"
	ImportTime      = "time"
	ImportError     = "errors"
)

type NonReadViewConfig struct {
	PackageName string
	Imports     []string
	FuncDef     *ViewFuncConfig
}

func GenerateNonReadView(storageConfig *config.StorageConfig, viewConfig *config.View) {
	conf := getNonReadConfig(storageConfig, viewConfig)
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, alias.GetTemplateForView[viewConfig.ViewType]))
	outFile, err := getOutputNonReadViewFile(tableName, viewConfig)
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
		PackageName: strings.ToLower(fmt.Sprintf(ViewPackageName, tableName)),
		Imports:     getNonReadImports(tableName, viewConfig.ViewType),
		FuncDef:     getFuncDef(tableName, viewConfig),
	}
}

func getNonReadImports(tableName string, viewType config.ViewType) []string {
	var imports = []string{
		ImportContext,
		ImportSQL,
		fmt.Sprintf(paths.ModelsImportPath, strings.ToLower(tableName)),
	}
	if viewType == config.VIEW_TYPE_UPDATE {
		imports = append(imports, ImportError)
	}
	if viewType == config.VIEW_TYPE_DELETE || viewType == config.VIEW_TYPE_UPDATE {
		imports = append(imports, ImportTime, fmt.Sprintf("%sutils", paths.BaseImportPath))
	}
	return imports
}

func getOutputNonReadViewFile(tableName string, viewConfig *config.View) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ViewsFilePath, paths.GeneratedFilePath, strings.ToLower(tableName), strings.ToLower(viewConfig.Name))
	return os.Create(outputFilePath)
}
