# Путь к директории с main.go
CMD_DIR := ./cmd
BIN_NAME := readDB

# Build бинарника
build:
	go build -o $(BIN_NAME) $(CMD_DIR)

# Удаление бинарников
clean:
	rm -f $(BIN_NAME)

# Тест запуск с JSON-файлом
json: build
	./$(BIN_NAME) -f assets/stolen_database.json

# Тест запуск с XML-файлом
xml: build
	./$(BIN_NAME) -f assets/original_database.xml
fmt:
	go fmt ./...

lint:
	golint ./...

tests:
	go test ./tests