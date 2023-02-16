proto:
	protoc --proto_path=internal/pb internal/pb/*.proto --go_out=. --go-grpc_out=.

build:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o authentication cmd/*

remove:
	rm -rf authentication