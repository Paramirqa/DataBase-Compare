package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/compare", compareHandler)

	fmt.Println("Server started on :8080")
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

	// Получаем первый файл
	file1, handler1, err := r.FormFile("file1")
	if err != nil {
		http.Error(w, "Failed to get file1", http.StatusBadRequest)
		return
	}
	defer file1.Close()

	// Получаем второй файл
	file2, handler2, err := r.FormFile("file2")
	if err != nil {
		http.Error(w, "Failed to get file2", http.StatusBadRequest)
		return
	}
	defer file2.Close()

	// Сохраняем файлы во временную папку
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

	// Вызываем функцию сравнения баз данных
	err = compareDatabases(path1, path2)
	if err != nil {
		http.Error(w, fmt.Sprintf("Comparison failed: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Databases compared successfully!")
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

// compareDatabases - сюда подключаешь свою логику
func compareDatabases(path1, path2 string) error {
	fmt.Printf("Comparing %s and %s\n", path1, path2)
	return nil
}
