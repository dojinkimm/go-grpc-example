# Makefile

.PHONY: generate-proto
generate-proto:
	protoc -I=. \
	    --go_out . --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    user/user.proto
