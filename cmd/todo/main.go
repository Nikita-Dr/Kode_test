package main

import (
	"Kode_test/config"
	"golang.org/x/exp/slog"
	"log"
	"os"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	log.Info("starting_logger", slog.String("env", cfg.Env))
	log.Debug("debug massages are enabled")

	//TODO init logger

	//TODO init storage

	//TODO init router

	//TODO run server
}
