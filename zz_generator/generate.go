package main

import (
	"io/ioutil"
	"log"

	"github.com/parijatpurohit/sidecar-sql/zz_generator/config"
	"github.com/parijatpurohit/sidecar-sql/zz_generator/genproto"
	"github.com/parijatpurohit/sidecar-sql/zz_generator/genserver"
	"github.com/parijatpurohit/sidecar-sql/zz_generator/genstorage"
)

func Generate() {
	files, err := ioutil.ReadDir(config.StorageConfigPath)
	if err != nil {
		log.Panic(err)
	}
	for _, f := range files {
		if f.Name() != config.CommonFileName {
			conf := config.GetStorageConfig(f.Name())
			generateForConfig(conf, f.Name())
		}
	}

}

func generateForConfig(conf *config.StorageConfig, fileName string) {
	genproto.GenerateEntityProto(conf, fileName)
	genstorage.Generate(conf, fileName)
	genserver.Generate(conf, fileName)
}
