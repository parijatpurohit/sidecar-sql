package genserver

type HandlersConfig struct {
	PackageName string
	Imports     []*ImportConfig
	Views       []*ViewConfig
}

type ImportConfig struct {
	ImportKey  string
	ImportPath string
}
type ViewConfig struct {
	ViewName string
}
