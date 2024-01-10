package main

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	APIKey    string
	DebugMode bool
	DummyAPI  bool
}

func initConfig() *Config {
	godotenv.Load()

	apiKey := os.Getenv("SHAI_OPENAI_API_KEY")
	debug := getConfigBool("SHAI_DEBUG")
	dummyAPI := getConfigBool("DUMMY_API")

	return &Config{
		APIKey:    apiKey,
		DebugMode: debug,
		DummyAPI:  dummyAPI,
	}
}

func getConfigBool(envName string) bool {
	val := os.Getenv(envName)
	return val != "" && val != "0"

}
