package generate

import (
	"fmt"
	"log"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

func Setup() {
	path := fmt.Sprintf("%s/%s", *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath)
	log.Println("cleaning up path: ", path)
	if err := os.RemoveAll(path); err != nil {
		log.Panic(err)
	}

	storagePath := fmt.Sprintf("%s/%s/%s", *config.GetFlags()[config.ServiceBasePath], paths.GeneratedFilePath, paths.ProtobufFilePath)
	log.Println("creating path: ", storagePath)
	if err := os.MkdirAll(storagePath, 0755); err != nil {
		log.Panic(err)
	}
}
