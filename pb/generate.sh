cd $GOPATH/src
mkdir -p github.com/hello/generated
protoc --go_out=github.com/hello/generated/ github.com/hello/pb/storage/user/schema.proto
protoc --proto_path=github.com/hello/pb/storage/user/ --go_out=github.com/hello/generated/ github.com/hello/pb/storage/user/FindByRollAndName.proto
