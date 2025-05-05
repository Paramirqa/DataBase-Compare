BIN_DIR := ./bin
COMPARE_CMD := ./cmd/compare
READING_CMD := ./cmd/reading
FS_CMD := ./cmd/fs

BIN_NAME := readDB
COMPARE_NAME := compareDB
FS_NAME := compareFS

VERSION := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

all: readDB compareDB

readDB:
	go build -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)" -o $(BIN_DIR)/$(BIN_NAME) $(READING_CMD)

compareDB:
	go build -o $(BIN_DIR)/$(COMPARE_NAME) $(COMPARE_CMD)

compareFS:
	go build -o $(BIN_DIR)/$(FS_NAME) $(FS_CMD)

build: all

clean:
	rm -f $(BIN_DIR)/$(BIN_NAME) $(BIN_DIR)/$(COMPARE_NAME) $(BIN_DIR)/$(FS_NAME) cover.out coverage.html

rebuild: clean
	$(MAKE) build

json: readDB
	$(BIN_DIR)/$(BIN_NAME) -f assets/stolen_database.json

xml: readDB
	$(BIN_DIR)/$(BIN_NAME) -f assets/original_database.xml

fmt:
	go fmt ./...

fmt-check:
	@UNFORMATTED=$$(gofmt -l .); \
	if [ -n "$$UNFORMATTED" ]; then \
		echo "❌ Неотформатированные файлы:"; \
		echo "$$UNFORMATTED"; \
		exit 1; \
	else \
		echo "✅ Все файлы отформатированы"; \
	fi

lint:
	staticcheck ./...

lint-readDB:
	staticcheck ./cmd/reading ./internal/...

lint-compareDB:
	staticcheck ./cmd/compare ./internal/...

lint-compareFS:
	staticcheck ./cmd/fs ./internal/...

tests:
	go test ./...

test-cover:
	go test -coverprofile=cover.out ./...

test-cover-html: test-cover
	go tool cover -html=cover.out -o coverage.html
	@if command -v xdg-open > /dev/null; then \
		xdg-open coverage.html; \
	elif command -v open > /dev/null; then \
		open coverage.html; \
	else \
		echo "coverage.html сгенерирован, но не удалось открыть автоматически."; \
	fi

run-compare: compareDB
	$(BIN_DIR)/$(COMPARE_NAME) --old assets/original_database.xml --new assets/stolen_database.json