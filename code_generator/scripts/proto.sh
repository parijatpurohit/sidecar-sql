cd $GOPATH/src

echo "generating go server and $1 client code"
# generate for argument go or golang
if [ "$1" == "go" ] || [ "$1" == "golang" ]
then
  mkdir -p github.com/parijatpurohit/sidecar-sql/zz_generated/go
  protoc --proto_path=github.com/parijatpurohit/sidecar-sql/zz_generated/pb/storage \
  --go_out=plugins=grpc:github.com/parijatpurohit/sidecar-sql/zz_generated/go \github.com/parijatpurohit/sidecar-sql/zz_generated/pb/storage/*.proto

# generate javascript clients
elif [ "$1" == "node" ] || [ "$1" == "nodejs" ] || [ "$1" == "javascript" ] || [ "$1" == "js" ]
then
  if [ "$2" == "" ]
  then
    echo "[ERROR]: nodejs setup requires destination for generated files"
    exit
  fi
  # clean up current temp location
  rm -rf github.com/parijatpurohit/sidecar-sql/zz_generated/js
  mkdir -p github.comparijatpurohit/sidecar-sql/zz_generated/js

  cd github.com/parijatpurohit/sidecar-sql/zz_generated/pb/storage
  grpc_tools_node_protoc --js_out=import_style=commonjs,binary:../../zz_generated/js \
  --grpc_out=../../zz_generated/js \
  --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` *.proto

  rm -rf "$2"
  mkdir -p "$2"
  mv ../../zz_generated/js "$2"

fi