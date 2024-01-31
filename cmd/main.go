package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kauefraga/lau-discord-bot/internal/discord"
)

func main() {
	s := discord.InitSession()

	err := s.Open()

	if err != nil {
		panic(err)
	}

	discord.RegisterSlashCommands(s)

	log.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	defer s.Close()
}
