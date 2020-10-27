package paths

const (
	OutputGoPath = "%s/go"

	BaseImportPath         = "github.com/parijatpurohit/sidecar-sql/"
	ModelsRelativePath     = "/go/storage/%s/models"
	ViewsRelativePath      = "/go/storage/%s/views"
	ProtoRelativePath      = "/go/protogen"
	TranslatorRelativePath = "/translator/%s"
	ServerRelativePath     = OutputGoPath + "/server"
	HandersRelativePath    = ServerRelativePath + HandlersOutputPath
	ModelsImportPath       = BaseImportPath + GeneratedFilePath + ModelsRelativePath
	ViewsImportPath        = BaseImportPath + GeneratedFilePath + ViewsRelativePath
	ProtoImportPath        = BaseImportPath + GeneratedFilePath + ProtoRelativePath
	TranslatorImportPath   = BaseImportPath + OutputGoPath + TranslatorRelativePath

	UtilsRelativePath   = "utils"
	TimestampImportPath = "github.com/golang/protobuf/ptypes/timestamp"
	PTypesImportPath    = "github.com/golang/protobuf/ptypes"
	// file paths
	EntityProtoPath = "%s/pb/storage/%s_schema.proto"
	ServerProtoPath = "%s/pb/storage/server.proto"
	ViewProtoPath   = "%s/pb/storage/%s_%s_View.proto"

	SchemaFilePath         = OutputGoPath + "/storage/%s/models/schema.go"
	QueryFilePath          = OutputGoPath + "/storage/%s/models/%s.go"
	ViewsFilePath          = OutputGoPath + "/storage/%s/views/%s.go"
	BaseTranslatorFilePath = OutputGoPath + TranslatorRelativePath + "/base.go"
	ViewTranslatorFilePath = OutputGoPath + TranslatorRelativePath + "/%s.go"

	BaseHanderFilePath  = HandersRelativePath + "/handlers.go"
	TableHanderFilePath = HandersRelativePath + "/%s_handlers.go"

	ViewDefsFileName = "views"

	ProtoTemplatePath      = "code_generator/templates/proto"
	StorageTemplatePath    = "code_generator/templates/storage"
	ServerTemplatePath     = "code_generator/templates/server"
	TranslatorTemplatePath = "code_generator/templates/translator"
	StorageConfigPath      = "configuration/storage"
	GeneratedFilePath      = "zz_generated"
	StorageOutputPath      = "go/storage"
	TranslatorOutputPath   = "go/translator"
	ServerOutputPath       = "go/server"
	ModelsOutputPath       = "/models"
	ViewsOutputPath        = "/views"

	HandlersOutputPath = "/handlers"

	// file names
	ProtoSchemaTemplateFile = "schema.tgo"
	ProtoServerTemplateFile = "server.tgo"
	ProtoViewTemplateFile   = "view.tgo"

	BaseHandlerTemplateFile  = "handlers/base_handler.tgo"
	TableHandlerTemplateFile = "handlers/table_handler.tgo"
	ServerTemplateFile       = "server.tgo"

	BaseTranslatorTemplateFile = "basetranslator.tgo"
	ViewTranslatorTemplateFile = "viewtranslator.tgo"

	ModelSchemaTemplateFile = "views/schema.tgo"
	ModelQueryTemplateFile  = "views/query.tgo"
	CreateViewTemplateFile  = "views/create.tgo"
	UpdateViewTemplateFile  = "views/update.tgo"
	DeleteViewTemplateFile  = "views/delete.tgo"
	ViewDefTemplateFile     = "views/viewsdef.tgo"
	CommonConfigFileName    = "common.yaml"
	SQLConfigFileName       = "sql/sql.yaml"
)
