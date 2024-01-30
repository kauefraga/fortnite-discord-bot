package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kauefraga/fortnite-discord-bot/internal/discord/commands"
)

func registerCommands(s *discordgo.Session) {
	s.AddHandler(commands.Hello)
}
