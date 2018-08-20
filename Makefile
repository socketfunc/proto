PROTOC_INCLUDE := -I=".:${GOPATH}/src/github.com/socketfunc"
DIRS := gateway runtime store

protoc:
	@for dir in $(DIRS); do \
	( \
		protoc ${PROTOC_INCLUDE} --go_out=plugins=grpc:${GOPATH}/src $$dir/*.proto; \
	) \
	done
