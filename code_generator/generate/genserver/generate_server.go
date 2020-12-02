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
	defaultGRPCPort   = int64(50010)
	netImportPath     = "net"
	grpcImportPath    = "google.golang.org/grpc"
	serverPackageName = "server"
)

func GenerateServer(config map[string]*config.StorageConfig) {
	conf := getServerConfig(config)
	tpl := generateUtils.GetTemplate(fmt.Sprintf("%s/%s", paths.ServerTemplatePath, paths.ServerTemplateFile))
	outFile, err := getOutputServerFile()
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(outFile, conf)
	if err != nil {
		panic(err)
	}
}

func getServerConfig(conf map[string]*config.StorageConfig) *ServerConfig {
	var tables []*TableConfig
	grpcPort := defaultGRPCPort
	for _, tableConfig := range conf {
		grpcPort = tableConfig.Common.GRPCPort
		tableName := generateUtils.GetTableName(tableConfig.Table, tableConfig.Common.IsPlural)
		tables = append(tables, &TableConfig{TableName: tableName, TableNameLower: strings.ToLower(tableName)})
	}
	return &ServerConfig{
		PackageName:     serverPackageName,
		Imports:         getServerImports(conf),
		Tables:          tables,
		GrpcPort:        grpcPort,
		ServiceBasePath: *config.GetFlags()[config.ServiceBasePath],
	}
}

func getServerImports(conf map[string]*config.StorageConfig) []*ImportConfig {
	imports := []*ImportConfig{
		{ImportKey: protoImportKey, ImportPath: fmt.Sprintf(paths.ProtoImportPath, config.GetBaseImportPath())},
		{ImportPath: netImportPath},
		{ImportPath: logImport},
		{ImportPath: grpcImportPath},
		{ImportPath: fmt.Sprintf(paths.HandlersRelativePath, config.GetBaseImportPath(), paths.GeneratedFilePath)},
		{ImportPath: paths.SqlConnImportPath},
		{ImportPath: paths.ConfigImportPath},
	}

	for _, tableConfig := range conf {
		tableName := generateUtils.GetTableName(tableConfig.Table, tableConfig.Common.IsPlural)
		viewFilePath := fmt.Sprintf(paths.ViewsImportPath, config.GetBaseImportPath(), strings.ToLower(tableName))
		importKey := fmt.Sprintf("%sviews", strings.ToLower(tableName))
		imports = append(imports, &ImportConfig{ImportKey: importKey, ImportPath: viewFilePath})
	}

	return imports
}

func getOutputServerFile() (*os.File, error) {
	outputFilePath := fmt.Sprintf(paths.ServerFilePath, *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath)
	return os.Create(outputFilePath)
}
