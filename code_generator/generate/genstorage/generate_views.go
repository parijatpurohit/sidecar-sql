package genstorage

import (
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genviews"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

var viewGenerators = map[config.ViewType]func(*config.StorageConfig, *config.View){
	config.VIEW_TYPE_CREATE: genviews.GenerateNonReadView,
	config.VIEW_TYPE_READ:   genviews.GenerateReadView,
	config.VIEW_TYPE_UPDATE: genviews.GenerateNonReadView,
	config.VIEW_TYPE_DELETE: genviews.GenerateNonReadView,
}

func GenerateViews(storageConfig *config.StorageConfig) {

	for _, view := range storageConfig.Views {
		generator := viewGenerators[view.ViewType]
		if generator == nil {
			log.Printf("no view generator found for view: %s, type: %s", view.Name, view.ViewType)
			continue
		}
		log.Printf("generating view for: %s type: %s", view.Name, view.ViewType)
		generator(storageConfig, view)
	}
}

func GenerateViewDef(storageConfig *config.StorageConfig) {
	genviews.GenerateDefs(storageConfig)
}
