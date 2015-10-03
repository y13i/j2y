PREFIX=/usr/local
RUNTIME_GOPATH=$(GOPATH):`pwd`
VERSION=`git tag | tail -n 1`
GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

j2y:	main.go dependencies
	rm -rf linux-*/ darwin-*/ windows-*/

	gox -os=linux -arch=386 -output linux-386/{{.Dir}}
	zip -r linux-386.zip linux-386/

	gox -os=linux -arch=amd64 -output linux-amd64/{{.Dir}}
	zip -r linux-amd64.zip linux-amd64/

	gox -os=darwin -arch=386 -output darwin-386/{{.Dir}}
	zip -r darwin-386.zip darwin-386/

	gox -os=darwin -arch=amd64 -output darwin-amd64/{{.Dir}}
	zip -r darwin-amd64.zip darwin-amd64/

	gox -os=windows -arch=386 -output windows-386/{{.Dir}}
	zip -r windows-386.zip windows-386/

	gox -os=windows -arch=amd64 -output windows-amd64/{{.Dir}}
	zip -r windows-amd64.zip windows-amd64/

dependencies:
	go get github.com/codegangsta/cli
	go get github.com/ghodss/yaml
	go get github.com/bitly/go-simplejson

install: j2y
	install -m 755 j2y $(DESTDIR)$(PREFIX)/bin/

clean:
	rm -rf j2y *.tar.gz *.zip linux-*/ darwin-*/ windows-*/

package: clean j2y
	tar zcf j2y-$(VERSION)-${GOOS}-$(GOARCH).tar.gz ./j2y
