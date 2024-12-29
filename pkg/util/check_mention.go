package util

import (
	"github.com/bwmarrin/discordgo"
	"github.com/loadept/pipeBot/pkg/throwable"
)

func CheckMentions(s *discordgo.Session, channelID string, mentions []*discordgo.User) error {
	// if len(content) < 3 {
	// 	return throwable.MisusedCommand
	// }
	if len(mentions) == 0 {
		return throwable.NoMentions
	}
	if len(mentions) > 1 {
		return throwable.MultipleMentions
	}

	return nil
}

func CheckMentionsRole(s *discordgo.Session, channelID string, roleMention []string, mentions []*discordgo.User) error {
	if len(mentions) == 0 {
		return throwable.NoMentions
	}
	if len(mentions) > 1 {
		return throwable.MultipleMentions
	}
	if len(roleMention) == 0 {
		return throwable.NoMentionsRole
	}
	if len(roleMention) > 1 {
		return throwable.MultipleMentionsRole
	}

	return nil
}
