package genproto

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/parijatpurohit/sidecar-sql/code_generator/generate/constants/paths"
	"github.com/parijatpurohit/sidecar-sql/lib/config"
)

const (
	protocCommand = "protoc --go_out=%s --proto_path=%s --go-grpc_out=%s %s/*.proto"
)

func GenerateProtoOutput() {
	basePath := *config.GetFlags()[config.ServiceBasePath]
	protoFilePath := fmt.Sprintf("%s/%s/%s", basePath, paths.GeneratedFilePath, paths.ProtobufFilePath)
	cmdStr := fmt.Sprintf(protocCommand, basePath, protoFilePath, basePath, protoFilePath)

	cmd := exec.Command("/bin/sh", "-c", cmdStr)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Panic(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	log.Println("Result: " + out.String())
}
