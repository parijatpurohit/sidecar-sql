package genstorage

import (
	"log"

	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genviews"
)

var viewGenerators = map[config.ViewType]func(*config.StorageConfig, *config.View){
	config.VIEW_TYPE_CREATE: genviews.GenerateCreateView,
	//config.VIEW_TYPE_READ:   genviews.GenerateCreateView,
	//config.VIEW_TYPE_UPDATE: genviews.GenerateCreateView,
	//config.VIEW_TYPE_DELETE: genviews.GenerateCreateView,
}

func GenerateViews(storageConfig *config.StorageConfig) {
	for _, view := range storageConfig.Views {
		generator := viewGenerators[view.ViewType]
		if generator == nil {
			log.Printf("no view generator found for view: %s, type: %s", view.Name, view.ViewType)
			continue
		}
		generator(storageConfig, view)
	}
}
