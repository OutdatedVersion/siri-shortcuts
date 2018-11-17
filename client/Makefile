default: build

clean:
	rm -rf build/

build: clean
	mkdir build/
	go build -o build/client

installDependencies:
	go get -u github.com/dgrr/fastws
