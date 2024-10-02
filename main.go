package main

import (
	"InfSec/pkg/logger"
	"InfSec/src/config"
	"InfSec/src/transport/rest/server"
	"context"
)

func main() {
	// Инициализация контекста приложения
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Загрузка конфигурации
	cfg, _ := config.MustLoad()

	// Инициализация логгера
	log := logger.NewLogger(cfg.ENV)

	//// Инициализация базы данных и проведение миграций
	//db, err := database.New(ctx, cfg)
	//if err != nil {
	//	log.Error("Error initializing database: %v", err)
	//	return err
	//}
	//defer db.Close()

	srv := server.New(cfg, log)
	if err := srv.Run(ctx); err != nil {
		log.Fatal("Error initializing server: %v", err)
	}
}
