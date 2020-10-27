package genserver

type HandlersConfig struct {
	PackageName string
	Imports     []*ImportConfig
	Tables      []*TableConfig
}

type ImportConfig struct {
	ImportKey  string
	ImportPath string
}
type TableConfig struct {
	TableName string
}

type ViewConfig struct {
	ViewName string
}

type TableHandlerConfig struct {
	PackageName    string
	Imports        []*ImportConfig
	TableName      string
	TableNameLower string
	Views          []*ViewConfig
}
