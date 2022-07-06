# share proto

### install protoc

1. download protoc at: https://github.com/protocolbuffers/protobuf/releases

### install go grpc plugin

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

```

#### gen message protobuf

---

```shell
protoc --go_out=.. protobuf/message/*.proto
```

### gen service protobuf

---

```shell
protoc --proto_path=protobuf --go-grpc_out=.. protobuf/rpc/*.proto
```
