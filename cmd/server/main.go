package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0daryo/staat/src/di"
	"github.com/0daryo/staat/src/route"
	"github.com/go-chi/chi"
)

const servicename = "staat"

var (
	revision string // Assigned at Build by Makefile
	addr     = ":3001"
)

func main() {
	var errLoggerGen error
	if errLoggerGen != nil {
		panic(errLoggerGen)
	}
	baseCtx := context.Background()

	// Dependency
	d := di.Dependency{}
	d.Inject(baseCtx)

	// Routing
	r := chi.NewRouter()
	route.Routing(r, d)

	// server
	server := http.Server{
		Addr:    addr,
		Handler: chi.ServerBaseContext(baseCtx, r),
	}

	// Run
	fmt.Println("main: start server")
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Println("main: closed server", err)
		}
	}()

	// graceful shuttdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	fmt.Printf("SIGNAL %d received, so server shutting down now...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Println("failed to gracefully shutdown")
	}

	fmt.Println("main: server shutdown completed")
}
