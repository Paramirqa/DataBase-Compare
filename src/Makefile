# Путь к директориям
COMPARE_CMD := ./cmd/compare
FS_CMD := ./cmd/fs
BIN_NAME := readDB
COMPARE_NAME := compareDB
FS_NAME := compareFS
# Сборка парсера (readDB)
build:
	go build -o $(BIN_NAME) ./cmd
# Сборка сравнения баз данных (compareDB)
compare:
	go build -o $(COMPARE_NAME) $(COMPARE_CMD)
# Сборка сравнения файловых систем (compareFS)
fs:
	go build -o $(FS_NAME) $(FS_CMD)
# Удаление бинарников
clean:
	rm -f $(BIN_NAME) $(COMPARE_NAME) $(FS_NAME)
# Тест запуск с JSON-файлом
json: build
	./$(BIN_NAME) -f assets/stolen_database.json
# Тест запуск с XML-файлом
xml: build
	./$(BIN_NAME) -f assets/original_database.xml
# Форматирование кода
fmt:
	go fmt ./...
# Проверка стиля
lint:
	golint ./...
# Запуск всех тестов
tests:
	go test ./...