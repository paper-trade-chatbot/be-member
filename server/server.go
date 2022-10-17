package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/paper-trade-chatbot/be-member/api"
	"github.com/paper-trade-chatbot/be-member/config"
	"github.com/paper-trade-chatbot/be-member/global"
	"github.com/paper-trade-chatbot/be-member/logging"
)

// CreateServer creates an HTTP server listening on the specified address.
func CreateHttpServer(ctx context.Context, address string) *http.Server {
	// Setup HTTP Server.
	server := &http.Server{
		Addr:    address,
		Handler: api.GetRouter(),
	}

	// Install the shutdown handler.
	installShutdownHandler(ctx, server)

	return server
}

// installShutdownHandler registers a shutdown handler for graceful shutdown.
func installShutdownHandler(ctx context.Context, server *http.Server) {
	// Create signal channel & shutdown timeout context.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	timeoutCtx, cancel := context.WithTimeout(ctx,
		config.GetMilliseconds("SERVER_SHUTDOWN_GRACE_PERIOD_MS"))

	// Catch signals in a separate goroutine.
	go func() {
		defer cancel()

		// Wait for signals.
		sig := <-sigChan
		signal.Stop(sigChan)
		logging.Warn("Received signal: %s.", sig.String())

		// Perform graceful shutdown.
		logging.Warn("Initiating graceful shutdown...")
		global.Alive = false
		if err := server.Shutdown(timeoutCtx); err != nil {
			logging.Error("Failed to shutdown: %s", err.Error())
		}
	}()
}

// CreateGRpcServer creates an Grpc server listening on the specified address.
func CreateGRpcServer(ctx context.Context, address string) *net.Listener {

	apiListener, err := net.Listen("tcp", address)
	if err != nil {
		logging.Debug(err.Error())
		panic(err)
	}

	logging.Info("register grpc ...")
	logging.Info("listening grpc on: %s", address)
	logging.Info("register grpc done")
	return &apiListener
}
