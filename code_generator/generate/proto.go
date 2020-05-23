package generate

import (
	"io/ioutil"
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genproto"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/paths"
)

func Proto() {
	files, err := ioutil.ReadDir(paths.StorageConfigPath)
	if err != nil {
		log.Panic(err)
	}
	for _, f := range files {
		if f.Name() != paths.CommonConfigFileName {
			conf := config.GetStorageConfig(f.Name())
			log.Printf("generating proto for: %s", f.Name())
			generateProtoForConfig(conf)
		}
	}
	genproto.GenerateServerProto(config.ParsedStorageConfig)
}

func generateProtoForConfig(conf *config.StorageConfig) {
	log.Printf("generating proto entity for: %s", conf.Table)
	genproto.GenerateEntityProto(conf)
	log.Printf("generating proto views for: %s", conf.Table)
	genproto.GenerateViewProto(conf)
}
