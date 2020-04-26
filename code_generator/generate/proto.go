package generate

import (
	"fmt"
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
			fmt.Println("generating proto for: ", f.Name())
			generateProtoForConfig(conf)
		}
	}
	genproto.GenerateServerProto(config.ParsedStorageConfig)
}

func generateProtoForConfig(conf *config.StorageConfig) {
	genproto.GenerateEntityProto(conf)
	genproto.GenerateViewProto(conf)
}
