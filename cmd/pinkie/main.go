package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kironono/pinkie/config"
	"github.com/kironono/pinkie/server"
)

func main() {
	if err := serve(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}

}

func serve(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed load config: %w", err)
	}

	port := cfg.HTTPPort
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen port %d: %w", port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %s", url)

	mux, cleanup, err := server.NewMux(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed: %w", err)
	}
	defer cleanup()

	s := server.NewServer(l, mux)
	return s.Run(ctx)
}
