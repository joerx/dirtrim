OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)

build:
	mkdir -p out
	GOOS=$(OS) GOARCH=$(ARCH) go build -o out/dirtrim .

clean:
	rm -rf out
