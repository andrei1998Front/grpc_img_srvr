compile:
	protoc -I ./pkg/proto ./pkg/proto/grpc_img_serv.proto --go_out=./pkg/proto/ --go_opt=paths=source_relative --go-grpc_out=./pkg/proto/ --go-grpc_opt=paths=source_relative

run:
	go run ./cmd/gis/main.go --config=./config/config.yaml

build:
	go build ./cmd/gis

test:
	mkdir internal\storage\test_imgs_dir
	go test ./internal/storage -cover
	rmdir /s /q internal\storage\test_imgs_dir
	go test ./internal/services
