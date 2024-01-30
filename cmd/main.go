package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/kauefraga/fortnite-discord-bot/internal/config"
)

func main() {
	env := config.GetEnv()

	discord, err := discordgo.New("Bot " + env.DiscordAuthToken)

	if err != nil {
		panic(err)
	}

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})

	err = discord.Open()

	if err != nil {
		panic(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	defer discord.Close()
}
