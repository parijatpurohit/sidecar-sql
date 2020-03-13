cd $GOPATH/src
mkdir -p github.com/hello/generated
protoc --go_out=github.com/parijatpurohit/sidecar-sql/generated/ github.com/parijatpurohit/sidecar-sql/pb/storage/User_schema.proto
protoc --proto_path=github.com/parijatpurohit/sidecar-sql/pb/storage/ --go_out=github.com/parijatpurohit/sidecar-sql/generated/ github.com/parijatpurohit/sidecar-sql/pb/storage/User_FindByRollAndName.proto