package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	ch, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Printf("Error to delete messages: %v", err)
		s.ChannelMessageSend(m.ChannelID, "Error to delete messages")
	}

	Music(ch.Name, s, m)
	Wallpaper(ch.Name, s, m)
}
