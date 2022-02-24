GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
ROOT_PATH=../..
APP_NAME=gateway
APP_BIN_PATH=bin/$(APP_NAME)
API_PROTO_PATH=$(ROOT_PATH)/api/$(APP_NAME)
API_PROTO_FILES=$(shell find ../../api/gateway -name *.proto)
APP_PROTO_FILES=$(shell find ./ -name *.proto)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=$(API_PROTO_PATH) \
	       --proto_path=$(ROOT_PATH)/third_party \
 	       --go_out=paths=source_relative:$(API_PROTO_PATH) \
 	       --go-http_out=paths=source_relative:$(API_PROTO_PATH) \
 	       --go-grpc_out=paths=source_relative:$(API_PROTO_PATH) \
 	       --openapi_out=paths=source_relative:$(API_PROTO_PATH) \
			--openapiv2_out $(API_PROTO_PATH) \
			--openapiv2_opt logtostderr=true \
			--openapiv2_opt json_names_for_fields=false \
	       $(API_PROTO_FILES)

.PHONY: error
# generate error proto
error:
	protoc --proto_path=. \
	       --proto_path=$(API_PROTO_PATH) \
	       --proto_path=$(ROOT_PATH)/third_party \
 	       --go_out=paths=source_relative:$(API_PROTO_PATH) \
		   --go-errors_out=paths=source_relative:$(API_PROTO_PATH) \
	       $(API_PROTO_FILES)

.PHONY: validate
# generate validate proto
validate:
	protoc --proto_path=. \
	       --proto_path=$(API_PROTO_PATH) \
	       --proto_path=$(ROOT_PATH)/third_party \
 	       --go_out=paths=source_relative:$(API_PROTO_PATH) \
		   --validate_out=lang=go,paths=source_relative:$(API_PROTO_PATH) \
	       $(API_PROTO_FILES)

.PHONY: config
# generate conf proto
config:
	protoc --proto_path=. \
	       --proto_path=$(ROOT_PATH)/third_party \
 	       --go_out=paths=source_relative:. \
	       $(APP_PROTO_FILES)

.PHONY: proto
# generate all proto
proto: api error validate config

.PHONY: wire
# wire
wire:
	cd cmd && wire check && wire

.PHONY: dao
# dao
dao:
	gentool -dsn "root:root@tcp(127.0.0.1:3306)/url_shorter?charset=utf8mb4&parseTime=True&loc=Local" -tables "" -outPath "internal/data/dao/query" -modelPkgName "internal/data/dao/model"

.PHONY: build
# build
build:
	mkdir -p bin/ && cd cmd && go build -ldflags "-X main.Version=$(VERSION)" -o ../$(APP_BIN_PATH) ./...

.PHONY: run
# run
run:
	./$(APP_BIN_PATH) -conf ./config

.PHONY: air
# run by air
air:
	air

.PHONY: clean
# clean
clean:
	rm -rf ./bin

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
