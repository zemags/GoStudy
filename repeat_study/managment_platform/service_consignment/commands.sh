# protoc api/consignment.proto --go_out=plugins=grpc:pb
protoc --proto_path=. --go_out=pb --micro_out=pb api/consignment.proto