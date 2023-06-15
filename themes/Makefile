all:
	make -C ./adminlte
	make fmt

fmt:
	GO111MODULE=off go fmt ./...
	GO111MODULE=off goimports -l -w .