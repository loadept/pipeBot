package util

import (
	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
)

func CheckMentions(s *discordgo.Session, channelID string, content []string, mentions []*discordgo.User) error {
	if len(content) < 3 {
		return throwable.MisusedCommand
	}
	if len(mentions) == 0 {
		return throwable.NoMentions
	}
	if len(mentions) > 1 {
		return throwable.MultipleMentions
	}

	return nil
}
