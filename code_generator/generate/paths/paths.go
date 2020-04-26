package paths

const (
	// file paths
	EntityProtoPath = "%s/pb/storage/%s_schema.proto"
	ServerProtoPath = "%s/pb/storage/server.proto"
	ViewProtoPath   = "%s/pb/storage/%s_%s_View.proto"

	SchemaFilePath = "%s/go/storage/%s/models/schema.go"
	QueryFilePath  = "%s/go/storage/%s/models/%s.go"

	ProtoTemplatePath   = "code_generator/templates/proto"
	StorageTemplatePath = "code_generator/templates/storage"
	StorageConfigPath   = "configuration/storage"
	GeneratedFilePath   = "zz_generated"
	StorageOutputPath   = "go/storage"
	ModelsOutputPath    = "/models"

	// file names
	ProtoSchemaTemplateFile = "schema.tgo"
	ProtoServerTemplateFile = "server.tgo"
	ProtoViewTemplateFile   = "view.tgo"

	ModelSchemaTemplateFile = "views/schema.tgo"
	ModelQueryTemplateFile  = "views/query.tgo"
	CommonConfigFileName    = "common.yaml"
	SQLConfigFileName       = "sql/sql.yaml"
)
