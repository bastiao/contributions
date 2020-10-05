usage:
	@echo ""
	@echo "Task                 : Description"
	@echo "-----------------    : -------------------"
	@echo "make install         : Install all necessary dependencies"
	@echo "make build           : Generate production build for current OS"
	@echo "make test            : Execute unit test suite"
	@echo ""


VERSION = "v0.0.1"

install:
	cp bin/contributions $$DESTDIR

build:
	go get -v -t -d ./...
	go build -o bin/contributions cmd/*

docker:
	@docker build -t bastiao/contributions:$(VERSION) . \
	&& docker tag -f bastiao/contributions:$(VERSION) bastiao/contributions::latest

publish: build
	@docker push bastiao/contributions$(VERSION) \
	&& docker push bastiao/contributions:latest	

run:
	go run cmd/*.go $$PHA_ARGS


test:
	go test cmd/*.go
	go test sourceCode/*.go
	go test config/*.go
	go test jenkins/*.go
	go test sourceCode/*.go


clean:
	rm -rf bin/*