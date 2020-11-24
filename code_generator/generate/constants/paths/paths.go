package paths

const (
	OutputGoPath = "%s/go"

	BaseImportPath         = "github.com/parijatpurohit/sidecar-sql/"
	ModelsRelativePath     = "/go/storage/%s/models"
	ViewsRelativePath      = "/go/storage/%s/views"
	ProtoRelativePath      = "/go/protogen"
	TranslatorRelativePath = "/translator/%s"
	ServerRelativePath     = OutputGoPath + "/server"
	HandlersRelativePath   = ServerRelativePath + HandlersOutputPath
	ModelsImportPath       = BaseImportPath + GeneratedFilePath + ModelsRelativePath
	ViewsImportPath        = BaseImportPath + GeneratedFilePath + ViewsRelativePath
	ProtoImportPath        = BaseImportPath + GeneratedFilePath + ProtoRelativePath
	HandlerImportPath      = BaseImportPath + HandlersRelativePath
	TranslatorImportPath   = BaseImportPath + OutputGoPath + TranslatorRelativePath

	UtilsRelativePath   = "utils"
	TimestampImportPath = "github.com/golang/protobuf/ptypes/timestamp"
	PTypesImportPath    = "github.com/golang/protobuf/ptypes"
	SqlConnImportPath   = BaseImportPath + "lib/sqlconn"
	ConfigImportPath    = BaseImportPath + "code_generator/config"

	// file paths
	EntityProtoPath = "%s/pb/storage/%s_schema.proto"
	ServerProtoPath = "%s/pb/storage/server.proto"
	ViewProtoPath   = "%s/pb/storage/%s_%s_View.proto"

	SchemaFilePath         = OutputGoPath + "/storage/%s/models/schema.go"
	QueryFilePath          = OutputGoPath + "/storage/%s/models/%s.go"
	ViewsFilePath          = OutputGoPath + "/storage/%s/views/%s.go"
	BaseTranslatorFilePath = OutputGoPath + TranslatorRelativePath + "/base.go"
	ViewTranslatorFilePath = OutputGoPath + TranslatorRelativePath + "/%s.go"

	BaseHanderFilePath  = HandlersRelativePath + "/handlers.go"
	TableHanderFilePath = HandlersRelativePath + "/%s_handlers.go"
	ServerFilePath      = ServerRelativePath + "/serve.go"

	ViewDefsFileName = "views"

	ProtoTemplatePath      = "code_generator/templates/proto"
	StorageTemplatePath    = "code_generator/templates/storage"
	ServerTemplatePath     = "code_generator/templates/server"
	TranslatorTemplatePath = "code_generator/templates/translator"

	ConfigBasePath    = "configuration"
	StorageConfigPath = "configuration/storage"

	GeneratedFilePath    = "zz_generated"
	StorageOutputPath    = "go/storage"
	TranslatorOutputPath = "go/translator"
	ServerOutputPath     = "go/server"
	ModelsOutputPath     = "/models"
	ViewsOutputPath      = "/views"

	HandlersOutputPath = "/handlers"

	// file names
	ProtoSchemaTemplateFile = "schema.tgo"
	ProtoServerTemplateFile = "server.tgo"
	ProtoViewTemplateFile   = "view.tgo"

	BaseHandlerTemplateFile  = "handlers/base_handler.tgo"
	TableHandlerTemplateFile = "handlers/table_handler.tgo"
	ServerTemplateFile       = "serve.tgo"

	BaseTranslatorTemplateFile = "basetranslator.tgo"
	ViewTranslatorTemplateFile = "viewtranslator.tgo"

	ModelSchemaTemplateFile = "views/schema.tgo"
	ModelQueryTemplateFile  = "views/query.tgo"
	CreateViewTemplateFile  = "views/create.tgo"
	UpdateViewTemplateFile  = "views/update.tgo"
	DeleteViewTemplateFile  = "views/delete.tgo"
	ReadViewTemplateFile    = "views/read.tgo"
	ViewDefTemplateFile     = "views/viewsdef.tgo"
	CommonConfigFileName    = "common.yaml"
	SQLConfigFileName       = "sql/sql.yaml"
)
