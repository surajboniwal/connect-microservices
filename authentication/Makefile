proto:
	protoc --proto_path=pkg/pb pkg/pb/*.proto --go_out=. --go-grpc_out=.

build:
	CGO_ENABLED=0 go build -o authentication cmd/*

remove:
	rm -rf authentication