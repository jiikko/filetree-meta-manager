.PHONY: build dist test clean

PLATFORMS = linux darwin
ARCHITECTURES = amd64 arm64

build:
	go build -o bin/ ./cmd/...

dist:
	@for platform in $(PLATFORMS); do \
		for arch in $(ARCHITECTURES); do \
			GOOS=$$platform GOARCH=$$arch go build -o dist/kecy-$$platform-$$arch ./cmd/...; \
		done \
	done

test:
	go test ./...

bench:
	go test -bench=. -benchmem -run=^# ./...

clean:
	rm -rf bin dist
