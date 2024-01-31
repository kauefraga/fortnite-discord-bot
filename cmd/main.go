package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kauefraga/lau-discord-bot/internal/discord"
)

func main() {
	discord := discord.InitSession()

	err := discord.Open()

	if err != nil {
		panic(err)
	}

	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	defer discord.Close()
}
