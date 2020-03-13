cd $GOPATH/src
mkdir -p github.com/hello/generated
protoc --go_out=github.com/parijatpurohit/sidecar-sql/generated/ github.com/parijatpurohit/sidecar-sql/pb/storage/user/schema.proto
protoc --proto_path=github.com/parijatpurohit/sidecar-sql/pb/storage/user/ --go_out=github.com/parijatpurohit/sidecar-sql/generated/ github.com/parijatpurohit/sidecar-sql/pb/storage/user/FindByRollAndName.proto