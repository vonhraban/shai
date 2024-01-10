package main

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	APIKey    string
	DebugMode bool
	DummyAPI  bool
	APIHost   string
}

const defaultAPIHost = "https://api.openai.com"

func initConfig() *Config {
	godotenv.Load()

	apiKey := os.Getenv("SHAI_OPENAI_API_KEY")
	debug := getConfigBool("SHAI_DEBUG")
	dummyAPI := getConfigBool("DUMMY_API")

	return &Config{
		APIKey:    apiKey,
		DebugMode: debug,
		DummyAPI:  dummyAPI,
		APIHost:   defaultAPIHost, // make this configurable for the use with other models
	}
}

func getConfigBool(envName string) bool {
	val := os.Getenv(envName)
	return val != "" && val != "0"

}
