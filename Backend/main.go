package main

import (
	"fmt"
	"log"
	"net/http"

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

	// Запуск сервера
	port := ":" + cfg.ServerPort

	if cfg.UseHTTPS {
		fmt.Println("Starting HTTPS server on port", cfg.ServerPort)
		server := &http.Server{
			Addr:    port,
			Handler: router,
		}
		if err := server.ListenAndServeTLS(cfg.SSLCertPath, cfg.SSLKeyPath); err != nil {
			log.Fatal("Failed to start HTTPS server:", err)
		}
	} else {
		fmt.Println("Starting HTTP server on port", cfg.ServerPort)
		if err := router.Run(port); err != nil {
			log.Fatal("Failed to start HTTP server:", err)
		}
	}
}
