package paths

const (
	// file paths
	EntityProtoPath   = "%s/pb/storage/%s_schema.proto"
	ServerProtoPath   = "%s/pb/storage/server.proto"
	ViewProtoPath     = "%s/pb/storage/%s_%s_View.proto"
	TemplatePath      = "code_generator/templates/proto"
	StorageConfigPath = "configuration/storage"
	GeneratedFilePath = "zz_generated"
	StorageOutputPath = "go/storage"

	// file names
	ProtoSchemaTemplateFile = "schema.tgo"
	ProtoServerTemplateFile = "server.tgo"
	ProtoViewTemplateFile   = "view.tgo"
	CommonConfigFileName    = "common.yaml"
	SQLConfigFileName       = "sql/sql.yaml"
)
