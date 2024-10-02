package app

import (
	"InfSec/internal/config"
	"InfSec/internal/transport/rest/server"
	"InfSec/pkg/logger"
	"context"
)

// Run запускает основное приложение
func Run() error {
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
		log.Error("Error initializing server: %v", err)
		return err
	}

	return nil
}
