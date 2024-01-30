package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DiscordAuthToken string
}

func GetEnv() *Env {
	godotenv.Load()

	env := Env{
		DiscordAuthToken: os.Getenv("DISCORD_AUTH_TOKEN"),
	}

	return &env
}
