proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. pkg/pb/*.proto   

server:
	go run cmd/main.go