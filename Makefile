
.PHONY: build
build: 
	@docker build --ssh default -t $(IMAGE):latest .

.PHONY: build-local
build-local: datalayer/postgres/queries/db.go
	go build -buildvcs=false -o ./bin/server ./cmd/server

.PHONY: test-local
test-local: build-local
	go test ./...

## Code generation
.PHONY: setup
setup:
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.14.0
	@apt install -y protobuf-compiler

datalayer/postgres/queries/db.go: datalayer/postgres/sqlc.yaml datalayer/postgres/queries/*.sql datalayer/postgres/migrations/*.sql
	@echo "Compiling postgres sqlc module..."
	@sqlc generate -f ./datalayer/postgres/sqlc.yaml


protobuf/service_grpc.pb.go: protobuf/service.proto
	@echo "Compiling protobufs..."
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./protobuf/service.proto

## Other tools

hasher:
	go build -o ./bin/hasher ./cmd/hasher