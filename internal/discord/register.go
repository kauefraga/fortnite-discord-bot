package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kauefraga/lau-discord-bot/internal/discord/commands"
)

func RegisterSlashCommands(s *discordgo.Session) {
	_, err := s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "pickgame",
		Description: "Escolhe um jogo aleatoriamente",
	})

	if err != nil {
		panic(err)
	}
}

func registerCommands(s *discordgo.Session) {
	s.AddHandler(commands.Hello)
	s.AddHandler(commands.PickGame)
	s.AddHandler(commands.SlashPickGame)
}
