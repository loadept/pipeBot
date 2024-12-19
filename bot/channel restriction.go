package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/util"
)

func Wallpaper(channelName string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if !util.CheckChName(channelName, ".*wallpapers?$") {
		return
	}

	if len(m.Attachments) == 0 || !strings.HasPrefix(m.Attachments[0].ContentType, "image/") {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
	}
}

func Music(channelName string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if util.CheckChName(channelName, ".*música$|.*music$") {
		return
	}

	embed := &discordgo.MessageEmbed{
		Title:       "🔴 Invalid action",
		Description: "You can't play music here",
		Color:       0xff0000,
	}

	if m.Interaction != nil && m.Interaction.Name == "play" {
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
		s.ChannelMessageDelete(m.ChannelID, m.ID)
	}
}
