package main

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	APIKey    string
	DebugMode bool
}

func initConfig() *Config {
	godotenv.Load()

	apiKey := os.Getenv("SHAI_OPENAI_API_KEY")

	debugEnvVar := os.Getenv("SHAI_DEBUG")
	debug := debugEnvVar != "" && debugEnvVar != "0"

	return &Config{
		APIKey:    apiKey,
		DebugMode: debug,
	}
}
