package main

import (
	"log"
	"os"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate"
	"github.com/parijatpurohit/sidecar-sql/code_generator/utils"
)

const (
	GenTypeProto   = "proto"
	GenTypeServer  = "server"
	GenTypeStorage = "storage"
	GenTypeAll     = "all"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("invalid arguments. please try again")
	}
	if utils.ContainsAnyStr(args, GenTypeProto, GenTypeAll) != -1 {
		log.Printf("generating proto files")
		generate.Proto()
	}
	if utils.ContainsAnyStr(args, GenTypeStorage, GenTypeAll) != -1 {
		log.Printf("generating storage files")
		generate.Storage()
	}
	if utils.ContainsAnyStr(args, GenTypeServer, GenTypeAll) != -1 {
		log.Printf("generating server files")
		generate.Server()
	}

}
