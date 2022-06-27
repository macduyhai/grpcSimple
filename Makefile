gen-pr:
	protoc simplepb/simplepb.proto --go_out=plugins=grpc:.