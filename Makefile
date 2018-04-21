build:
	go build .

clean:
	go clean .

run:
	dev_appserver.py app.yaml

init:
	go get -u google.golang.org/appengine/...
	go get -u cloud.google.com/go/...

.PHONY: \
	build \
	clean \
	run \
	init
