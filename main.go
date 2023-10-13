package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/tranleconghau77/go-restful/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)


	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
