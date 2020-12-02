set -e
go run code_generator/main.go -gentype=init -basepath=$1 -gopath=$GOPATH
go run code_generator/main.go -gentype=proto -basepath=$1 -gopath=$GOPATH
go run code_generator/main.go -gentype=protogen -basepath=$1 -gopath=$GOPATH
go run code_generator/main.go -gentype=storage -basepath=$1 -gopath=$GOPATH
go run code_generator/main.go -gentype=translator -basepath=$1 -gopath=$GOPATH
go run code_generator/main.go -gentype=server -basepath=$1 -gopath=$GOPATH