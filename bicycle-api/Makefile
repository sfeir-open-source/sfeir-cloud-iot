init:
	cp .env.dist .env

coding-style:
	go mod tidy
	gofmt -l -w ./
	go vet ./...
	goimports -l -w ./
	golint -set_exit_status ./...
