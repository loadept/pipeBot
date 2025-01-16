package handler

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
)

func AsyncEvent(s *discordgo.Session, channels []string) {
	ticker := time.NewTicker(24 * time.Hour * 13)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, channelID := range channels {
				var lastMessageID string
				messagesDeleted := 0
				for {
					var messageIDs []string
					channelMessages, err := s.ChannelMessages(channelID, 100, lastMessageID, "", "")
					if err != nil {
						throwable.SendErrorEmbed(s, channelID, "Error getting messages to delete")
						break
					}
					if len(channelMessages) == 0 {
						break
					}

					for _, m := range channelMessages {
						messageIDs = append(messageIDs, m.ID)
					}

					if len(messageIDs) > 0 {
						if err := s.ChannelMessagesBulkDelete(channelID, messageIDs); err != nil {
							throwable.SendErrorEmbed(s, channelID, "Error deleting messages from channel")
							continue
						}
						messagesDeleted += len(messageIDs)
					}
					lastMessageID = channelMessages[len(channelMessages)-1].ID
					time.Sleep(500 * time.Millisecond)
				}

				if messagesDeleted > 0 {
					embed := &discordgo.MessageEmbed{
						Title:       "✅ Cleanup Complete",
						Description: fmt.Sprintf("%d messages have been deleted successfully", messagesDeleted),
						Color:       0xe8ff00,
					}
					s.ChannelMessageSendEmbed(channelID, embed)
				}
			}
		}
	}
}
