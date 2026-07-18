package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

// Fetching the env variables and reporting back any missing values
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Environmental variable %s is not set, Please make sure to double check it before proceeding.", key)
	}
	return value
}

// Defining global values for pulling env variables in other files
func EnvMongoURI()  string { return getEnv("MONGO_URI")}
func EnvJWTSecret() string { return getEnv("JWT_SECRET") }
func EnvDiscordClient() string { return getEnv("DISCORD_CLIENT_ID")}
func EnvDiscordSecret() string { return getEnv("DISCORD_CLIENT_SECRET") }
func EnvDiscordBotToken() string { return getEnv("DISCORD_BOT_TOKEN") }
