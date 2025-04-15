BIN_DIR := ./bin
COMPARE_CMD := ./cmd/compare
FS_CMD := ./cmd/fs

BIN_NAME := readDB
COMPARE_NAME := compareDB
FS_NAME := compareFS

all: readDB compareDB compareFS

readDB:
	go build -o $(BIN_DIR)/$(BIN_NAME) ./cmd

compareDB:
	go build -o $(BIN_DIR)/$(COMPARE_NAME) $(COMPARE_CMD)

compareFS:
	go build -o $(BIN_DIR)/$(FS_NAME) $(FS_CMD)

build: all

clean:
	rm -f $(BIN_DIR)/$(BIN_NAME) $(BIN_DIR)/$(COMPARE_NAME) $(BIN_DIR)/$(FS_NAME)

json: readDB
	$(BIN_DIR)/$(BIN_NAME) -f assets/stolen_database.json

xml: readDB
	$(BIN_DIR)/$(BIN_NAME) -f assets/original_database.xml

fmt:
	go fmt ./...

lint:
	golint ./...

tests:
	go test ./...