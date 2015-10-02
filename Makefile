PREFIX=/usr/local
RUNTIME_GOPATH=$(GOPATH):`pwd`
VERSION=`git tag | tail -n 1`
GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

j2y:	main.go dependencies
	GOPATH=$(RUNTIME_GOPATH) go build -o j2y main.go
	zip j2y.zip j2y

dependencies:
	go get github.com/codegangsta/cli
	go get github.com/ghodss/yaml
	go get github.com/bitly/go-simplejson

install: j2y
	install -m 755 j2y $(DESTDIR)$(PREFIX)/bin/

clean:
	rm -f j2y *.tar.gz

package: clean j2y
	tar zcf j2y-$(VERSION)-${GOOS}-$(GOARCH).tar.gz ./j2y
