package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/parijatpurohit/sidecar-sql/zz_generator/configreader"
)

func Generate() {
	files, err := ioutil.ReadDir(configreader.StorageConfigPath)
	if err != nil {
		log.Panic(err)
	}

	for _, f := range files {
		if f.Name() != configreader.CommonFileName {
			config, _ := json.Marshal(configreader.GetStorageConfig(f.Name()))
			fmt.Println(string(config))
		}
	}
}
