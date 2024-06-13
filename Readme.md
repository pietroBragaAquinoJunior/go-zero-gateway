goctl rpc protoc ./protos/zrpc.proto --go_out=./common/pb --go-grpc_out=./common/pb --zrpc_out=./zrpc

protoc --include_imports --proto_path=./protos --descriptor_set_out=common/pb/zrpc_descriptor.pb zrpc.proto