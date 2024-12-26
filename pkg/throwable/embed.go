package throwable

import "github.com/bwmarrin/discordgo"

func SendErrorEmbed(s *discordgo.Session, channelID, description string) {
	embed := &discordgo.MessageEmbed{
		Title:       "🔴 Invalid action!",
		Description: description,
		Color:       0xff0000,
	}
	s.ChannelMessageSendEmbed(channelID, embed)
}
