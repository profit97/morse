package main

import (
	//"fmt"
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаём логгер
	logger := log.New(os.Stderr, "logger: ", log.LstdFlags)

	// Создаём сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	log.Print("Сервер запущен")
	err := srv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
