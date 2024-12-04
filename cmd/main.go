package main

import (
	"flag"
	"log/slog"

	"github.com/Abhinav7903/split/server"
)

func main() {
	envType := flag.String("env", "dev", "Environment type production or development")
	flag.Parse()
	slog.Info("Environment type: ", "env", *envType)
	slog.Info("Starting server...")
	server.Run(envType)
}
