package generate

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/gentranslator"
	generateUtils "github.com/parijatpurohit/sidecar-sql/code_generator/generate/utils"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

func Translator() {
	for _, conf := range config.GetAllStorage() {
		tableName := generateUtils.GetTableName(conf.Table, conf.Common.IsPlural)
		createTranslatorPath(tableName)
		log.Printf("generating base translator for: %s", tableName)
		gentranslator.GenerateBaseTranslator(conf)
		gentranslator.GenerateViewTranslator(conf)
	}
}

func createTranslatorPath(tableName string) {
	path := fmt.Sprintf("%s/%s/%s", paths.GeneratedFilePath, paths.TranslatorOutputPath, strings.ToLower(tableName))
	if err := os.MkdirAll(path, 0755); err != nil {
		log.Panic(err)
	}
}
