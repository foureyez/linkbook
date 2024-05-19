package http

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/foureyez/linkbook/config"
	"github.com/foureyez/linkbook/internal/api"
	"github.com/foureyez/linkbook/internal/logger"
)

func StartServer(ctx context.Context, config *config.Server, handlers []api.Handler) error {
	mux := http.NewServeMux()
	for _, h := range handlers {
		if err := h.RegisterRoutes(mux); err != nil {
			return err
		}
	}

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: mux,
	}

	go func() {
		logger.Get().Infof("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Get().Fatal("error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.Get().Errorf("error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
