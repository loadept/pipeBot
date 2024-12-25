package util

import (
	"fmt"

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
