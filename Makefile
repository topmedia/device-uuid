VERSION := $(shell git describe --tags --always --dirty)

.PHONY: bin bin/device-uuid

all: bin/device-uuid bin/device-uuid.exe

clean:
	rm -f bin/device-uuid bin/device-uuid.exe

windows: bin/device-uuid.exe

bin/device-uuid.exe:
	GOOS=windows go build -i -v -o $@ -ldflags "-X main.Version=$(VERSION)" ./cmd/device-uuid

bin:
	mkdir -p bin

bin/device-uuid: bin
		go build -i -v -o $@ -ldflags "-X main.Version=$(VERSION)" ./cmd/device-uuid
