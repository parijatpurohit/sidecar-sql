cd $GOPATH/src

# generate for argument go or golang
if [ "$1" == "go" ] || [ "$1" == "golang" ]
then
  mkdir -p github.com/sidecar-storage/zz_generated/go
  protoc --proto_path=github.com/sidecar-storage/pb/storage \
  --go_out=plugins=grpc:github.com/sidecar-storage/zz_generated/go \github.com/sidecar-storage/pb/storage/*.proto

# generate javascript clients
elif [ "$1" == "node" ] || [ "$1" == "nodejs" ] || [ "$1" == "javascript" ] || [ "$1" == "js" ]
then
  mkdir -p github.com/sidecar-storage/zz_generated/js
  protoc --proto_path=github.com/sidecar-storage/pb/storage \
  --js_out=import_style=commonjs,binary:github.com/sidecar-storage/zz_generated/js github.com/sidecar-storage/pb/storage/*.proto
fi