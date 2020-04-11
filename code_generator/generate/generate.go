package generate

import (
	"io/ioutil"
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/genproto"
	"github.com/parijatpurohit/sidecar-sql/code_generator/genserver"
	"github.com/parijatpurohit/sidecar-sql/code_generator/genstorage"
)

func Generate() {
	files, err := ioutil.ReadDir(config.StorageConfigPath)
	if err != nil {
		log.Panic(err)
	}
	for _, f := range files {
		if f.Name() != config.CommonFileName {
			conf := config.GetStorageConfig(f.Name())
			generateForConfig(conf)
		}
	}
	genproto.GenerateServerProto(config.ParsedStorageConfig)

}

func generateForConfig(conf *config.StorageConfig) {
	genproto.GenerateEntityProto(conf)
	genproto.GenerateViewProto(conf)
	genstorage.Generate(conf)
	genserver.Generate(conf)
}
