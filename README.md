install protoc => see in site https://grpc.io/docs/languages/go/quickstart/ how should be installed


protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet/proto/greet.proto
protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/greet.proto 


go mod tidy