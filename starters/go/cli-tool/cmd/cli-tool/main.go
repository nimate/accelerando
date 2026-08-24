package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

type Config struct {
	ServiceName string
	Name        string
}

func newLogger(service string) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})).With("service", service)
}

func loadConfig(args []string) (Config, error) {
	fs := flag.NewFlagSet("cli-tool", flag.ContinueOnError)

	cfg := Config{
		ServiceName: "cli-tool",
		Name:        "world",
	}

	fs.StringVar(&cfg.Name, "name", cfg.Name, "name to greet")
	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}

	if raw := os.Getenv("SERVICE_NAME"); raw != "" {
		cfg.ServiceName = raw
	}

	return cfg, nil
}

func main() {
	cfg, err := loadConfig(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	logger := newLogger(cfg.ServiceName)
	logger.Info("command started", "name", cfg.Name)
	fmt.Printf("hello, %s\n", cfg.Name)
}
