package gentranslator

import "github.com/parijatpurohit/sidecar-sql/lib/config"

type BaseTranslatorConfig struct {
	PackageName   string
	ViewName      string
	ViewNameLower string
	Imports       []*ImportConfig
	PrimaryKeys   []*ViewFieldConfig
	ViewFields    []*ViewFieldConfig
}

type ViewTranslatorConfig struct {
	PackageName    string
	TableName      string
	TableNameLower string
	ViewName       string
	ReturnType     config.ReturnType
	IsRead         bool
	MultiInput     bool
	MultiReturn    bool
	Imports        []*ImportConfig
	QueryFields    []*ViewFieldConfig
}

type ViewFieldConfig struct {
	ProtoFieldName string
	ViewFieldName  string
}
type ImportConfig struct {
	ImportKey  string
	ImportPath string
}
