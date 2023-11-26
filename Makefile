
.PHONY: build
build:
	@docker build --ssh default -t $(IMAGE):latest .

.PHONY: build-local
build-local: build-comics
	@echo "building all servers"

.PHONY: build-comics
build-comics: pkg/repositories/comics/postgresstore/queries/db.go protobufs/gen/go/comics.pb.go pkg/repositories/login/postgresstore/queries/db.go
	@go build -buildvcs=false -o ./bin/server/comics ./cmd/server/comics


.PHONY: test-local
test-local: build-local
	@go test ./...

## Code generation
.PHONY: setup
setup:
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.14.0
	@apt install -y protobuf-compiler

pkg/repositories/comics/postgresstore/queries/db.go: pkg/repositories/comics/postgresstore/queries/*.sql pkg/repositories/comics/postgresstore/migrations/*.sql pkg/repositories/comics/postgresstore/sqlc.yaml
	@echo "Compiling comics repo's postgres sqlc module..."
	@sqlc generate -f pkg/repositories/comics/postgresstore/sqlc.yaml

pkg/repositories/login/postgresstore/queries/db.go: pkg/repositories/login/postgresstore/queries/*.sql pkg/repositories/login/postgresstore/migrations/*.sql pkg/repositories/login/postgresstore/sqlc.yaml
	@echo "Compiling login repo's postgres sqlc module..."
	@sqlc generate -f pkg/repositories/login/postgresstore/sqlc.yaml

protobufs/gen/go/comics.pb.go: protobufs/*.proto
	@echo "Compiling protobufs..."
	@cd protobufs; buf generate ; cd ..


## Other tools

hasher:
	go build -o ./bin/hasher ./cmd/hasher
