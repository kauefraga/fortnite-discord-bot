package config

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	DiscordAuthToken string
	FortniteApiKey   string
}

var Env *env

func GetEnv() {
	godotenv.Load()

	Env = &env{
		DiscordAuthToken: os.Getenv("DISCORD_AUTH_TOKEN"),
		FortniteApiKey:   os.Getenv("FORTNITE_API_KEY"),
	}
}
