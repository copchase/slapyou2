package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/copchase/slapyou2/internal/application"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()

		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		<-signalChan
	}()

	app := application.NewApplication(ctx)
	app.Run()
}
