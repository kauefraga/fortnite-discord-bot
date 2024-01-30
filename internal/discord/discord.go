package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kauefraga/fortnite-discord-bot/internal/config"
)

func InitSession() *discordgo.Session {
	env := config.GetEnv()

	discord, err := discordgo.New("Bot " + env.DiscordAuthToken)

	if err != nil {
		panic(err)
	}

	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Println("Bot is ready")
	})

	registerCommands(discord)

	return discord
}
