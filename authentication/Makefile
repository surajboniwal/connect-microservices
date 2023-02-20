proto:
	protoc --proto_path=internal/pb internal/pb/*.proto --go_out=. --go-grpc_out=.

build:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o authentication cmd/*

remove:
	rm -rf authentication

createdb:
	docker exec -it connect-authentication-db-1 createdb --username=surajboniwal --owner=surajboniwal socialmedia

dropdb:
	docker exec -it connect-authentication-db-1 dropdb socialmedia

migrate-up:
	migrate -path internal/db/migration -database "postgresql://surajboniwal:suraj1335@localhost:3001/socialmedia?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/db/migration -database "postgresql://surajboniwal:suraj1335@localhost:3001/socialmedia?sslmode=disable" -verbose down