package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"
	"gopkg.in/yaml.v2"
)

const (
	ServiceBasePath = "basepath"
	GenTypeFlag     = "gentype"
	GoPathFlag      = "gopath"
)

var (
	parsedStorageConfig map[string]*StorageConfig
	once                sync.Once
	flagOnce            sync.Once
	flagKeys            = []string{ServiceBasePath, GenTypeFlag, GoPathFlag}
	flagMap             = map[string]*string{}
)

func InitStorageConfig() {
	once.Do(func() {
		parsedStorageConfig = map[string]*StorageConfig{}
		files, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", *GetFlags()[ServiceBasePath], paths.StorageConfigPath))
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
	filePathStr := "%s/%s/%s"
	conf = getYamlConfig(fmt.Sprintf(filePathStr, *GetFlags()[ServiceBasePath], paths.StorageConfigPath, paths.CommonConfigFileName), conf)
	return getYamlConfig(fmt.Sprintf(filePathStr, *GetFlags()[ServiceBasePath], paths.StorageConfigPath, entity), conf)
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

func GetFlags() map[string]*string {
	flagOnce.Do(func() {
		for _, f := range flagKeys {
			flagMap[f] = flag.String(f, "", "")
		}
		flag.Parse()
	})
	return flagMap
}

func GetBaseImportPath() string {
	gopath := fmt.Sprintf("%s/src/", *GetFlags()[GoPathFlag])
	importPath := *GetFlags()[ServiceBasePath]
	return strings.Replace(importPath, gopath, "", 1)

}
