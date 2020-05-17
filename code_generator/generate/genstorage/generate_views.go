package genstorage

import (
	"github.com/parijatpurohit/sidecar-sql/code_generator/config"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/genstorage/genviews"
)

type ViewConfig struct {
	View              *config.View
	PrimaryKeys       []string
	TableName         string
	HasTimestampField bool
}

func GenerateViews(storageConfig *config.StorageConfig) {
	for _, view := range storageConfig.Views {
		switch view.ViewType {
		case config.VIEW_TYPE_CREATE:
			genviews.GenerateCreateView(storageConfig, view)
		}
		if view.ViewType == config.VIEW_TYPE_CREATE {

		}
	}
}
