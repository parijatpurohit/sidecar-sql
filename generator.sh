mkdir -p $GOPATH/src/github.com/parijatpurohit/sidecar-sql/zz_generated/pb/storage
go run code_generator/main.go proto
./code_generator/scripts/proto.sh "$1"
go run code_generator/main.go storage