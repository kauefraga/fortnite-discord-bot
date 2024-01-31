package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kauefraga/lau-discord-bot/internal/discord/commands"
)

func registerCommands(s *discordgo.Session) {
	s.AddHandler(commands.Hello)
	s.AddHandler(commands.DailyShop)
}
