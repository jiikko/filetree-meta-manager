.PHONY: build dist test clean

PLATFORMS = linux darwin
ARCHITECTURES = amd64 arm64
VERSION ?= unknown

build:
	go build -ldflags="-X main.version=$(VERSION)" -o bin/ ./cmd/...

dist:
ifndef VERSION
	$(error VERSION is not set)
endif
	@for platform in $(PLATFORMS); do \
			for arch in $(ARCHITECTURES); do \
					GOOS=$$platform GOARCH=$$arch go build -ldflags="-X main.version=$(VERSION)" -o dist/filetree_dumper-$$platform-$$arch ./cmd/...; \
			done \
	done

test:
	go test ./...

bench:
	go test -bench=. -benchmem -run=^# ./...

clean:
	rm -rf bin dist
