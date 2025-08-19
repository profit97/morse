package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	possiblePaths := []string{
		"index.html",
		"../index.html",
	}

	var data []byte
	var err error

	for _, path := range possiblePaths {
		data, err = os.ReadFile(path)
		if err == nil {
			log.Printf("Успешно загружен index.html из: %s", path)
			break
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fmt.Println("Загрузка файла:", handler.Filename)

	// Читаем данные из файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Передаем данные в функцию автоопределения
	convertedString, _ := service.ConvertString(string(data))

	// Создаем локальный файл
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	fileName := timestamp + filepath.Ext(handler.Filename)
	outFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Записываем в локальный файл результат конвертации строки
	_, err = outFile.Write([]byte(convertedString))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Конвертация выполнена: " + convertedString))
}
