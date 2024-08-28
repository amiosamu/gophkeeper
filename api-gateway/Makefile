proto:
	protoc internal/services/pb/*.proto --go_out=. --go_opt=paths=internal/services/pb/ --go-grpc_out=. --go-grpc_opt=paths=source_relative
	protoc internal/query/pb/*.proto --go_out=. --go_opt=paths=internal/services/pb/ --go-grpc_out=. --go-grpc_opt=paths=source_relative
	protoc internal/auth/pb/*.proto --go_out=. --go_opt=paths=internal/services/pb/ --go-grpc_out=. --go-grpc_opt=paths=source_relative
	protoc internal/command/pb/*.proto --go_out=. --go_opt=paths=internal/services/pb/ --go-grpc_out=. --go-grpc_opt=paths=source_relative

server:
	go run cmd/main.go