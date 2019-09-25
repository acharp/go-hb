ensure:
	dep ensure

build:
	@go build -o bin/go-hb cmd/*.go


run: build
	@./bin/go-hb $(ARGS)
