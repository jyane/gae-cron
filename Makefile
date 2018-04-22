all: \
	clean \
	init \
	build \
	test

build:
	go build .

clean:
	go clean .

init:
	go get -u google.golang.org/appengine/...
	go get -u cloud.google.com/go/...

run:
	dev_appserver.py app.yaml

test:
	go test .

.PHONY: \
	build \
	clean \
	run \
	init
