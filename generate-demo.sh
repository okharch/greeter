#/bin/sh
# Change to the script directory
SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"
cd "$SCRIPT_DIR"
# make sure protoc installed
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
# remove previous artifacts
rm -rf github.com *.pb.go
protoc --go_out=. --go-grpc_out=. greeter.proto
mv github.com/okharch/greeter/* .
rm -rf github.com
echo make sure port 50051 is free
kill `lsof -i :50051 -t`
go run cmd/greeting-server.go &
echo waiting server
sleep 1
go run cmd/client.go
echo killing grpc greeting server...
kill `lsof -i :50051 -t`

