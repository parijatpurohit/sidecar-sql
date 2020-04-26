package config

import (
	"fmt"
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"

	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"

	"gopkg.in/yaml.v2"
)

var sqlConfig *SQLConfig

func GetSQLConfig() *SQLConfig {
	if sqlConfig != nil {
		sqlConfig = getSQLConfig()
	}
	return sqlConfig
}

func getSQLConfig() *SQLConfig {
	conf := &SQLConfig{}
	filePathStr := "%s/%s"
	bytes, err := utils.GetFileBytes(fmt.Sprintf(filePathStr, paths.StorageConfigPath, paths.SQLConfigFileName))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		log.Panicf("could not unmarshal file, err: %v", err)
	}
	return conf
}
