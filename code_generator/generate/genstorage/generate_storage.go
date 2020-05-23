package genstorage

import (
	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genmodels"
)

func GenerateStorage(conf *config.StorageConfig) {
	genmodels.GenerateSchema(conf)
	genmodels.GenerateQueryModels(conf)
}
