package config

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	DiscordAuthToken string
}

var Env *env

func GetEnv() {
	godotenv.Load()

	Env = &env{
		DiscordAuthToken: os.Getenv("DISCORD_AUTH_TOKEN"),
	}
}
