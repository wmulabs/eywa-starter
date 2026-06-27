package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wmulabs/eywa-starter/internal/infrastructure/bootstrap"
	"github.com/wmulabs/eywa-starter/internal/infrastructure/driver/rest"
)

func main() {
	ctx := context.Background()

	app, err := bootstrap.Initialize(ctx)
	if err != nil {
		log.Fatalf("initialize: %v", err)
	}
	defer app.Shutdown(ctx)

	server, err := rest.NewServer(rest.Deps{
		Weave:             app.Engine.Weave,
		Port:              app.Config.Server.Port,
		AdminAPIKey:       app.Config.Admin.APIKey,
		RequireEventToken: app.Config.RequireEventToken,
		Echo:              app.Repositories.Echo,
		EchoQuery:         app.Repositories.Echo,
		Chronicle:         app.Repositories.Chronicle,
		AppToken:          app.Repositories.AppToken,
	})
	if err != nil {
		log.Fatalf("server: %v", err)
	}

	go func() {
		log.Printf("listening on :%s", app.Config.Server.Port)
		if err := server.Start(); err != nil {
			log.Fatalf("server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down")
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("shutdown: %v", err)
	}
}
