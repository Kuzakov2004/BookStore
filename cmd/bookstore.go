package main

import (
	"BookStore/internal/app"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a := app.NewStoreApp()
	if a == nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		<-c
		cancel()
	}()

	a.Run(ctx)
}
