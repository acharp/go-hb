ensure:
	dep ensure

build:
	@go build -o bin/go-hb cmd/*.go

run: build
	@./bin/go-hb $(ARGS)

zip:
	cp cmd/main.go deployment/cmd.go
	cp -R vendor deployment/
	cd deployment && zip -r go-hb.zip cmd.go vendor
