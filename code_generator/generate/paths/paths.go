package paths

const (
	// file paths
	EntityProtoPath   = "%s/pb/storage/%s_schema.proto"
	ServerProtoPath   = "%s/pb/storage/server.proto"
	ViewProtoPath     = "%s/pb/storage/%s_%s_View.proto"
	TemplatePath      = "templates/proto"
	StorageConfigPath = "configuration/storage"
	GeneratedFilePath = "zz_generated"

	// file names
	SchemaTemplateFile = "schema.tgo"
	ServerTemplateFile = "server.tgo"
	ViewTemplateFile   = "view.tgo"
	CommonFileName     = "common.yaml"
)
