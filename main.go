package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/jsbento/butterfly-api/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

		<-c
		cancel()
	}()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbUri := os.Getenv("MONGODB_URI")

	s := server.NewServer(ctx, dbUri)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return s.Start()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return s.Stop(context.Background())
	})

	if err := g.Wait(); err != nil {
		panic(err)
	}
}
