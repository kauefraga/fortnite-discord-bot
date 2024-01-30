package commands

import "github.com/bwmarrin/discordgo"

func Hello(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "Eae, bão ^^")
	}
}
