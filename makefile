GO=$(or ${RUNT_GO_VERSION},go)
export PATH := $(shell ${GO} env GOROOT)/bin:$(PATH)

mod:
	@echo "---> make mod <---"
	go mod tidy
	go mod download

build:
	@echo "---> make build <---"
	./build.sh

test:
	@echo "---> make test <---"
	go test

proto:
	@echo "---> make proto <---"
	./build_proto.sh

mock:
	@echo "---> make mock <---"
	mockgen -source=./pkg/executor/funcmanager/getter.go -destination=./pkg/mock/executor/funcmanager/getter.go -package=mock
