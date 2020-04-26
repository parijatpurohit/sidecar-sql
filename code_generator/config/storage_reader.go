package config

import (
	"fmt"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"
	"gopkg.in/yaml.v2"
)

var ParsedStorageConfig map[string]*StorageConfig

func GetStorageConfig(entity string) *StorageConfig {
	if ParsedStorageConfig == nil {
		ParsedStorageConfig = map[string]*StorageConfig{}
	}
	if ParsedStorageConfig[entity] == nil {
		ParsedStorageConfig[entity] = getFinalConfig(entity)
	}
	return ParsedStorageConfig[entity]
}

func getFinalConfig(entity string) *StorageConfig {
	conf := &StorageConfig{}
	filePathStr := "%s/%s"
	conf = getYamlConfig(fmt.Sprintf(filePathStr, paths.StorageConfigPath, paths.CommonConfigFileName), conf)
	return getYamlConfig(fmt.Sprintf(filePathStr, paths.StorageConfigPath, entity), conf)
}

func getYamlConfig(path string, conf *StorageConfig) *StorageConfig {
	bytes, err := utils.GetFileBytes(path)
	if err != nil {
		panic(err)
	}
	return getYamlConfigForData(bytes, conf)
}

func getYamlConfigForData(data []byte, config *StorageConfig) *StorageConfig {
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		panic(fmt.Sprintf("could not unmarshal file, err: %v", err))
	}
	return config
}
