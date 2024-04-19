go get google.golang.org/grpc
protoc BookPb/Book/Book.proto --go-grpc_out=. --go_out=.
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

## !!!
For Grpc-Gateway setup..
create folder tools and file inside tools.go
// import (
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
    _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
    _ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
run go mod tidy

protoc -I . --grpc-gateway_out=BookPb/Book \
    --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    your/service/v1/your_service.proto

proxy server -> http://localhost:8081/v1/example/echo