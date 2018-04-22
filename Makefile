build:
	go build .

clean:
	go clean .

ci:
	fmt \
	clean \
	init \
	build \
	test

deploy:
	gcloud app deploy app.yaml \cron.yaml

fmt:
	go fmt .

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
	ci \
	deploy \
	fmt \
	init \
	run \
	test \
