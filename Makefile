

build:
	GOPROXY=https://goproxy.cn go build -o bin/api

install:
	go install

lint:
	golangci-lint run ./...

proto:
	protoc -I ./third_party \
	-I ./api/v1 \
	--openapiv2_out ./docs --openapiv2_opt logtostderr=true \
	--openapiv2_opt json_names_for_fields=false \
	--go_out ./api/v1 --go_opt=paths=source_relative \
	--go-gin_out ./api/v1 --go-gin_opt=paths=source_relative \
	./api/v1/*.proto
	protoc-go-inject-tag -input=./api/v1/hello.pb.go