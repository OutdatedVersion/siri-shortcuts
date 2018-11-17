default: build

clean:
	rm -rf build/

build: clean
	mkdir build/
	go build -o build/client

installDependencies:
	go get -v -u github.com/dgrr/fastws
