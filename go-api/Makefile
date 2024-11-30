ORG="gogotechhk"
REPO=gogovan-gin-sample-codes
COV_FILE=bin/cov-gogovan-gin-sample-codes.out
DOCKER_BUILD_COMMAND=podman
EXECUTABLE=app

swagger:
	swag init
build-image:
	$(DOCKER_BUILD_COMMAND) build  -t $(ORG)/$(REPO):latest .
sonarcloud-test:
	mkdir -p bin
	go test -short -coverprofile=$(COV_FILE) `go list ./... | grep -v vendor/`
	go tool cover -func=$(COV_FILE)
clean-sonarcloud-testfiles:
	rm -f $(COV_FILE)
	rm -f bin/$(EXECUTABLE)
unit-test:
	go test -v ./...
start-container:
	$(DOCKER_BUILD_COMMAND) start $(REPO)
stop-container:
	$(DOCKER_BUILD_COMMAND) stop $(REPO)
delete-image: clean-image
clean-image:
	- $(DOCKER_BUILD_COMMAND) rm $(REPO)
	$(DOCKER_BUILD_COMMAND) rmi $(ORG)/$(REPO):latest 
clean-logs:
	rm -f logs/*.log
delete-all: clean-all
clean-all: clean-logs
	- $(DOCKER_BUILD_COMMAND) rmi $(ORG)/$(REPO):latest
run-container:
	$(DOCKER_BUILD_COMMAND) run -d -p 8080:8080 --rm --name $(REPO) $(ORG)/$(REPO):latest
