build:
	go build main.go

clean:
	go clean .

run:
	dev_appserver.py app.yaml

init:
	go get -u -d google.golang.org/appengine/...
	go get -u -d cloud.google.com/go/...

.PHONY: \
	build \
	clean \
	run \
	init
