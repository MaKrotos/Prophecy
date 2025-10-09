package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"prophecy/backend/config"
	"prophecy/backend/database"
	"prophecy/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg := config.GetConfig()

	// Инициализация подключения к базе данных
	database.InitDB()
	defer database.DB.Close()

	// Установка режима релиза для Gin
	gin.SetMode(gin.ReleaseMode)

	// Создание роутера Gin
	router := gin.Default()

	// Регистрация маршрутов
	routes.RegisterRoutes(router)

	var wg sync.WaitGroup

	// Запуск HTTP сервера
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpPort := ":" + cfg.ServerPort
		fmt.Println("Starting HTTP server on port", cfg.ServerPort)
		if err := http.ListenAndServe(httpPort, router); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	// Запуск HTTPS сервера (если включен)
	if cfg.UseHTTPS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			httpsPort := ":8443" // используйте другой порт для HTTPS
			fmt.Println("Starting HTTPS server on port 8443")
			if err := http.ListenAndServeTLS(httpsPort, cfg.SSLCertPath, cfg.SSLKeyPath, router); err != nil {
				log.Printf("HTTPS server error: %v", err)
			}
		}()
	}

	// Ожидаем завершения всех серверов
	wg.Wait()
}
