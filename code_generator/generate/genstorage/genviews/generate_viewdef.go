package genviews

import (
	"fmt"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"

	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

const (
	InterfaceName   = "I%sViews"
	ImportSync      = "sync"
	ImportGorm      = "github.com/jinzhu/gorm"
	ImportSQLConn   = paths.BaseImportPath + "lib/sqlconn"
	CreatePrefix    = "Create"
	UpdatePrefix    = "Update"
	DeletePrefix    = "Delete"
	FieldDelimiter  = ","
	QueryFieldName  = "query"
	EntityFieldName = "entities"
	ResFieldName    = "res"

	QueryFieldType = "%s_%sQuery"
)

type FuncFieldConfig struct {
	FieldName string
	FieldType string
}

type ViewFuncConfig struct {
	Name        string
	InputVars   string
	ReturnVars  string
	ResponseVar string
}

type ViewDefConfig struct {
	PackageName   string
	Imports       []string
	InterfaceName string
	Views         []*ViewFuncConfig
}

func GenerateDefs(storageConfig *config.StorageConfig) {
	conf := getDefConfig(storageConfig)
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.StorageTemplatePath, paths.ViewDefTemplateFile))
	outFile, err := getOutputViewDefFile(tableName)
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getDefConfig(storageConfig *config.StorageConfig) *ViewDefConfig {
	tableName := generateUtils.GetTableName(storageConfig.Table, storageConfig.Common.IsPlural)
	return &ViewDefConfig{
		PackageName:   strings.ToLower(fmt.Sprintf(ViewPackageName, tableName)),
		Imports:       getDefImports(tableName),
		InterfaceName: fmt.Sprintf(InterfaceName, tableName),
		Views:         getViews(tableName, storageConfig),
	}
}

func getDefImports(tableName string) []string {
	imports := []string{
		ImportContext,
		ImportSync,
		ImportGorm,
		ImportSQLConn,
		fmt.Sprintf(paths.ModelsImportPath, strings.ToLower(tableName)),
	}
	return imports
}
func getViews(tableName string, storageConfig *config.StorageConfig) []*ViewFuncConfig {
	var configs []*ViewFuncConfig
	for _, view := range storageConfig.Views {
		configs = append(configs, getFuncDef(tableName, view))
	}
	return configs
}

func getOutputViewDefFile(tableName string) (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ViewsFilePath, paths.GeneratedFilePath, strings.ToLower(tableName), paths.ViewDefsFileName)
	return os.Create(outputFilePath)
}
