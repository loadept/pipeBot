package util

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func GetVoiceChannel(s *discordgo.Session, userID string) (string, error) {
	guilds := s.State.Guilds
	for _, guild := range guilds {
		for _, vc := range guild.VoiceStates {
			if vc.UserID == userID {
				return vc.ChannelID, nil
			}
		}
	}
	return "", fmt.Errorf("User not in voice channel")
}

func GetGuildChannel(s *discordgo.Session, guildID string, matchString string) string {
	channels, err := s.GuildChannels(guildID)
	if err != nil {
		return ""
	}

	for _, ch := range channels {
		match, err := regexp.MatchString(matchString, ch.Name)
		if err != nil {
			return ""
		}
		if !match {
			continue
		}

		return ch.ID
	}

	return ""
}
