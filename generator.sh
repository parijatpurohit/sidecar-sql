rm -rf zz_generated/
mkdir -p zz_generated/pb/storage
go run code_generator/main.go proto
./code_generator/scripts/proto.sh "$1"
go run code_generator/main.go storage
go run code_generator/main.go server