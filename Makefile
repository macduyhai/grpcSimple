gen-pr:
	protoc proto/hello.proto --go_out=plugins=grpc:.