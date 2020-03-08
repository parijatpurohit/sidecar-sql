cd $GOPATH/src
mkdir -p github.com/hello/generated
protoc --proto_path=github.com/hello/pb/storage --go_out=plugins=grpc:github.com/hello/generated/ github.com/hello/pb/storage/*.proto