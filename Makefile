# Makefile

.PHONY: generate-user-proto
generate-user-proto:
	protoc -I=. \
	    --go_out . --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    protos/v1/user/user.proto

.PHONY: generate-post-proto
generate-post-proto:
	protoc -I=. \
	    --go_out . --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    protos/v1/post/post.proto
