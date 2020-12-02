package main

import (
	"flag"
	"log"

	"github.com/parijatpurohit/sidecar-sql/lib/config"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate"
)

const (
	GenTypeProto      = "proto"
	GenTypeProtogen   = "protogen"
	GenTypeServer     = "server"
	GenTypeStorage    = "storage"
	GenTypeTranslator = "translator"
	GenTypeInit       = "init"
	GenTypeAll        = "all"

	GenTypeDesc = "generation type. ex: proto, protogen, server, storage, translator, all"
)

func main() {
	config.InitStorageConfig()
	flag.Parse()
	genType := *config.GetFlags()[config.GenTypeFlag]

	if genType == GenTypeInit || genType == GenTypeAll {
		log.Printf("*********SETTING UP INITIAL DIRECTORY********")
		generate.Setup()
		log.Printf("*********INIT COMPLETE********")
	}
	if genType == GenTypeProto || genType == GenTypeAll {
		log.Printf("*********GENERATING PROTOBUF FILES********")
		generate.Proto()
		log.Printf("*********PROTOBUF GENERATION COMPLETE********")
	}

	if genType == GenTypeProtogen || genType == GenTypeAll {
		log.Printf("*********GENERATING PROTOBUF OUTPUT FILES********")
		generate.Protogen()
		log.Printf("*********PROTOBUF OUTPUT GENERATION COMPLETE********")
	}

	if genType == GenTypeStorage || genType == GenTypeAll {
		log.Printf("*********GENERATING STORAGE FILES*********")
		generate.Storage()
		log.Printf("*********STORAGE GENERATION COMPLETE********")

	}
	if genType == GenTypeTranslator || genType == GenTypeAll {
		log.Printf("*********GENERATING TRANSLATOR FILES*********")
		generate.Translator()
		log.Printf("*********TRANSLATOR GENERATION COMPLETE*********")
	}
	if genType == GenTypeServer || genType == GenTypeAll {
		log.Printf("*********GENERATING SERVER FILES*********")
		generate.Server()
		log.Printf("*********SERVER GENERATION COMPLETE*********")
	}
}
