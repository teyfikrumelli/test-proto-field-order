.PHONY: build run generate-proto

build:
	docker build -t test-proto-field-order .

run: build
	docker run --rm test-proto-field-order

generate-proto:
	docker run --rm -v `pwd`:/defs namely/protoc-all -f message_v1.proto -l go -o ./out
	docker run --rm -v `pwd`:/defs namely/protoc-all -f message_v2.proto -l go -o ./out
