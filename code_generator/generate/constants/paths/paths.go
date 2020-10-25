package paths

const (
	OutputGoPath = "%s/go"

	BaseImportPath      = "github.com/parijatpurohit/sidecar-sql/"
	ModelsRelativePath  = "/go/storage/%s/models"
	ViewsRelativePath   = "/go/storage/%s/views"
	ProtoRelativePath   = "/go/protogen"
	ServerRelativePath  = OutputGoPath + "/server"
	HandersRelativePath = ServerRelativePath + HandlersOutputPath
	ModelsImportPath    = BaseImportPath + GeneratedFilePath + ModelsRelativePath
	ViewsImportPath     = BaseImportPath + GeneratedFilePath + ViewsRelativePath
	ProtoImportPath     = BaseImportPath + GeneratedFilePath + ProtoRelativePath

	// file paths
	EntityProtoPath = "%s/pb/storage/%s_schema.proto"
	ServerProtoPath = "%s/pb/storage/server.proto"
	ViewProtoPath   = "%s/pb/storage/%s_%s_View.proto"

	SchemaFilePath = OutputGoPath + "/storage/%s/models/schema.go"
	QueryFilePath  = OutputGoPath + "/storage/%s/models/%s.go"
	ViewsFilePath  = OutputGoPath + "/storage/%s/views/%s.go"

	HandersFilePath = HandersRelativePath + "/handlers.go"

	ViewDefsFileName = "views"

	ProtoTemplatePath   = "code_generator/templates/proto"
	StorageTemplatePath = "code_generator/templates/storage"
	ServerTemplatePath  = "code_generator/templates/server"
	StorageConfigPath   = "configuration/storage"
	GeneratedFilePath   = "zz_generated"
	StorageOutputPath   = "go/storage"
	ServerOutputPath    = "go/server"
	ModelsOutputPath    = "/models"
	ViewsOutputPath     = "/views"

	HandlersOutputPath = "/handlers"

	// file names
	ProtoSchemaTemplateFile = "schema.tgo"
	ProtoServerTemplateFile = "server.tgo"
	ProtoViewTemplateFile   = "view.tgo"

	HandlersTemplateFile = "handlers/handlers.tgo"
	ServerTemplateFile   = "server.tgo"

	ModelSchemaTemplateFile = "views/schema.tgo"
	ModelQueryTemplateFile  = "views/query.tgo"
	CreateViewTemplateFile  = "views/create.tgo"
	UpdateViewTemplateFile  = "views/update.tgo"
	DeleteViewTemplateFile  = "views/delete.tgo"
	ViewDefTemplateFile     = "views/viewsdef.tgo"
	CommonConfigFileName    = "common.yaml"
	SQLConfigFileName       = "sql/sql.yaml"
)
