package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/gohxap/etc"
	"github.com/HsiaoCz/gohxap/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		slog.Error("load env error", "err", err)
		return
	}

	var (
		homePageHandler = handlers.NewHomePageHandler()
		port            = etc.GetPort("PORT")
	)

	r := gin.Default()
	r.GET("/", homePageHandler.HandleHome)

	srv := http.Server{
		Handler:      r,
		Addr:         port,
		ReadTimeout:  time.Millisecond * 1500,
		WriteTimeout: time.Millisecond * 1500,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen : %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown error %v\n", err)
	}

	slog.Info("the server is shutdown")
}
