package main

import (
	"context"
	"fmt"
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Onefootball/simple-service/db"

	"github.com/Onefootball/simple-service/config"
	"github.com/Onefootball/simple-service/http"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	cfg := config.NewConfig()
	pg := db.NewPostgres(cfg.PostgresURL)

	router := chi.NewRouter()
	router.Get("/live", netHttp.HandlerFunc(
		func(w netHttp.ResponseWriter, r *netHttp.Request) {
			http.NewHandler(pg).ServeHTTP(w, r)
		},
	))

	server := &netHttp.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	go func() {
		fmt.Printf("Listening on port %d\n", cfg.Port)
		if err := server.ListenAndServe(); err != nil {
			if err == netHttp.ErrServerClosed {
				fmt.Printf("closing http server...\ns")
				return
			}
			os.Exit(1)
		}
	}()

	sig := <-signalChan
	fmt.Printf("Received signal: %q, shutting down...", sig.String())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Avoid context leak

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("could not gracefully shutdown\n")
		os.Exit(1)
	}
}
