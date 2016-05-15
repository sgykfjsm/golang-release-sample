VERSION := $(shell cat VERSION)
DIST_DIRS := find * -type d -exec

deps:
	go get -v github.com/Masterminds/glide
	glide install

deps-release:
	go get -v github.com/mitchellh/gox

build:
	gox -verbose -output="dist/{{.Dir}}_${VERSION}_{{.OS}}_{{.Arch}}"

dist: build
	ghr -t $GITHUB_TOKEN -u $CIRCLE_PROJECT_USERNAME -r $CIRCLE_PROJECT_REPONAME dist/

test:
	go vet $(glide novendor)
	go test -v -race -cover $(glide novendor)

clean:
	rm -frv dist
