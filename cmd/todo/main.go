package main

import (
	"Kode_test/config"
	"Kode_test/internal/app"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	// Запуск приложения
	applocation := app.NewApp(cfg)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return applocation.Run()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return applocation.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
