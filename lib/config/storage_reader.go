package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"
	"gopkg.in/yaml.v2"
)

var parsedStorageConfig map[string]*StorageConfig
var once sync.Once

func InitStorageConfig() {
	once.Do(func() {
		parsedStorageConfig = map[string]*StorageConfig{}
		files, err := ioutil.ReadDir(paths.StorageConfigPath)
		if err != nil {
			log.Panic(err)
		}
		for _, f := range files {
			if f.Name() != paths.CommonConfigFileName {
				if parsedStorageConfig[f.Name()] == nil {
					parsedStorageConfig[f.Name()] = getFinalConfig(f.Name())
				}
			}
		}
	})
}

func GetAllStorage() map[string]*StorageConfig {
	return parsedStorageConfig
}
func GetStorageConfig(entity string) *StorageConfig {
	return parsedStorageConfig[entity]
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