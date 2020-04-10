package configreader

type CommonConfig struct {
	IsPlural          bool  `yaml:"isPlural"`
	MaxIdle           int64 `yaml:"maxIdle"`
	MaxOpen           int64 `yaml:"maxOpen"`
	SoftDeleteEnabled bool  `yaml:"softDeleteEnabled"`
}

type StorageConfig struct {
	Table  string       `yaml:"table"`
	Fields []*Field     `yaml:"fields"`
	Views  []*View      `yaml:"views"`
	Common CommonConfig `yaml:"common"`
}

type Field struct {
	FieldName  string    `yaml:"fieldName"`
	FieldType  FieldType `yaml:"fieldType"`
	PrimaryKey bool      `yaml:"primaryKey"`
}

type View struct {
	Name     string      `yaml:"name"`
	Config   *ViewConfig `yaml:"config"`
	ViewType ViewType    `yaml:"type"`
	Query    *Query      `yaml:"query"`
}

type ViewConfig struct {
	EnableRequestValidation bool       `yaml:"enableRequestValidation"`
	MultiInput              bool       `yaml:"multiInput"`
	MultiReturn             bool       `yaml:"multiReturn"`
	ReturnType              ReturnType `yaml:"returnType"`
}

type Query struct {
	Fields []*QueryField `yaml:"fields"`
}

type QueryField struct {
	FieldName  string     `yaml:"fieldName"`
	SearchType SearchType `yaml:"searchType"`
	IsOptional bool       `yaml:"optional"`
}
