package main

import (
	"Kode_test/config"
	"Kode_test/internal/storage/postgres"
	"Kode_test/pkg/logger/sl"
	"golang.org/x/exp/slog"
	"os"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	log.Info("starting_logger", slog.String("env", cfg.Env))
	log.Debug("debug massages are enabled")

	storage, err := postgres.New(cfg.DB)
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
	//TODO init router

	//TODO run server
}
