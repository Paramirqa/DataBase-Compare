package main

import (
	"Go_day01-1/nydiamig/internal/parser"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	// 1. Отдавать статические файлы из папки static
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// 2. Роут для /compare (обработка формы)
	http.HandleFunc("/compare", compareHandler)

	fmt.Println("Server started on :8080")

	go func() {
		// Пытаемся открыть браузер автоматически
		err := openBrowser("http://localhost:8080")
		if err != nil {
			fmt.Println("Failed to open browser:", err)
		}
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// compareHandler принимает два файла и вызывает сравнение
func compareHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file1, handler1, err := r.FormFile("file1")
	if err != nil {
		http.Error(w, "Failed to get file1", http.StatusBadRequest)
		return
	}
	ext1 := strings.ToLower(filepath.Ext(handler1.Filename))
	if ext1 != ".json" && ext1 != ".xml" {
		http.Error(w, "Unsupported file format for file1", http.StatusBadRequest)
		return
	}
	defer file1.Close()

	file2, handler2, err := r.FormFile("file2")
	if err != nil {
		http.Error(w, "Failed to get file2", http.StatusBadRequest)
		return
	}
	ext2 := strings.ToLower(filepath.Ext(handler2.Filename))
	if ext2 != ".json" && ext2 != ".xml" {
		http.Error(w, "Unsupported file format for file2", http.StatusBadRequest)
		return
	}
	defer file2.Close()

	path1, err := saveUploadedFile(file1, handler1.Filename)
	if err != nil {
		http.Error(w, "Failed to save file1", http.StatusInternalServerError)
		return
	}

	path2, err := saveUploadedFile(file2, handler2.Filename)
	if err != nil {
		http.Error(w, "Failed to save file2", http.StatusInternalServerError)
		return
	}

	// Новый способ сравнения баз
	diffResult, err := compareDatabases(path1, path2)
	if err != nil {
		http.Error(w, fmt.Sprintf("Comparison failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diffResult)
}

// saveUploadedFile сохраняет файл в /tmp/uploads
func saveUploadedFile(file io.Reader, filename string) (string, error) {
	dir := "/tmp/uploads"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, filepath.Base(filename))
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return path, nil
}

// compareDatabases - загружает и сравнивает базы данных
func compareDatabases(path1, path2 string) (parser.DiffResult, error) {
	recipe1, err := loadRecipe(path1)
	if err != nil {
		return parser.DiffResult{}, fmt.Errorf("failed to load first database: %w", err)
	}

	recipe2, err := loadRecipe(path2)
	if err != nil {
		return parser.DiffResult{}, fmt.Errorf("failed to load second database: %w", err)
	}

	diff := parser.CompareRecipes(recipe1, recipe2)
	return diff, nil
}

// loadRecipe - утилита для чтения файла и парсинга
func loadRecipe(filePath string) (parser.Recipe, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return parser.Recipe{}, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	format := parser.DetectFormatFromExtension(filePath)
	if !parser.IsSupportedFormat(format) {
		return parser.Recipe{}, fmt.Errorf("unsupported or unknown file format: %s", format)
	}

	recipe, err := parser.ParseData(format, data)
	if err != nil {
		return parser.Recipe{}, fmt.Errorf("failed to parse %s file: %w", format, err)
	}

	return recipe, nil
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch os := runtime.GOOS; os {
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	default: // Linux и прочие
		cmd = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmd, args...).Start()
}
