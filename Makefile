PREFIX=/usr/local
RUNTIME_GOPATH=$(GOPATH):`pwd`
VERSION=`git tag | tail -n 1`
GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

j2y:	main.go dependencies
	rm -rf linux-*/ darwin-*/ windows-*/

	gox -os=linux -arch=386 -output build/linux-386/{{.Dir}}
	cd build/linux-386/ && zip -r ../j2y-linux-386.zip . && cd ../../

	gox -os=linux -arch=amd64 -output build/linux-amd64/{{.Dir}}
	cd build/linux-amd64/ && zip -r ../j2y-linux-amd64.zip . && cd ../../

	gox -os=darwin -arch=amd64 -output build/darwin-amd64/{{.Dir}}
	cd build/darwin-amd64/ && zip -r ../j2y-darwin-amd64.zip . && cd ../../

	gox -os=windows -arch=386 -output build/windows-386/{{.Dir}}
	cd build/windows-386/ && zip -r ../j2y-windows-386.zip . && cd ../../

	gox -os=windows -arch=amd64 -output build/windows-amd64/{{.Dir}}
	cd build/windows-amd64/ && zip -r ../j2y-windows-amd64.zip . && cd ../../

dependencies:
	go get github.com/codegangsta/cli
	go get github.com/ghodss/yaml
	go get github.com/bitly/go-simplejson

install: j2y
	install -m 755 j2y $(DESTDIR)$(PREFIX)/bin/

clean:
	rm -rf ./j2y build/
