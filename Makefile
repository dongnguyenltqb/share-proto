default: clean build

clean:
	@rm -R proto-gen

build:
	@protoc --go_out=.. protobuf/message/*.proto
	@protoc --proto_path=protobuf --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=.. protobuf/rpc/*.proto
