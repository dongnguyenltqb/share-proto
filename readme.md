# share proto

## gen message proto

protoc --go_out=plugins=grpc:.. protobuf/message/*.proto

## gen service proto

protoc --proto_path=protobuf --go_out=plugins=grpc:.. protobuf/rpc/*.proto
