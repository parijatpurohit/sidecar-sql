cd $GOPATH/src
mkdir -p github.com/sidecar-storage/generated
protoc --proto_path=github.com/sidecar-storage/pb/storage --go_out=plugins=grpc:github.com/sidecar-storage/generated/ github.com/sidecar-storage/pb/storage/*.proto