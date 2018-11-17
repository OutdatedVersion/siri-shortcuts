default: build

clean:
	rm -rf build/

build: clean
	mkdir build/
	go build -o build/server

installDependencies:
	go get -v -u github.com/dgrr/fastws
	go get -v -u github.com/fasthttp/router
	go get -v -u github.com/valyala/fasthttp
