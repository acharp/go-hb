ensure:
	dep ensure

build:
	@go build -o bin/go-hb cmd/*.go

run: build
	@./bin/go-hb $(ARGS)

# First change the package name, the main function signature and uncomment the net/http package then zip and it's good to deploy
zip:
	cp cmd/main.go deployment/cmd.go
	cp -R vendor deployment/
	cd deployment && zip -r go-hb.zip cmd.go vendor
