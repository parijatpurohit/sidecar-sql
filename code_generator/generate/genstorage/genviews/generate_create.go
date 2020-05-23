package genviews

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genmodels"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
)

const (
	InputPackageName = "%s_models"
	ViewPackageName  = "%s_views"
	ImportSQL        = "database/sql"
	ImportContext    = "context"
)

type CreateViewConfig struct {
	PackageName string
	Imports     []string
	ViewName    string
	InputType   string
}

func GenerateCreateView(storageConfig *config.StorageConfig, viewConfig *config.View) {
	conf := getCreateConfig(storageConfig, viewConfig)
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, paths.CreateViewTemplateFile))
	outFile, err := getOutputCreateViewFile(tableName)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getCreateConfig(storageConfig *config.StorageConfig, viewConfig *config.View) *CreateViewConfig {
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	return &CreateViewConfig{
		PackageName: fmt.Sprintf(ViewPackageName, tableName),
		Imports:     getCreateImports(tableName),
		ViewName:    tableName,
		InputType:   getInputType(tableName),
	}
}

func getInputType(tableName string) string {
	return fmt.Sprintf("[]*%s.%s", fmt.Sprintf(genmodels.SchemaPackageFormat, strings.ToLower(tableName)), tableName)
}

func getCreateImports(tableName string) []string {
	var imports = []string{
		ImportContext,
		ImportSQL,
		fmt.Sprintf(paths.ModelsImportPath, tableName),
	}
	return imports
}

func getOutputCreateViewFile(tableName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ViewsFilePath, paths.GeneratedFilePath, strings.ToLower(tableName))
	return os.Create(outputFilePath)
}
