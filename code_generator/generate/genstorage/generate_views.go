package genstorage

import "github.com/parijatpurohit/sidecar-sql/code_generator/config"

type ViewConfig struct {
	View              *config.View
	PrimaryKeys       []string
	TableName         string
	HasTimestampField bool
}

func GenerateViews(sqlConfig *config.SQLConfig, storageConfig *config.StorageConfig) {
	for _, view := range storageConfig.Views {
		if view.ViewType == config.VIEW_TYPE_CREATE {
			GenerateCreateView(view)
		}
	}
}

func GenerateCreateView(viewConfig *config.View) {

}
