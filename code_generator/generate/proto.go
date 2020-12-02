package generate

import (
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genproto"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

func Proto() {
	for name, conf := range config.GetAllStorage() {
		log.Printf("generating proto for: %s", name)
		generateProtoForConfig(conf)
	}
	genproto.GenerateServerProto(config.GetAllStorage())
}

func Protogen() {
	genproto.GenerateProtoOutput()
}

func generateProtoForConfig(conf *config.StorageConfig) {
	log.Printf("generating proto entity for: %s", conf.Table)
	genproto.GenerateEntityProto(conf)
	log.Printf("generating proto views for: %s", conf.Table)
	genproto.GenerateViewProto(conf)
}
