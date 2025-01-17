package handler

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
	"github.com/loadept/pipeBot/pkg/util"
)

type WallpaperChannel struct{}

func (wp *WallpaperChannel) NotifyMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	if !util.CheckChName(ch.Name, ".*wallpapers?$") {
		return
	}

	if len(m.Attachments) == 0 || !strings.HasPrefix(m.Attachments[0].ContentType, "image/") {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		throwable.SendErrorEmbed(
			s,
			logsChannel,
			"An attempt was made to upload unauthorized content to the wallpapers channel.",
		)
	}
}
