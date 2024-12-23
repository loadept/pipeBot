package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/util"
)

type MusicChannel struct{}

func (ms *MusicChannel) NotifyMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	ch, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Printf("Error to delete messages: %v", err)
		s.ChannelMessageSend(m.ChannelID, "Error to delete messages")
	}

	if util.CheckChName(ch.Name, ".*música$|.*music$") {
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🔴 Invalid action",
		Description: "You can't play music here",
		Color:       0xff0000,
	}

	if m.Interaction != nil && m.Interaction.Name == "play" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}
