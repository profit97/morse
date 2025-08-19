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
	// Читаем содержимое файла index.html
	file, err := os.Open("../index.html")
	if err != nil {
		log.Println("Ошибка при открытии файла:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Отправляем содержимое файла в ответ
	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму
	err := r.ParseMultipartForm(0)
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
