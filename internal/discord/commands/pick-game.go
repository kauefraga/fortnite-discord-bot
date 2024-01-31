package commands

import (
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
var defaultGames = [9]string{
	"Minecraft",
	"Fall Guys",
	"Among Us",
	"Phasmophobia",
	"Fortnite",
	"Don't Starve Together",
	"Pummel Party",
	"Stardew Valley",
	"Rocket League",
}

func GetRandomGame(games []string) string {
	i := rng.Intn(len(games) - 1)

	return games[i]
}

func SlashPickGame(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: GetRandomGame(defaultGames[:]),
		},
	})
}

func PickGame(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	cmd := strings.Split(m.Content, " ")

	if cmd[0] == "!pickgame" && len(cmd) == 1 {
		s.ChannelMessageSend(m.ChannelID, GetRandomGame(defaultGames[:]))
	}

	if cmd[0] == "!pickgame" && len(cmd) > 1 {
		s.ChannelMessageSend(m.ChannelID, GetRandomGame(cmd[1:]))
	}
}
