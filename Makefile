.PHONY: all clean build docker docker-build docker-push

GORUN = go run
GOBUILD = go build

all: remove build

remove:
	rm -rf bin

build: remove
	$(GOBUILD) -o bin/api cmd/api/main.go

docker-build:
	test $(DOCKERREPO)
	docker build . -t $(DOCKERREPO)

docker-push:
	test $(DOCKERREPO)
	docker push $(DOCKERREPO)

docker: docker-build docker-push
