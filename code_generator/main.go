package main

import (
	"log"
	"os"

	"github.com/parijatpurohit/sidecar-sql/lib/config"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate"
	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"
)

const (
	GenTypeProto      = "proto"
	GenTypeServer     = "server"
	GenTypeStorage    = "storage"
	GenTypeTranslator = "translator"
	GenTypeAll        = "all"
)

func main() {
	args := os.Args[1:]
	config.InitStorageConfig()
	if len(args) == 0 {
		panic("invalid arguments. please try again")
	}
	if utils.ContainsAnyStr(args, GenTypeProto, GenTypeAll) != -1 {
		log.Printf("*********GENERATING PROTOBUF FILES********")
		generate.Proto()
		log.Printf("*********PROTOBUF GENERATION COMPLETE********")
	}
	if utils.ContainsAnyStr(args, GenTypeStorage, GenTypeAll) != -1 {
		log.Printf("*********GENERATING STORAGE FILES*********")
		generate.Storage()
		log.Printf("*********STORAGE GENERATION COMPLETE********")

	}
	if utils.ContainsAnyStr(args, GenTypeTranslator, GenTypeAll) != -1 {
		log.Printf("*********GENERATING TRANSLATOR FILES*********")
		generate.Translator()
		log.Printf("*********TRANSLATOR GENERATION COMPLETE*********")
	}
	if utils.ContainsAnyStr(args, GenTypeServer, GenTypeAll) != -1 {
		log.Printf("*********GENERATING SERVER FILES*********")
		generate.Server()
		log.Printf("*********SERVER GENERATION COMPLETE*********")
	}
}
