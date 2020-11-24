package genstorage

import (
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genmodels"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

func GenerateStorage(conf *config.StorageConfig) {
	genmodels.GenerateSchema(conf)
	genmodels.GenerateQueryModels(conf)
}
