package configreader

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var storageConfig map[string]*StorageConfig

const (
	StorageConfigPath = "configuration/storage"
	CommonFileName    = "common.yaml"
)

func GetStorageConfig(entity string) *StorageConfig {
	if storageConfig == nil {
		storageConfig = map[string]*StorageConfig{}
	}
	if storageConfig[entity] == nil {
		storageConfig[entity] = getFinalConfig(entity)
	}
	return storageConfig[entity]
}

func getFinalConfig(entity string) *StorageConfig {
	config := &StorageConfig{}
	filePathStr := "%s/%s"
	config = getYamlConfig(fmt.Sprintf(filePathStr, StorageConfigPath, CommonFileName), config)
	return getYamlConfig(fmt.Sprintf(filePathStr, StorageConfigPath, entity), config)
}

func getYamlConfig(path string, config *StorageConfig) *StorageConfig {
	bytes, err := getFileBytes(path)
	if err != nil {
		panic(err)
	}
	return getYamlConfigForData(bytes, config)
}

func getYamlConfigForData(data []byte, config *StorageConfig) *StorageConfig {
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("could not unmarshal file, err: %v", err)
	}
	return config
}

func getFileBytes(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("could not read file, err: %v", err)
		return nil, err
	}
	return content, nil
}
