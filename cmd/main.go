package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	companiesadpt "github.com/IsaErtunga/companies/internal/adapters/companies"
	companiessrv "github.com/IsaErtunga/companies/internal/core/services/companies_srv"
	companiesrepo "github.com/IsaErtunga/companies/internal/repositories/companies_repo"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	companiesRepository := companiesrepo.NewMemKVS()
	companiesService := companiessrv.New(companiesRepository)
	companiesAdapter := companiesadpt.NewHTTPAdapter(companiesService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	router.Mount("/companies", companiesAdapter.Routes())

	server := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      router,            // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
