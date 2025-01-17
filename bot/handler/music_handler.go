package handler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
	"github.com/loadept/pipeBot/pkg/util"
)

type MusicChannel struct{}

func (ms *MusicChannel) NotifyMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	logsChannel := util.GetGuildChannel(s, m.GuildID, ".*logs?")
	if len(logsChannel) == 0 {
		throwable.SendErrorEmbed(s, logsChannel, throwable.SomethingWentWrong.Error())
	}

	ch, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Printf("Error to delete messages: %v", err)
		s.ChannelMessageSend(logsChannel, "Error to delete messages")
	}

	if util.CheckChName(ch.Name, ".*música$|.*music$") {
		return
	}

	if m.Interaction != nil && m.Interaction.Name == "play" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		throwable.SendErrorEmbed(s, m.ChannelID, "You can't play music here")
	}
}
