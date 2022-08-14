folder: ## create folder in project
	@mkdir -p cmd
	@mkdir -p cmd/worker
	@mkdir -p cmd/server
	@mkdir -p config
	@mkdir -p db
	@mkdir -p docs
	@mkdir -p genproto
	@mkdir -p internal
	@mkdir -p internal/adapter
	@mkdir -p internal/facade
	@mkdir -p internal/api
	@mkdir -p internal/common
	@mkdir -p internal/dto
	@mkdir -p internal/helper
	@mkdir -p internal/registry
	@mkdir -p internal/service
	@mkdir -p internal/usecase
	@mkdir -p internal/utils
	@mkdir -p thirdparty
	@mkdir -p proto
	@mkdir -p sql

1:
	protoc -I ./proto \
   	--go_out ./proto --go_opt paths=source_relative \
   	--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
   	--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
   	./proto/api/api.proto

gen:
	protoc \
		-I thirdparty/googleapis \
		--go_out=./genproto --go_opt=paths=source_relative \
    	--go-grpc_out=./genproto --go-grpc_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		proto/*.proto

get-grpc:
	go get \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc \

gen1:
	protoc -I . \
		-I thirdparty/googleapis \
		--plugin=protoc-gen-grpc-gateway=$GOPATH/bin/protoc-gen-grpc-gateway \
		--grpc-gateway_out ./genproto \
    	--grpc-gateway_opt logtostderr=true \
    	--grpc-gateway_opt paths=source_relative \
		--go_out=./genproto --go_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
    	--go-grpc_out=./genproto --go-grpc_opt=paths=source_relative \
    	proto/*.proto

install:
	go install \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    	google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc

buildlinux:
	docker build -t simpson . --platform=linux/amd64

deps:
	docker-compose -f docker-compose.yml up -d

build:
	docker build -t simpson .

run:
	docker run -it -m 256m --cpus=0.5  --name=simpson -p 8080:8080 -d simpson

loadtest:
	locust -f index.py